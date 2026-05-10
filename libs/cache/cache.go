package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Del(ctx context.Context, key string) error
}

type redisCache struct {
	client *redis.Client
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *redisCache) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func NewCache(lc fx.Lifecycle) (Cache, error) {
	redisURL := os.Getenv("REDIS_CACHE_CONNECTION_STRING")
	if redisURL == "" {
		return nil, fmt.Errorf("REDIS_CACHE_CONNECTION_STRING environment variable is required")
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}

	client := redis.NewClient(opt)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Ping(ctx).Err()
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return &redisCache{client: client}, nil
}

var Module = fx.Module("cache",
	fx.Provide(NewCache),
)
