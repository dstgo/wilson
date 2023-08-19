package middleware

import (
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/dstgo/wilson/app/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"net/http"
)

func UseJwtAuthenticate(v auth.Authenticator) gin.HandlerFunc {
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
			var userClaims auth.UserClaims
			if res, e := jwtToken.Claims.(*auth.UserClaims); e {
				userClaims = *res
			}
			auth.SetContextTokenInfo(ctx, userClaims)
			ctx.Next()
		} else {
			ctx.Abort()
			var respErr *resp.ResponseError
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				respErr = resp.NewErr().Status(http.StatusUnauthorized).I18n("jwt.expired").Err(err)
			default:
				respErr = resp.NewErr().Status(http.StatusForbidden).I18n("jwt.parsedFailed").Err(err)
			}
			resp.Fail(ctx).Code(respErr.HttpStatus * 10).MsgI18n("error.forbidden").Error(respErr).Send()
		}
	}
}

func UseCasbin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
