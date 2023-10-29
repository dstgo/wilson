package system

import (
	"context"
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/core/role"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/cache"
	"github.com/dstgo/wilson/internal/types/auth"
	"github.com/dstgo/wilson/internal/types/system"
)

func NewAPIKey(source *data.DataSource) ApiKey {
	keyCache := cache.NewAPIKeyCache(source.Redis())
	resolver := role.NewGormResolver(source.ORM())
	return ApiKey{
		KeyCache:  keyCache,
		keyIssuer: authen.NewAPIKeyCacheAuthor(source, resolver, keyCache),
	}
}

type ApiKey struct {
	KeyCache  cache.KeyCache
	keyIssuer authen.KeyIssuer
}

func (a ApiKey) CreateAPiKey(ctx context.Context, option auth.CreateKeyOption) error {
	_, err := a.keyIssuer.Issue(ctx, option.Uid, option.Name, option.ExpiredAt, option.Perms)
	return err
}

func (a ApiKey) RemoveApiKey(ctx context.Context, userId, KeyId string) error {
	err := a.KeyCache.Del(ctx, userId, KeyId)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func (a ApiKey) ListApiKey(ctx context.Context, userId string) ([]auth.APIKey, error) {
	var keys []auth.APIKey

	list, err := a.KeyCache.List(ctx, userId)
	if err != nil {
		return keys, system.ErrDatabase.Wrap(err)
	}

	for _, key := range list {
		keys = append(keys, auth.APIKey{
			Name:      key.Name,
			Key:       key.Key,
			ExpiredAt: key.ExpiredAt,
		})
	}

	return keys, nil
}
