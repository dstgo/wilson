package authen

import (
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/gin-gonic/gin"
)

const (
	contextTokenInfo = "gin.context.user.info"

	contextKeyInfo = "gin.context.key.info"
)

func SetContextAPIInfo(ctx *gin.Context, key entity.APIKey) {
	ctx.Set(contextTokenInfo, key)
}

func GetContextTokenInfo(ctx *gin.Context) UserClaims {
	value, _ := ctx.Get(contextTokenInfo)
	claims, _ := value.(UserClaims)
	return claims
}

func SetContextTokenInfo(ctx *gin.Context, claims UserClaims) {
	ctx.Set(contextTokenInfo, claims)
}

func GetContextKeyInfo(ctx *gin.Context) entity.APIKey {
	value, _ := ctx.Get(contextTokenInfo)
	key, _ := value.(entity.APIKey)
	return key
}
