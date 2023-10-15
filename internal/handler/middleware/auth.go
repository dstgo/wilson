package middleware

import (
	"github.com/dstgo/wilson/internal/handler/auth"
	"github.com/dstgo/wilson/internal/pkg/httpx"
	resp2 "github.com/dstgo/wilson/internal/pkg/resp"
	"github.com/dstgo/wilson/internal/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"net/http"
)

func UseAuthenticate(v auth.Parser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		routeMeta := route.MetaFromCtx(ctx)

		// if no need to auth, skip
		if routeMeta.Has(meta.NoAuth.Key) {
			ctx.Next()
			return
		}

		token := httpx.GetBearerTokenFromCtx(ctx)
		jwtToken, err := v.Parse(ctx, token)
		if err == nil {
			var userClaims auth.UserClaims
			if res, e := jwtToken.Claims.(*auth.UserClaims); e {
				userClaims = *res
			}
			auth.SetContextTokenInfo(ctx, userClaims)
			ctx.Next()
		} else {
			ctx.Abort()
			var respErr *resp2.ResponseError
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				respErr = resp2.NewErr().Status(http.StatusUnauthorized).I18n("jwt.expired").Err(err)
			default:
				respErr = resp2.NewErr().Status(http.StatusForbidden).I18n("jwt.parsedFailed").Err(err)
			}
			resp2.Fail(ctx).Code(respErr.HttpStatus * 10).MsgI18n("error.forbidden").Error(respErr).Send()
		}
	}
}

func UseCasbin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UseApiKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
