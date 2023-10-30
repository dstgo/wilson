package authen

import (
	"context"
	"errors"
	"github.com/dstgo/wilson/internal/core/role"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/cache"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/auth"
	roleType "github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/google/uuid"
	"slices"
	"time"
)

type KeyAuthor interface {
	Authenticate(ctx context.Context, key, obj, act string) (entity.APIKey, error)
}

type KeyIssuer interface {
	Issue(ctx context.Context, userId string, name string, exp int64, perms []uint) (entity.APIKey, error)
}

func NewAPIKeyCacheAuthor(ds *data.DataSource, resolver role.Resolver, cache cache.KeyCache) APIKeyCacheAuthor {
	return APIKeyCacheAuthor{
		ds:       ds,
		cache:    cache,
		resolver: resolver,
	}
}

type APIKeyCacheAuthor struct {
	ds       *data.DataSource
	cache    cache.KeyCache
	resolver role.Resolver
}

func (a APIKeyCacheAuthor) Issue(ctx context.Context, userId string, name string, exp int64, perms []uint) (entity.APIKey, error) {

	// check if permission exists
	permList, err := a.resolver.GetPermInBatch(perms, system.OpenAPI)
	if err != nil {
		return entity.APIKey{}, err
	} else if len(permList) < len(perms) {
		return entity.APIKey{}, roleType.ErrPermNotFound
	}

	var (
		key = cryptor.Sha256WithBase64(uuid.NewString())
		now = time.Now()
		ttl = time.Unix(0, exp).Sub(now)
	)

	apikey := entity.APIKey{
		Key:       key,
		UserId:    userId,
		Name:      name,
		Perms:     perms,
		CreatedAt: uint64(now.Nanosecond()),
		ExpiredAt: uint64(exp),
	}

	if apikey.ExpiredAt <= 0 || ttl <= 0 {
		return apikey, auth.ErrInvalidKeyExpration
	}

	err = a.cache.Set(ctx, apikey, ttl)
	if err != nil {
		return apikey, err
	}

	return apikey, nil
}

func (a APIKeyCacheAuthor) Authenticate(ctx context.Context, key, obj, act string) (entity.APIKey, error) {
	if len(key) == 0 {
		return entity.APIKey{}, auth.ErrInvalidKey
	}

	apikey, b, err := a.cache.Find(ctx, key)
	if err == nil && !b || errors.Is(err, cache.ErrDuplicateKey) {
		return apikey, auth.ErrInvalidKey
	} else if err != nil {
		return apikey, err
	}
	perm, err := a.resolver.MatchPerm("", obj, act, "", system.OpenAPI)
	if err != nil {
		return apikey, system.ErrDatabase.Wrap(err)
	}

	if !slices.Contains(apikey.Perms, perm.Id) {
		return apikey, auth.ErrKeyNoPerm
	}

	return apikey, nil
}
