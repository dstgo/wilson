package email

import (
	"context"
	"fmt"
	"github.com/dstgo/wilson/internal/data"
	"github.com/go-redis/redis/v8"
	"time"
)

func NewEmailCodeCache(source *data.DataSource) CodeCache {
	return CodeCache{
		cache: source.Redis(),
	}
}

func CodeCacheKey(code string) string {
	return fmt.Sprintf("email:code:%s", code)
}

type CodeCache struct {
	cache *redis.Client
}

func (c CodeCache) Get(ctx context.Context, code string) (string, error) {
	return c.cache.Get(ctx, CodeCacheKey(code)).Result()
}

func (c CodeCache) Set(ctx context.Context, code string, email string, expire time.Duration) error {
	return c.cache.Set(ctx, CodeCacheKey(code), email, expire).Err()
}

func (c CodeCache) Del(ctx context.Context, code string) error {
	return c.cache.Del(ctx, CodeCacheKey(code)).Err()
}

// Check get and remove
func (c CodeCache) Check(ctx context.Context, code string) (string, error) {
	cache, err := c.Get(ctx, code)
	if err != nil {
		return "", err
	}
	return cache, c.Del(ctx, code)
}
