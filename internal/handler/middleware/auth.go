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
	"github.com/pkg/errors"
	"net/http"
)

func UseAuthenticate(v authen.TokenParser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		routeMeta := ginx.MetaFromCtx(ctx)

		// if no need to auth, skip
		if routeMeta.Has(meta.NoAuth.Key) {
			ctx.Next()
			return
		}

		token := httpx.GetBearerTokenFromCtx(ctx)
		parsedToken, err := v.Parse(ctx, token)
		if err == nil {
			authen.SetContextTokenInfo(ctx, parsedToken.Access.Payload)
			ctx.Next()
		} else {
			ctx.Abort()
			var respErr errs.LocaleError
			switch {
			case errors.Is(err, authen.ErrTokenExpired):
				respErr = auth.ErrTokenExpired.Wrap(err)
			case errors.Is(err, authen.ErrTokenNeedRefreshed):
				respErr = auth.ErrTokenNeedRefresh.Wrap(err)
			default:
				respErr = auth.ErrTokenParsedFailed.Wrap(err)
			}
			resp.Fail(ctx).Status(http.StatusUnauthorized).MsgI18n("err.unauthorized").Error(respErr).Transparent().Send()
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
				resp.Forbidden(ctx).MsgI18n("err.forbidden").Error(roleType.ErrPermNoAccess).Send()
			} else {
				resp.InternalFailed(ctx).MsgI18n("err.internal").Error(system.ErrDatabase.Wrap(err)).Transparent().Send()
			}
			return
		}

		ctx.Next()
	}
}

func UseOpenAPIAuth(author authen.KeyAuthor) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			obj = ctx.FullPath()
			act = ctx.Request.Method
			key = ctx.Query("key")
		)
		apikey, err := author.Authenticate(ctx, key, obj, act)
		if err == nil {
			authen.SetContextAPIInfo(ctx, apikey)
			ctx.Next()
		} else {
			ctx.Abort()
			switch {
			case errors.Is(err, system.ErrDatabase):
				resp.InternalFailed(ctx).MsgI18n("err.internal").Error(err).Transparent().Send()
			default:
				resp.Forbidden(ctx).MsgI18n("err.forbidden").Error(err).Transparent().Send()
			}
		}
	}
}
