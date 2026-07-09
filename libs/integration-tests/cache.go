package integrationtests

import (
	"context"
	"testing"

	"github.com/kaua-nasc/gymtrack-go/libs/cache"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

func StartRedis(t *testing.T) (cache.Cache, func()) {
	t.Helper()

	ctx := context.Background()

	redisContainer, err := redis.Run(ctx, "redis:7-alpine")
	if err != nil {
		t.Fatalf("start redis container: %s", err)
	}
	t.Cleanup(func() {
		if err := redisContainer.Terminate(ctx); err != nil {
			t.Logf("failed to terminate redis container: %s", err)
		}
	})

	connStr, err := redisContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatalf("get redis connection string: %s", err)
	}

	client, err := cache.NewCache(connStr)
	if err != nil {
		t.Fatalf("connect redis: %s", err)
	}

	return client, func() { client.Close(ctx) }
}
