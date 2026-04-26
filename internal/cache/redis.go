package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func RedisClient(ctx context.Context, redisDSN string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisDSN)
	opt.PoolSize = 50 // depend on the througput
	opt.MaxRetries = 3
	opt.MinIdleConns = 10
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opt)
	// inital ping pong
	_, err = rdb.Ping(ctx).Result()
	return rdb, err
}
