package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/dstgo/wilson/internal/data"
	"github.com/go-redis/redis/v8"
	"time"
)

type TokenCache interface {
	TTL(ctx context.Context, tokenId string) (time.Duration, error)
	Key(tokenId string) string
	Get(ctx context.Context, tokenId string) (string, bool, error)
	Set(ctx context.Context, tokenId string, token string, expire time.Duration) error
	Del(ctx context.Context, tokenId string) error
}

func NewAccessTokenCache(d *data.DataSource) RedisTokenCache {
	return RedisTokenCache{cache: d.Redis(), keyFn: AccessTokenCacheKey}
}

func NewRefreshTokenCache(d *data.DataSource) RedisTokenCache {
	return RedisTokenCache{cache: d.Redis(), keyFn: RefreshTokenCacheKey}
}

var tokenCache TokenCache = RedisTokenCache{}

type RedisTokenCache struct {
	cache *redis.Client
	keyFn func(key string) string
}

func (t RedisTokenCache) TTL(ctx context.Context, tokenId string) (time.Duration, error) {
	ttl := t.cache.TTL(ctx, t.keyFn(tokenId))
	return ttl.Val(), ttl.Err()
}

func (t RedisTokenCache) Key(tokenId string) string {
	if t.keyFn != nil {
		return t.keyFn(tokenId)
	}
	return ""
}

func (t RedisTokenCache) Get(ctx context.Context, tokenId string) (string, bool, error) {
	var (
		res   string
		err   error
		exist bool
	)
	op := t.cache.Get(ctx, t.keyFn(tokenId))

	if errors.Is(op.Err(), redis.Nil) {
		exist = false
	} else if op.Err() != nil {
		err = op.Err()
	} else {
		res = op.String()
		exist = true
	}

	return res, exist, err
}

func (t RedisTokenCache) Set(ctx context.Context, tokenId string, token string, expire time.Duration) error {
	return t.cache.Set(ctx, t.keyFn(tokenId), token, expire).Err()
}

func (t RedisTokenCache) Del(ctx context.Context, tokenId string) error {
	return t.cache.Del(ctx, t.keyFn(tokenId)).Err()
}

func AccessTokenCacheKey(tokenId string) string {
	return fmt.Sprintf("token:%s", tokenId)
}

func RefreshTokenCacheKey(tokenId string) string {
	return fmt.Sprintf("refresh:%s", tokenId)
}
