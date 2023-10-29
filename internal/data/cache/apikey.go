package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ErrDuplicateKey = errors.New("duplicate key")
)

type KeyCache interface {
	Key(userId, keyId string) string
	TTL(ctx context.Context, userId, keyId string) (time.Duration, error)
	Get(ctx context.Context, userId, keyId string) (entity.APIKey, bool, error)
	Find(ctx context.Context, keyId string) (entity.APIKey, bool, error)
	Set(ctx context.Context, apikey entity.APIKey, ttl time.Duration) error
	Del(ctx context.Context, userId, keyId string) error
	List(ctx context.Context, userId string) ([]entity.APIKey, error)
}

var redisApiKeyCache = KeyCache(RedisApiKeyCache{})

func APIKeyCacheKey(userId, keyId string) string {
	return fmt.Sprintf("apikey:%s:%s", userId, keyId)
}

func NewAPIKeyCache(client *redis.Client) RedisApiKeyCache {
	return RedisApiKeyCache{cache: client, keyFn: APIKeyCacheKey}
}

type RedisApiKeyCache struct {
	cache *redis.Client
	keyFn func(keyId, userId string) string
}

func (r RedisApiKeyCache) unmarshal(b []byte) (entity.APIKey, error) {
	var apikey entity.APIKey
	err := json.Unmarshal(b, &apikey)
	if err != nil {
		return apikey, err
	}
	return apikey, nil
}

func (r RedisApiKeyCache) marshal(key entity.APIKey) ([]byte, error) {
	marshal, err := json.Marshal(key)
	if err != nil {
		return []byte{}, err
	}
	return marshal, nil
}

func (r RedisApiKeyCache) Key(userId, keyId string) string {
	return r.keyFn(userId, keyId)
}

func (r RedisApiKeyCache) TTL(ctx context.Context, userId, keyId string) (time.Duration, error) {
	key := r.Key(userId, keyId)
	ttl := r.cache.TTL(ctx, key)
	return ttl.Val(), ttl.Err()
}

func (r RedisApiKeyCache) get(ctx context.Context, key string) (entity.APIKey, bool, error) {
	var apikey entity.APIKey
	bytes, err := r.cache.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return apikey, false, nil
	} else if err != nil {
		return apikey, false, err
	}

	unmarshal, err := r.unmarshal(bytes)
	if err != nil {
		return apikey, false, err
	}
	apikey = unmarshal
	return apikey, true, nil
}

func (r RedisApiKeyCache) Get(ctx context.Context, userId, keyId string) (entity.APIKey, bool, error) {
	return r.get(ctx, r.Key(userId, keyId))
}

func (r RedisApiKeyCache) Find(ctx context.Context, keyId string) (entity.APIKey, bool, error) {
	scan, err := r.scan(ctx, r.Key("*", keyId))
	if err != nil {
		return entity.APIKey{}, false, err
	}
	if len(scan) == 0 {
		return entity.APIKey{}, false, nil
	} else if len(scan) == 1 {
		return scan[0], true, nil
	} else {
		return entity.APIKey{}, false, ErrDuplicateKey
	}
}

func (r RedisApiKeyCache) Set(ctx context.Context, apikey entity.APIKey, ttl time.Duration) error {
	key := r.Key(apikey.UserId, apikey.Key)
	marshal, err := r.marshal(apikey)
	if err != nil {
		return err
	}
	return r.cache.Set(ctx, key, marshal, ttl).Err()
}

func (r RedisApiKeyCache) Del(ctx context.Context, userId, keyId string) error {
	key := r.Key(userId, keyId)
	return r.cache.Del(ctx, key).Err()
}

func (r RedisApiKeyCache) scan(ctx context.Context, pattern string) ([]entity.APIKey, error) {
	var (
		cursor  uint64
		match   = pattern
		count   = int64(20)
		apikeys = make([]entity.APIKey, 0, count*2)
		keys    []string
	)

	for {
		result := r.cache.Scan(ctx, cursor, match, count)
		keys, cursor = result.Val()

		for _, key := range keys {
			apikey, exist, err := r.get(ctx, key)
			if err != nil {
				return apikeys, err
			} else if !exist {
				continue
			}
			apikeys = append(apikeys, apikey)
		}

		if cursor == 0 {
			break
		}
	}

	return apikeys, nil
}

func (r RedisApiKeyCache) List(ctx context.Context, userId string) ([]entity.APIKey, error) {
	return r.scan(ctx, r.Key(userId, "*"))
}
