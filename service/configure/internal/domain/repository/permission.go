package repository

import "github.com/dstgo/wilson/framework/kratosx"

type Permission interface {
	//GetPermission 获取当前用户，指定key的权限
	GetPermission(ctx kratosx.Context, keyword string) (bool, []uint32, error)

	// GetEnv 获取当前用户对于env的权限
	GetEnv(ctx kratosx.Context) (bool, []uint32, error)

	// HasEnv 获取当前用户是否具有指定env的权限
	HasEnv(ctx kratosx.Context, id uint32) bool

	// GetServer 获取当前用户是对于server的权限
	GetServer(ctx kratosx.Context) (bool, []uint32, error)

	// HasServer 获取当前用户是具有指定server的权限
	HasServer(ctx kratosx.Context, id uint32) bool
}
