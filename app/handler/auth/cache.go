package auth

import (
	"context"
	"fmt"
	"github.com/dstgo/wilson/app/data"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"time"
)

type TokenCache interface {
	Get(ctx context.Context, tokenId string) (string, error)
	Set(ctx context.Context, tokenId string, token string, expire time.Duration) error
	Del(ctx context.Context, tokenId string) error
}

func TokenCacheKey(tokenId string) string {
	return fmt.Sprintf("token:%s", tokenId)
}

var TokenCacheProviderSet = wire.NewSet(
	NewTokenRedisCache,
	wire.Bind(new(TokenCache), new(TokenRedisCache)),
)

func NewTokenRedisCache(d *data.DataSource) TokenRedisCache {
	return TokenRedisCache{cache: d.Redis()}
}

type TokenRedisCache struct {
	cache *redis.Client
}

func (t TokenRedisCache) Get(ctx context.Context, tokenId string) (string, error) {
	return t.cache.Get(ctx, TokenCacheKey(tokenId)).Result()
}

func (t TokenRedisCache) Set(ctx context.Context, tokenId string, token string, expire time.Duration) error {
	return t.cache.Set(ctx, TokenCacheKey(tokenId), token, expire).Err()
}

func (t TokenRedisCache) Del(ctx context.Context, tokenId string) error {
	return t.cache.Del(ctx, TokenCacheKey(tokenId)).Err()
}
