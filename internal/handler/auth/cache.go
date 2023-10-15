package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dstgo/wilson/internal/data"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"time"
)

type TokenCache interface {
	Get(ctx context.Context, tokenId string) (string, bool, error)
	Set(ctx context.Context, tokenId string, token string, expire time.Duration) error
	Del(ctx context.Context, tokenId string) error
}

func TokenCacheKey(tokenId string) string {
	return fmt.Sprintf("token:%s", tokenId)
}

var TokenCacheProviderSet = wire.NewSet(
	NewTokenRedisCache,
	wire.Bind(new(TokenCache), new(RedisTokenCache)),
)

func NewTokenRedisCache(d *data.DataSource) RedisTokenCache {
	return RedisTokenCache{cache: d.Redis()}
}

type RedisTokenCache struct {
	cache *redis.Client
}

func (t RedisTokenCache) Get(ctx context.Context, tokenId string) (string, bool, error) {
	var (
		res   string
		err   error
		exist bool
	)
	op := t.cache.Get(ctx, TokenCacheKey(tokenId))

	if !errors.Is(op.Err(), redis.Nil) {
		err = op.Err()
	}

	return res, exist, err
}

func (t RedisTokenCache) Set(ctx context.Context, tokenId string, token string, expire time.Duration) error {
	return t.cache.Set(ctx, TokenCacheKey(tokenId), token, expire).Err()
}

func (t RedisTokenCache) Del(ctx context.Context, tokenId string) error {
	return t.cache.Del(ctx, TokenCacheKey(tokenId)).Err()
}
