package middleware

import (
	"github.com/dstgo/wilson/internal/pkg/httpx"
	"github.com/dstgo/wilson/internal/sys/authenticate"
	"github.com/dstgo/wilson/internal/sys/resp"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/dstgo/wilson/internal/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func UseAuthenticate(v authenticate.Parser) gin.HandlerFunc {
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
			var userClaims authenticate.UserClaims
			if res, e := jwtToken.Claims.(*authenticate.UserClaims); e {
				userClaims = *res
			}
			authenticate.SetContextTokenInfo(ctx, userClaims)
			ctx.Next()
		} else {
			ctx.Abort()
			var respErr *errs.ResponseError
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				respErr = errs.UnAuthorized(err).I18n("jwt.expired")
			default:
				respErr = errs.Forbidden(err).I18n("jwt.parsedFailed")
			}
			resp.Fail(ctx).Code(respErr.HttpStatus * 10).MsgI18n("error.forbidden").Error(respErr).Send()
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
