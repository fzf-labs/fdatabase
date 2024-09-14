package goredisdbcache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "0.0.0.0:6379",
})

func TestGoRedisCache_Fetch(t *testing.T) {
	cache := NewGoRedisDBCache(client, WithName("test"), WithTTL(time.Minute))
	ctx := context.Background()
	fetch, err := cache.Fetch(ctx, "GoRedisCache_Fetch", func() (string, error) {
		fmt.Println("do Fetch")
		return "GoRedisCache_Fetch: result", nil
	})
	fmt.Println(fetch)
	fmt.Println(err)
	assert.Equal(t, nil, err)
}

func TestGoRedisCache_FetchBatch(t *testing.T) {
	cache := NewGoRedisDBCache(client)
	ctx := context.Background()
	keys := []string{
		"GoRedisCache_Fetch_a",
		"GoRedisCache_Fetch_b",
		"GoRedisCache_Fetch_c",
		"GoRedisCache_Fetch_d",
	}
	fetch, err := cache.FetchBatch(ctx, keys, func(miss []string) (map[string]string, error) {
		fmt.Println("do FetchBatch")
		return map[string]string{
			"GoRedisCache_Fetch_a": "1",
			"GoRedisCache_Fetch_b": "2",
			"GoRedisCache_Fetch_c": "3",
		}, nil
	})
	fmt.Println(fetch)
	fmt.Println(err)
	assert.Equal(t, nil, err)
}

func TestCache_Del(t *testing.T) {
	cache := NewGoRedisDBCache(client)
	ctx := context.Background()
	err := cache.Del(ctx, "GoRedisCache_Fetch")
	if err != nil {
		return
	}
	assert.Equal(t, nil, err)
}

func TestCache_DelBatch(t *testing.T) {
	cache := NewGoRedisDBCache(client)
	ctx := context.Background()
	keys := []string{
		"GoRedisCache_Fetch_a",
		"GoRedisCache_Fetch_b",
		"GoRedisCache_Fetch_c",
		"GoRedisCache_Fetch_d",
	}
	err := cache.DelBatch(ctx, keys)
	if err != nil {
		return
	}
	assert.Equal(t, nil, err)
}

func TestCache_Key(t *testing.T) {
	cache := NewGoRedisDBCache(client, WithName("test"), WithTTL(time.Minute))
	key := cache.Key("a", "b", "c")
	assert.Equal(t, key, "test:a:b:c")
}