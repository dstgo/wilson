package manager

import (
	"sync"

	"github.com/samber/lo"

	v1 "github.com/dstgo/wilson/api/gen/manager/auth/v1"
	mresourcev1 "github.com/dstgo/wilson/api/gen/manager/resource/v1"
	"github.com/dstgo/wilson/framework/kratosx"
)

const (
	Manager = "Manager"
	Env     = "cfg_env"
	Server  = "cfg_server"
)

type Permission struct {
}

var (
	permissionIns  *Permission
	permissionOnce sync.Once
)

func NewPermission() *Permission {
	permissionOnce.Do(func() {
		permissionIns = &Permission{}
	})
	return permissionIns
}

// GetPermission 获取当前用户，指定key的权限
func (p *Permission) GetPermission(ctx kratosx.Context, keyword string) (bool, []uint32, error) {
	var (
		info = &v1.AuthReply{}
		err  error
	)
	if ctx.Token() != "" {
		err = ctx.JWT().Parse(ctx, info)
	} else {
		err = ctx.Authentication().ParseAuthFromMD(ctx, &info)
	}

	if err != nil {
		return false, nil, err
	}
	if info.UserId == 1 || info.RoleId == 1 {
		return true, nil, nil
	}

	client, err := mResourceClient(ctx)
	if err != nil {
		return false, nil, err
	}

	mdCtx, err := ctx.Authentication().SetAuthMD(ctx, info)
	if err != nil {
		return false, nil, err
	}

	reply, err := client.GetResourceScopes(mdCtx, &mresourcev1.GetResourceScopesRequest{
		Keyword: keyword,
	})

	if err != nil {
		return false, nil, err
	}
	return reply.All, reply.Scopes, nil
}

// GetEnv 获取当前用户对于env的权限
func (p *Permission) GetEnv(ctx kratosx.Context) (bool, []uint32, error) {
	all, ids, err := p.GetPermission(ctx, Env)
	if ids == nil {
		ids = []uint32{}
	}
	return all, ids, err
}

// HasEnv 获取当前用户是否具有指定env的权限
func (p *Permission) HasEnv(ctx kratosx.Context, id uint32) bool {
	all, ids, err := p.GetPermission(ctx, Env)
	if err != nil {
		return false
	}
	if all {
		return true
	}
	return lo.Contains(ids, id)
}

// GetServer 获取当前用户是对于server的权限
func (p *Permission) GetServer(ctx kratosx.Context) (bool, []uint32, error) {
	all, ids, err := p.GetPermission(ctx, Server)
	if ids == nil {
		ids = []uint32{}
	}
	return all, ids, err
}

// HasServer 获取当前用户是具有指定server的权限
func (p *Permission) HasServer(ctx kratosx.Context, id uint32) bool {
	all, ids, err := p.GetPermission(ctx, Server)
	if err != nil {
		return false
	}
	if all {
		return true
	}
	return lo.Contains(ids, id)
}
