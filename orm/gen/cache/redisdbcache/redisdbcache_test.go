package redisdbcache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.yc345.tv/backend/utils/v2/client/redis"
)

var client = redis.NewClient(&redis.Config{
	Addr:     "10.8.8.117:6379",
	Password: "",
	DB:       7,
})

func TestCache_Fetch(t *testing.T) {
	cache := NewRedisDBCache(client, WithName("test"), WithTTL(time.Minute))
	ctx := context.Background()
	fetch, err := cache.Fetch(ctx, "RedisDBCache_Fetch", func() (string, error) {
		fmt.Println("do Fetch")
		return "RedisDBCache_Fetch:Result", nil
	})
	fmt.Println(fetch)
	fmt.Println(err)
	assert.Equal(t, nil, err)
}

func TestCache_FetchBatch(t *testing.T) {
	cache := NewRedisDBCache(client, WithName("test"), WithTTL(time.Minute))
	ctx := context.Background()
	keys := []string{
		"RedisDBCache_Fetch_a",
		"RedisDBCache_Fetch_b",
		"RedisDBCache_Fetch_c",
		"RedisDBCache_Fetch_d",
	}
	fetch, err := cache.FetchBatch(ctx, keys, func(miss []string) (map[string]string, error) {
		fmt.Println("do FetchBatch")
		return map[string]string{
			"RedisDBCache_Fetch_a": "1",
			"RedisDBCache_Fetch_b": "2",
			"RedisDBCache_Fetch_c": "3",
		}, nil
	})
	fmt.Println(fetch)
	fmt.Println(err)
	assert.Equal(t, nil, err)
}

func TestCache_Del(t *testing.T) {
	cache := NewRedisDBCache(client, WithName("test"), WithTTL(time.Minute))
	ctx := context.Background()
	err := cache.Del(ctx, "RedisDBCache_Fetch")
	if err != nil {
		return
	}
	assert.Equal(t, nil, err)
}

func TestCache_DelBatch(t *testing.T) {
	cache := NewRedisDBCache(client, WithName("test"), WithTTL(time.Minute))
	ctx := context.Background()
	keys := []string{
		"RedisDBCache_Fetch_a",
		"RedisDBCache_Fetch_b",
		"RedisDBCache_Fetch_c",
		"RedisDBCache_Fetch_d",
	}
	err := cache.DelBatch(ctx, keys)
	if err != nil {
		return
	}
	assert.Equal(t, nil, err)
}

func TestCache_Key(t *testing.T) {
	cache := NewRedisDBCache(client, WithName("test"), WithTTL(time.Minute))
	key := cache.Key("a", "b", "c")
	assert.Equal(t, key, "test:a:b:c")
}
