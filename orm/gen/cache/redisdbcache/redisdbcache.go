package redisdbcache

import (
	"context"
	"errors"
	"gitlab.yc345.tv/backend/utils/v2/orm/gen/cache"
	"math/rand"
	"strings"
	"time"

	redisV8 "github.com/go-redis/redis/v8"
	"gitlab.yc345.tv/backend/utils/v2/client/redis"
	"golang.org/x/sync/singleflight"
)

type Cache struct {
	name   string
	client redis.IClient
	ttl    time.Duration
	sf     singleflight.Group
}

func NewRedisDBCache(client redis.IClient, opts ...CacheOption) *Cache {
	r := &Cache{
		name:   "GormCache",
		client: client,
		ttl:    time.Hour * 24,
	}
	if len(opts) > 0 {
		for _, v := range opts {
			v(r)
		}
	}
	return r
}

type CacheOption func(cache *Cache)

func WithName(name string) CacheOption {
	return func(r *Cache) {
		r.name = name
	}
}

func WithTTL(ttl time.Duration) CacheOption {
	return func(r *Cache) {
		r.ttl = ttl
	}
}

func (r *Cache) Key(keys ...interface{}) string {
	keyStr := make([]string, 0)
	keyStr = append(keyStr, r.name)
	for _, v := range keys {
		keyStr = append(keyStr, cache.KeyFormat(v))
	}
	return strings.Join(keyStr, ":")
}

func (r *Cache) TTL(ttl time.Duration) time.Duration {
	return ttl - time.Duration(rand.Float64()*0.1*float64(ttl))
}

func (r *Cache) Fetch(ctx context.Context, key string, fn func() (string, error)) (string, error) {
	return r.Fetch2(ctx, key, fn, r.ttl)
}

func (r *Cache) Fetch2(ctx context.Context, key string, fn func() (string, error), expire time.Duration) (string, error) {
	do, err, _ := r.sf.Do(key, func() (interface{}, error) {
		result, err := r.client.Get(ctx, key)
		if err != nil && !errors.Is(err, redisV8.Nil) {
			return "", err
		}
		if result == "" && errors.Is(err, redisV8.Nil) {
			result, err = fn()
			if err != nil {
				return "", err
			}
			_, err := r.client.Set(ctx, key, result, r.TTL(expire))
			if err != nil {
				return nil, err
			}
		}
		return result, nil
	})
	if err != nil {
		return "", err
	}
	return do.(string), nil
}

func (r *Cache) FetchBatch(ctx context.Context, keys []string, fn func(miss []string) (map[string]string, error)) (map[string]string, error) {
	return r.FetchBatch2(ctx, keys, fn, r.ttl)
}

func (r *Cache) FetchBatch2(ctx context.Context, keys []string, fn func(miss []string) (map[string]string, error), expire time.Duration) (map[string]string, error) {
	resp := make(map[string]string)
	miss := make([]string, 0)
	//查询数据
	pipelineGet := r.client.Pipeline()
	for _, v := range keys {
		_, err := pipelineGet.Get(ctx, v).Result()
		if err != nil {
			return nil, err
		}
	}
	pipelineGetResult, err := pipelineGet.Exec(ctx)
	if err != nil && !errors.Is(err, redisV8.Nil) {
		return nil, err
	}
	for k, cmder := range pipelineGetResult {
		if errors.Is(cmder.Err(), redisV8.Nil) {
			miss = append(miss, keys[k])
		}
		resp[keys[k]] = cmder.(*redisV8.StringCmd).Val()
	}
	if len(miss) > 0 {
		dbValue, err := fn(miss)
		if err != nil {
			return nil, err
		}
		pipelineSet := r.client.Pipeline()
		for _, v := range miss {
			resp[v] = dbValue[v]
			err = pipelineSet.Set(ctx, v, dbValue[v], r.TTL(expire)).Err()
			if err != nil {
				return nil, err
			}
		}
		_, err = pipelineSet.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (r *Cache) Del(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (r *Cache) DelBatch(ctx context.Context, keys []string) error {
	_, err := r.client.Del(ctx, keys...)
	if err != nil {
		return err
	}
	return nil
}
