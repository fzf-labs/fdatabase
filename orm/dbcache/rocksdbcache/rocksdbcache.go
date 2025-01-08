package rocksdbcache

import (
	"context"
	"strings"
	"time"

	"github.com/dtm-labs/rockscache"
	"github.com/fzf-labs/fdatabase/orm/dbcache"
	"golang.org/x/sync/errgroup"
)

type Cache struct {
	name              string             // 缓存名称
	rocksCache        *rockscache.Client // RocksCache缓存客户端
	redisTTL          time.Duration      // redis缓存过期时间
	redisLuaBatchSize int                // redis lua 批量查询数量  默认100 有些云厂商对lua的keys有限制
}

func NewRocksDBCache(rocksCache *rockscache.Client, opts ...CacheOption) *Cache {
	r := &Cache{
		name:              "GormCache",
		rocksCache:        rocksCache,
		redisTTL:          time.Hour * 24,
		redisLuaBatchSize: 100,
	}
	if len(opts) > 0 {
		for _, v := range opts {
			v(r)
		}
	}
	return r
}

type CacheOption func(cache *Cache)

// WithName 设置缓存名称
func WithName(name string) CacheOption {
	return func(r *Cache) {
		r.name = name
	}
}

// WithRedisTTL 设置redis缓存过期时间
func WithRedisTTL(ttl time.Duration) CacheOption {
	return func(r *Cache) {
		r.redisTTL = ttl
	}
}

// WithRedisLuaBatchSize 设置RocksCache批量查询数量
func WithRedisLuaBatchSize(batchSize int) CacheOption {
	return func(r *Cache) {
		r.redisLuaBatchSize = batchSize
	}
}

func (r *Cache) Key(keys ...any) string {
	keyStr := make([]string, 0)
	keyStr = append(keyStr, r.name)
	for _, v := range keys {
		keyStr = append(keyStr, dbcache.KeyFormat(v))
	}
	return strings.Join(keyStr, ":")
}

func (r *Cache) Fetch(ctx context.Context, key string, fn func() (string, error)) (string, error) {
	return r.Fetch2(ctx, key, fn, r.redisTTL)
}

func (r *Cache) Fetch2(ctx context.Context, key string, fn func() (string, error), expire time.Duration) (string, error) {
	// 查询redis缓存
	rocksCacheValue, err := r.rocksCache.Fetch2(ctx, key, expire, fn)
	if err != nil {
		return "", err
	}
	return rocksCacheValue, nil
}

func (r *Cache) FetchBatch(ctx context.Context, keys []string, fn func(miss []string) (map[string]string, error)) (map[string]string, error) {
	return r.FetchBatch2(ctx, keys, fn, r.redisTTL)
}

func (r *Cache) FetchBatch2(ctx context.Context, keys []string, fn func(miss []string) (map[string]string, error), expire time.Duration) (map[string]string, error) {
	resp := make(map[string]string)
	// 去重
	keys = unique(keys)
	// 查询redis缓存
	batch := chunk(keys, r.redisLuaBatchSize)
	// 使用`errgroup`并发查询
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(100)
	// 创建一个channel用于接收每个goroutine的结果
	resultCh := make(chan map[string]string, len(batch))
	for k := range batch {
		i := k
		g.Go(func() error {
			rocksCacheResult, err := r.fetchBatchItem(ctx, batch[i], fn, expire)
			if err != nil {
				return err
			}
			// 将结果发送到channel
			resultCh <- rocksCacheResult
			return nil
		})
	}
	// 等待所有goroutine执行完毕
	err := g.Wait()
	if err != nil {
		return nil, err
	}
	// 关闭channel
	close(resultCh)
	// 从channel中读取结果
	for result := range resultCh {
		for k, v := range result {
			resp[k] = v
		}
	}
	return resp, nil
}

// fetchBatchItem 批量查询
func (r *Cache) fetchBatchItem(ctx context.Context, keys []string, fn func(miss []string) (map[string]string, error), expire time.Duration) (map[string]string, error) {
	resp := make(map[string]string)
	// 查询redis缓存
	rocksCacheResult, err := r.rocksCache.FetchBatch2(ctx, keys, expire, func(idx []int) (map[int]string, error) {
		result := make(map[int]string)
		miss := make([]string, 0)
		for _, v := range idx {
			result[v] = ""
			miss = append(miss, keys[v])
		}
		dbValue, err := fn(miss)
		if err != nil {
			return nil, err
		}
		keyToInt := make(map[string]int)
		for k, v := range keys {
			keyToInt[v] = k
		}
		for k, v := range dbValue {
			result[keyToInt[k]] = v
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	for k, v := range rocksCacheResult {
		resp[keys[k]] = v
	}
	return resp, nil
}

func (r *Cache) Del(ctx context.Context, key string) error {
	err := r.rocksCache.TagAsDeleted2(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (r *Cache) DelBatch(ctx context.Context, keys []string) error {
	keys = unique(keys)
	batch := chunk(keys, r.redisLuaBatchSize)
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(100)
	for k := range batch {
		i := k
		g.Go(func() error {
			err := r.rocksCache.TagAsDeletedBatch2(ctx, batch[i])
			if err != nil {
				return err
			}
			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		return err
	}
	return nil
}

// unique 去重
func unique(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}
	// here no use map filter. if use it, the result slice element order is random, not same as origin slice
	var result []string
	for i := 0; i < len(slice); i++ {
		v := slice[i]
		skip := true
		for j := range result {
			if v == result[j] {
				skip = false
				break
			}
		}
		if skip {
			result = append(result, v)
		}
	}
	return result
}

// chunk 将一个数组分成多个数组，每个数组包含size个元素，最后一个数组可能包含少于size个元素。
func chunk(collection []string, size int) [][]string {
	if size <= 0 {
		panic("Second parameter must be greater than 0")
	}
	chunksNum := len(collection) / size
	if len(collection)%size != 0 {
		chunksNum += 1
	}
	result := make([][]string, 0, chunksNum)
	for i := 0; i < chunksNum; i++ {
		last := (i + 1) * size
		if last > len(collection) {
			last = len(collection)
		}
		result = append(result, collection[i*size:last])
	}
	return result
}
