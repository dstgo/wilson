package dbs

import (
	"sync"

	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/service/manager/internal/domain/entity"
	"github.com/dstgo/wilson/service/manager/internal/types"
)

type Rbac struct {
}

var (
	rbacIns  *Rbac
	rbacOnce sync.Once
)

func NewRbac() *Rbac {
	rbacOnce.Do(func() {
		rbacIns = &Rbac{}
	})
	return rbacIns
}

func (infra *Rbac) CreateRbacRolesApi(ctx kratosx.Context, roles []string, req types.MenuApi) error {
	defer func() {
		err := ctx.Authentication().Enforce().LoadPolicy()
		if err != nil {
			ctx.Logger().Errorf("CreateRbacRolesApi LoadPolicy() failed: %s", err.Error())
		}
	}()

	var list []*entity.CasbinRule
	for _, role := range roles {
		list = append(list, &entity.CasbinRule{
			Ptype: "p",
			V0:    role,
			V1:    req.Api,
			V2:    req.Method,
		})
	}
	return ctx.DB().Create(&list).Error
}

func (infra *Rbac) DeleteRbacApi(ctx kratosx.Context, api, method string) error {
	defer func() {
		err := ctx.Authentication().Enforce().LoadPolicy()
		if err != nil {
			ctx.Logger().Errorf("DeleteRbacApi LoadPolicy() failed: %s", err.Error())
		}
	}()

	return ctx.DB().Where("v1=? and v2=?", api, method).Delete(entity.CasbinRule{}).Error
}

func (infra *Rbac) UpdateRbacApi(ctx kratosx.Context, old types.MenuApi, now types.MenuApi) error {
	defer func() {
		err := ctx.Authentication().Enforce().LoadPolicy()
		if err != nil {
			ctx.Logger().Errorf("UpdateRbacApi LoadPolicy() failed: %s", err.Error())
		}
	}()

	return ctx.DB().
		Model(entity.CasbinRule{}).
		Where("v1=? and v2=?", old.Api, old.Method).
		UpdateColumn("v1", now.Api).
		UpdateColumn("v2", now.Method).
		Error
}

func (infra *Rbac) UpdateRbacRoleApis(ctx kratosx.Context, role string, apis []*types.MenuApi) error {
	var list []*entity.CasbinRule
	for _, item := range apis {
		list = append(list, &entity.CasbinRule{
			Ptype: "p",
			V0:    role,
			V1:    item.Api,
			V2:    item.Method,
		})
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		defer func() {
			err := ctx.Authentication().Enforce().LoadPolicy()
			if err != nil {
				ctx.Logger().Errorf("UpdateRbacRoleApis LoadPolicy() failed: %s", err.Error())
			}
		}()
		if err := ctx.DB().Where("v0=?", role).Delete(&entity.CasbinRule{}).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return nil
		}
		return ctx.DB().Create(&list).Error
	})
}

func (infra *Rbac) DeleteRbacRoles(ctx kratosx.Context, roles []string) error {
	defer func() {
		err := ctx.Authentication().Enforce().LoadPolicy()
		if err != nil {
			ctx.Logger().Errorf("DeleteRbacRoles LoadPolicy() failed: %s", err.Error())
		}
	}()
	return ctx.DB().Where("v0 in ?", roles).Delete(&entity.CasbinRule{}).Error
}
