package redisChache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	rdc *redis.Client
}

func NewRedisCache(rdc *redis.Client) *redisCache {
	return &redisCache{
		rdc: rdc,
	}
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	return r.rdc.Get(ctx, key).Result()
}

func (r *redisCache) Set(ctx context.Context, key, value string) error {
	fmt.Printf("here")
	return r.rdc.Set(ctx, key, value, 0).Err()
}
