package repository

import (
	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Env interface {
	// GetEnvByToken 获取指定的环境信息
	GetEnvByToken(ctx kratosx.Context, token string) (*entity.Env, error)

	// GetEnvByKeyword 获取指定的环境信息
	GetEnvByKeyword(ctx kratosx.Context, keyword string) (*entity.Env, error)

	// GetEnv 获取指定的环境信息
	GetEnv(ctx kratosx.Context, id uint32) (*entity.Env, error)

	// ListEnv 获取环境信息列表
	ListEnv(ctx kratosx.Context, req *types.ListEnvRequest) ([]*entity.Env, uint32, error)

	// CreateEnv 创建环境信息
	CreateEnv(ctx kratosx.Context, env *entity.Env) (uint32, error)

	// UpdateEnv 更新环境信息
	UpdateEnv(ctx kratosx.Context, env *entity.Env) error

	// DeleteEnv 删除环境信息
	DeleteEnv(ctx kratosx.Context, ids uint32) error
}
