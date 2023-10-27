package authen

import "github.com/gin-gonic/gin"

const (
	contextUserInfo = "gin.context.user.info"
)

func SetContextTokenInfo(ctx *gin.Context, claims UserClaims) {
	ctx.Set(contextUserInfo, claims)
}

func GetContextTokenInfo(ctx *gin.Context) UserClaims {
	var userClaims UserClaims
	value, exists := ctx.Get(contextUserInfo)
	if !exists {
		return userClaims
	}
	if claims, ok := value.(UserClaims); ok {
		userClaims = claims
	}
	return userClaims
}
