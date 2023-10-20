package middleware

import (
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/core/role"
	"github.com/dstgo/wilson/internal/pkg/httpx"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/dstgo/wilson/internal/types/meta"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func UseAuthenticate(v authen.Parser) gin.HandlerFunc {
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
			var userClaims authen.UserClaims
			if res, e := jwtToken.Claims.(*authen.UserClaims); e {
				userClaims = *res
			}
			authen.SetContextTokenInfo(ctx, userClaims)
			ctx.Next()
		} else {
			ctx.Abort()
			var respErr *errs.LocaleError
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				respErr = errs.UnAuthorized(err).I18n("jwt.expired")
			default:
				respErr = errs.Forbidden(err).I18n("jwt.parsedFailed")
			}
			resp.Fail(ctx).MsgI18n("err.forbidden").Error(respErr).Send()
		}
	}
}

func UseRoleAuthorize(resolver role.Resolver) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		info := authen.GetContextTokenInfo(ctx)
		var (
			roles  = info.Roles
			object = ctx.FullPath()
			action = ctx.Request.Method
		)

		err := resolver.ResolveAny(object, action, roles...)
		if err != nil {
			resp.Fail(ctx).MsgI18n("error.unauthorized").Error(errs.UnAuthorized(err)).Send()
			return
		}

		ctx.Next()
	}
}
