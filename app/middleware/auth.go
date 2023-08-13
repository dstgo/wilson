package middleware

import (
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func UserHs256Jwt(v auth.Authenticator, l *locale.Locale) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := httpx.GetBearerTokenFromCtx(ctx)
		jwtToken, err := v.Authenticate(ctx, token)
		if err == nil {
			ctx.Next()
			var userClaims auth.UserClaims
			if res, e := jwtToken.Claims.(auth.UserClaims); e {
				userClaims = res
			}
			auth.SetContextUserInfo(ctx, userClaims)
		} else {
			ctx.Abort()
			var httpError httpx.Error
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				httpError = httpx.NewErrorMsg(401, l.GetWithCtx(ctx, "jwt.expired"))
			default:
				httpError = httpx.NewErrorMsg(403, l.GetWithCtx(ctx, "jwt.parsedFailed"))
			}
			httpx.Failed(ctx, httpError.Code*10, httpError)
		}
	}
}

func UseCasbin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
