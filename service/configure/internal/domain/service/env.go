package service

import (
	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/api/gen/errors"

	"github.com/dstgo/wilson/service/configure/internal/conf"
	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/domain/repository"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Env struct {
	conf       *conf.Config
	repo       repository.Env
	permission repository.Permission
}

func NewEnv(
	conf *conf.Config,
	repo repository.Env,
	permission repository.Permission,
) *Env {
	return &Env{
		conf:       conf,
		repo:       repo,
		permission: permission,
	}
}

// ListEnv 获取环境信息列表
func (u *Env) ListEnv(ctx kratosx.Context, req *types.ListEnvRequest) ([]*entity.Env, uint32, error) {
	// 获取具有权限的环境id列表
	all, scopes, err := u.permission.GetEnv(ctx)
	if err != nil {
		return nil, 0, err
	}
	if !all {
		req.Ids = scopes
	}

	// 获取环境列表
	list, total, err := u.repo.ListEnv(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// GetEnvToken 获取环境token
func (u *Env) GetEnvToken(ctx kratosx.Context, id uint32) (string, error) {
	// 环境鉴权
	if !u.permission.HasEnv(ctx, id) {
		return "", errors.NotPermissionError()
	}

	// 获取指定环境
	env, err := u.repo.GetEnv(ctx, id)
	if err != nil {
		return "", errors.CreateError(err.Error())
	}
	return env.Token, nil
}

// CreateEnv 创建环境信息
func (u *Env) CreateEnv(ctx kratosx.Context, req *entity.Env) (uint32, error) {
	id, err := u.repo.CreateEnv(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateEnv 更新环境信息
func (u *Env) UpdateEnv(ctx kratosx.Context, req *entity.Env) error {
	// 环境鉴权
	if !u.permission.HasEnv(ctx, req.Id) {
		return errors.NotPermissionError()
	}
	// 更新环境
	if err := u.repo.UpdateEnv(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteEnv 删除环境信息
func (u *Env) DeleteEnv(ctx kratosx.Context, id uint32) error {
	// 环境鉴权
	if !u.permission.HasEnv(ctx, id) {
		return errors.NotPermissionError()
	}
	// 删除环境
	if err := u.repo.DeleteEnv(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
