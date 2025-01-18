package service

import (
	"github.com/samber/lo"

	"github.com/dstgo/wilson/api/gen/errors"
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/kratosx/library/md"
	"github.com/dstgo/wilson/service/manager/internal/conf"
	"github.com/dstgo/wilson/service/manager/internal/types"
)

type Auth struct {
	conf *conf.Config
}

func NewAuth(conf *conf.Config) *Auth {
	return &Auth{
		conf: conf,
	}
}

// Auth 外部接口鉴权
func (u *Auth) Auth(ctx kratosx.Context, in *types.AuthRequest) (*md.Auth, error) {
	info := md.GetAuthInfo(ctx)

	if lo.Contains(ctx.Config().App().Authentication.SkipRole, info.RoleKeyword) {
		return info, nil
	}

	author := ctx.Authentication()
	if author.IsWhitelist(in.Path, in.Method) {
		return info, nil
	}

	enforce := ctx.Authentication().Enforce()
	isAuth, _ := enforce.Enforce(info.RoleKeyword, in.Path, in.Method)
	if !isAuth {
		return nil, errors.ForbiddenError()
	}

	return info, nil
}
