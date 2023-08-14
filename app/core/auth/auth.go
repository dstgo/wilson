package auth

import (
	"context"
	"github.com/dstgo/wilson/app/pkg/jwtx"
	"github.com/gin-gonic/gin"
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

func SetContextUserInfo(ctx *gin.Context, claims UserClaims) {
	ctx.Set(ContextUserInfo, claims)
}

func GetContextUserInfo(ctx *gin.Context) UserClaims {
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
