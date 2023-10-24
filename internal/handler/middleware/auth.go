package middleware

import (
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/core/role"
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/dstgo/wilson/internal/types/auth"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/dstgo/wilson/internal/types/meta"
	roleType "github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/dstgo/wilson/pkg/ginx"
	"github.com/dstgo/wilson/pkg/ginx/httpx"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"net/http"
)

func UseAuthenticate(v authen.Parser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		routeMeta := ginx.MetaFromCtx(ctx)

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
			var respErr errs.LocaleError
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				respErr = auth.ErrJwtExpired.Wrap(err).Status(http.StatusUnauthorized)
			default:
				respErr = auth.ErrJwtParsedFailed.Wrap(err).Status(http.StatusForbidden)
			}
			resp.Fail(ctx).MsgI18n("err.forbidden").Error(respErr).Send()
		}
	}
}

func UseRoleAuthorize(resolver role.Resolver, userRole user.UserInfo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// no need to authorize if no need to authenticate
		metaData := ginx.MetaFromCtx(ctx)
		if metaData.Has(meta.NoAuth.Key) || metaData.Has(meta.Anonymous.Key) {
			ctx.Next()
			return
		}

		info := authen.GetContextTokenInfo(ctx)
		userRoles, err := userRole.GetUserRoleCodes(info.UUID)
		if err != nil {
			ctx.Abort()
			resp.InternalFailed(ctx).MsgI18n("err.internal").Error(err).Send()
			return
		}

		var (
			roles  = userRoles
			object = ctx.FullPath()
			action = ctx.Request.Method
		)

		if err := resolver.ResolveAny(object, action, roles...); err != nil {
			ctx.Abort()
			if errors.Is(err, role.ErrHasNoPermission) {
				resp.Forbidden(ctx).MsgI18n("err.unauthorized").Error(roleType.ErrNoPemrAccess).Send()
			} else {
				resp.InternalFailed(ctx).MsgI18n("err.internal").Error(system.ErrDatabase.Wrap(err)).Send()
			}
			return
		}

		ctx.Next()
	}
}
