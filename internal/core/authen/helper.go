package authen

import (
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/gin-gonic/gin"
)

const (
	contextTokenInfo = "gin.context.user.info"

	contextKeyInfo = "gin.context.key.info"
)

// GetContextTokenInfo returns the token info in current request
func GetContextTokenInfo(ctx *gin.Context) UserClaims {
	value, _ := ctx.Get(contextTokenInfo)
	claims, _ := value.(UserClaims)
	return claims
}

func SetContextTokenInfo(ctx *gin.Context, claims UserClaims) {
	ctx.Set(contextTokenInfo, claims)
}

// GetContextKeyInfo returns the key info in current request
func GetContextKeyInfo(ctx *gin.Context) entity.APIKey {
	value, _ := ctx.Get(contextKeyInfo)
	key, _ := value.(entity.APIKey)
	return key
}

func SetContextAPIInfo(ctx *gin.Context, key entity.APIKey) {
	ctx.Set(contextKeyInfo, key)
}
