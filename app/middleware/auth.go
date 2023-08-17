package middleware

import (
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/dstgo/wilson/app/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func UseJwtAuthenticate(v auth.Authenticator, l *locale.Locale) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		routeMeta := route.MetaFromCtx(ctx)

		// if no need to auth, skip
		if routeMeta.Has(meta.NoAuth.Key) {
			ctx.Next()
			return
		}

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
			var httpError resp.Error
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				httpError = resp.NewI18nErr(401, "jwt.expired")
			default:
				httpError = resp.NewI18nErr(403, "jwt.parsedFailed")
			}
			resp.Fail(ctx, httpError.Code*10, httpError)
		}
	}
}

func UseCasbin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
