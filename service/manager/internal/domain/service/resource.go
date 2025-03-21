package service

import (
	"github.com/dstgo/wilson/api/gen/errors"
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/kratosx/library/md"
	"github.com/dstgo/wilson/framework/pkg/slicex"
	"github.com/dstgo/wilson/service/manager/internal/conf"
	"github.com/dstgo/wilson/service/manager/internal/domain/entity"
	"github.com/dstgo/wilson/service/manager/internal/domain/repository"
	"github.com/dstgo/wilson/service/manager/internal/types"
)

type Resource struct {
	conf *conf.Config
	repo repository.Resource
	dept repository.Department
}

func NewResource(config *conf.Config,
	repo repository.Resource,
	dept repository.Department,
) *Resource {
	return &Resource{conf: config, repo: repo, dept: dept}
}

// GetResourceScopes 获取指定的资源权限
func (u *Resource) GetResourceScopes(ctx kratosx.Context, keyword string) (bool, []uint32, error) {
	// 获取用户当前的部门权限
	all, scopes, err := u.dept.GetDepartmentDataScope(ctx, md.UserId(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "get resource scopes error", "err", err.Error())
		return false, nil, errors.DatabaseError()
	}
	if all {
		return true, nil, nil
	}

	// 获取准许部门的资源列表
	ids, err := u.repo.GetResourceScopes(ctx, &types.GetResourceScopesRequest{
		Keyword:       keyword,
		DepartmentIds: scopes,
	})
	if err != nil {
		return false, nil, err
	}

	return false, ids, nil
}

// GetResource 获取指定的资源权限
func (u *Resource) GetResource(ctx kratosx.Context, req *types.GetResourceRequest) ([]uint32, error) {
	// 获取用户当前的部门权限
	all, scopes, err := u.dept.GetDepartmentDataScope(ctx, md.UserId(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "get resource scopes error", "err", err.Error())
		return nil, errors.DatabaseError()
	}
	if !all {
		req.DepartmentIds = scopes
	}

	// 获取资源的部门
	ids, err := u.repo.GetResource(ctx, req)
	if err != nil {
		return nil, errors.DatabaseErrorWrap(err)
	}
	return ids, nil
}

// UpdateResource 更新资源权限
func (u *Resource) UpdateResource(ctx kratosx.Context, req *types.UpdateResourceRequest) error {
	// 获取用户当前的部门权限
	all, scopes, err := u.dept.GetDepartmentDataScope(ctx, md.UserId(ctx))
	if err != nil {
		ctx.Logger().Warnw("msg", "update resource scopes error", "err", err.Error())
		return errors.DatabaseError()
	}
	if !all {
		req.DepartmentIds = scopes
	}

	var (
		list []*entity.Resource
		set  = slicex.ToBoolSetOrdered(scopes)
	)
	for _, id := range req.DepartmentIds {
		// 过滤管理权限外的部门
		if all || set[id] {
			list = append(list, &entity.Resource{
				Keyword:      req.Keyword,
				ResourceId:   req.ResourceId,
				DepartmentId: id,
			})
		}
	}

	if err := ctx.Transaction(func(ctx kratosx.Context) error {
		// 删除资源权限
		delReq := &types.DeleteResourceRequest{
			ResourceId: req.ResourceId,
			Keyword:    req.Keyword,
		}
		if !all {
			delReq.DepartmentIds = scopes
		}
		if err := u.repo.DeleteResource(ctx, delReq); err != nil {
			return err
		}

		// 设置新的资源权限
		return u.repo.CreateResources(ctx, list)
	}); err != nil {
		return errors.UpdateErrorWrap(err)
	}
	return nil
}
