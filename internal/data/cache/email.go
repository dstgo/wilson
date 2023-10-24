package cache

import (
	"context"
	"fmt"
	"github.com/dstgo/wilson/internal/data"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"time"
)

type EmailCodeCache interface {
	Get(ctx context.Context, code string) (string, error)
	Set(ctx context.Context, code string, email string, expire time.Duration) error
	Del(ctx context.Context, code string) error
	Check(ctx context.Context, code string) (string, error)
}

var EmailCodeCacheProvider = wire.NewSet(
	NewRedisEmailCodeCache,
	wire.Bind(new(EmailCodeCache), new(RedisEmailCodeCache)),
)

func NewRedisEmailCodeCache(source *data.DataSource) RedisEmailCodeCache {
	return RedisEmailCodeCache{
		cache: source.Redis(),
	}
}

func CodeCacheKey(code string) string {
	return fmt.Sprintf("email:code:%s", code)
}

type RedisEmailCodeCache struct {
	cache *redis.Client
}

func (c RedisEmailCodeCache) Get(ctx context.Context, code string) (string, error) {
	return c.cache.Get(ctx, CodeCacheKey(code)).Result()
}

func (c RedisEmailCodeCache) Set(ctx context.Context, code string, email string, expire time.Duration) error {
	return c.cache.Set(ctx, CodeCacheKey(code), email, expire).Err()
}

func (c RedisEmailCodeCache) Del(ctx context.Context, code string) error {
	return c.cache.Del(ctx, CodeCacheKey(code)).Err()
}

// Check get and remove
func (c RedisEmailCodeCache) Check(ctx context.Context, code string) (string, error) {
	cache, err := c.Get(ctx, code)
	if err != nil {
		return "", err
	}
	return cache, c.Del(ctx, code)
}
