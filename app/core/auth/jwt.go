package auth

import (
	"context"
	"fmt"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/pkg/jwtx"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

// Authenticator
// The authenticator should verify whether the request has been authenticated.
type Authenticator interface {
	Authenticate(ctx context.Context, token string) (jwtx.Jwt, error)
}

// Issuer
// The Issuer should issue a new jwt token and return the token info
type Issuer interface {
	Issue(ctx context.Context, user UserPayload, exp time.Duration) (jwtx.Jwt, error)
}

// UserPayload
// basic user info
type UserPayload struct {
	Username string `json:"Username"`
	UserId   string `json:"userId"`
}

var (
	ContextUserInfo = "gin.context.user.info"
)

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

type UserClaims struct {
	UserPayload
	jwt.RegisteredClaims
}

func NewJwtAuthenticator(cfg *conf.JwtConf, client *redis.Client) *JwtAuthenticator {
	return &JwtAuthenticator{
		redis:  client,
		cfg:    cfg,
		method: jwt.SigningMethodHS256,
	}
}

type JwtAuthenticator struct {
	redis  *redis.Client
	cfg    *conf.JwtConf
	method jwt.SigningMethod
}

func (j *JwtAuthenticator) Issue(ctx context.Context, user UserPayload, exp time.Duration) (jwtx.Jwt, error) {
	now := time.Now()
	expiredAt := now.Add(exp)
	if exp <= 0 {
		expiredAt = now.Add(j.cfg.Exp)
	}

	userClaims := UserClaims{
		UserPayload: user,
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

	// 将token的唯一ID存入redis
	if err = j.redis.Set(ctx, fmt.Sprintf("token:%s", userClaims.ID), newJwt.SignedJwt, j.cfg.Exp).Err(); err != nil {
		return jwtx.Jwt{}, err
	}

	return newJwt, nil
}

func (j *JwtAuthenticator) Authenticate(ctx context.Context, token string) (jwtx.Jwt, error) {
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

	// search from redis
	if res, err := j.redis.Get(ctx, fmt.Sprintf("token:%s", userClaims.ID)).Result(); err != nil {
		return jwtV, err
	} else if len(res) == 0 {
		return jwtV, errors.New("token not found")
	}

	return parsedJwt, nil
}
