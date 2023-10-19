package authen

import (
	"context"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data/cache"
	"github.com/dstgo/wilson/internal/pkg/jwtx"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

// Parser
// The Parser should verify if the request has been authenticated.
type Parser interface {
	Parse(ctx context.Context, token string) (jwtx.Jwt, error)
}

// Issuer
// The Issuer should issue a new jwt token and return the token info
type Issuer interface {
	Issue(ctx context.Context, payload UserPayload, exp time.Duration) (jwtx.Jwt, error)
}

var (
	ContextUserInfo = "gin.context.user.info"
)

// UserPayload
// basic user info
type UserPayload struct {
	Username string   `json:"username"`
	UUID     string   `json:"uuid"`
	Roles    []string `json:"roles"`
}

type UserClaims struct {
	UserPayload
	jwt.RegisteredClaims
}

func SetContextTokenInfo(ctx *gin.Context, claims UserClaims) {
	ctx.Set(ContextUserInfo, claims)
}

func GetContextTokenInfo(ctx *gin.Context) UserClaims {
	var userClaims UserClaims
	value, exists := ctx.Get(ContextUserInfo)
	if !exists {
		return userClaims
	}
	if claims, ok := value.(UserClaims); ok {
		userClaims = claims
	}
	return userClaims
}

func NewCacheAuthor(cfg *conf.JwtConf, cache cache.TokenCache) *CacheAuthor {
	return &CacheAuthor{
		cache:  cache,
		cfg:    cfg,
		method: jwt.SigningMethodHS256,
	}
}

type CacheAuthor struct {
	cache  cache.TokenCache
	cfg    *conf.JwtConf
	method jwt.SigningMethod
}

func (j *CacheAuthor) Issue(ctx context.Context, payload UserPayload, exp time.Duration) (jwtx.Jwt, error) {
	now := time.Now()
	expiredAt := now.Add(exp)
	if exp <= 0 {
		expiredAt = now.Add(j.cfg.Exp)
	}

	userClaims := UserClaims{
		UserPayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.Isu,
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		},
	}

	newJwt, err := jwtx.NewJwt(j.cfg.Sig, j.method, userClaims)
	if err != nil {
		return jwtx.Jwt{}, err
	}

	// store jwt info cache
	if err := j.cache.Set(ctx, userClaims.ID, newJwt.SignedJwt, j.cfg.Exp); err != nil {
		return jwtx.Jwt{}, err
	}

	return newJwt, nil
}

func (j *CacheAuthor) Parse(ctx context.Context, token string) (jwtx.Jwt, error) {
	var (
		jwtV       jwtx.Jwt
		secret     = j.cfg.Sig
		method     = j.method
		userClaims UserClaims
	)

	// try to parse token
	parsedJwt, err := jwtx.ParseJwt(token, secret, method, &UserClaims{})
	if err != nil {
		return jwtV, err
	}

	if claims, e := parsedJwt.Claims.(*UserClaims); e {
		userClaims = *claims
	}

	// find jwt from cache
	_, e, err := j.cache.Get(ctx, userClaims.ID)
	if !e && err == nil {
		return jwtV, errors.New("token not found")
	} else if err != nil {
		return jwtV, err
	}

	return parsedJwt, nil
}
