package system

import (
	"errors"
	"github.com/dstgo/wilson/internal/core/role"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/types/errs"
	roleType "github.com/dstgo/wilson/internal/types/role"
)

func NewRoleEnforcer(source *data.DataSource) RoleEnforcer {
	return RoleEnforcer{resolver: role.NewGormResolver(source.ORM())}
}

type RoleEnforcer struct {
	resolver role.Resolver
}

func (r RoleEnforcer) ListRole(page roleType.PageOption) ([]roleType.RoleInfo, error) {
	listRole, err := r.resolver.ListRole(page)
	if err != nil {
		return []roleType.RoleInfo{}, errs.DataBaseErr(err)
	}
	return listRole, nil
}

func (r RoleEnforcer) ListRolePerms(roleId uint) ([]roleType.PermGroup, error) {
	perms, err := r.resolver.ListRolePerms(roleId)
	if err != nil {
		return []roleType.PermGroup{}, errs.DataBaseErr(err)
	}
	return perms, nil
}

func (r RoleEnforcer) CreateRole(option roleType.CreateRoleOption) error {
	info := roleType.RoleInfo{
		Name: option.Name,
		Code: option.Code,
	}
	err := r.resolver.CreateRole(info)
	if err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (r RoleEnforcer) UpdateRole(option roleType.UpdateRoleOption) error {
	info := roleType.RoleInfo{
		ID:   option.Id,
		Name: option.Name,
	}
	err := r.resolver.UpdateRole(info)
	if err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (r RoleEnforcer) UpdateRolePerms(option roleType.GrantOption) error {
	err := r.resolver.UpdateRolePerms(option.RoleId, option.Tag, option.PermIds)
	if errors.Is(err, role.ErrPermissionNotFound) {
		return errs.NewI18nError("perm.notfound").Err(err)
	} else {
		return errs.DataBaseErr(err)
	}
}

func (r RoleEnforcer) RemoveRole(roleId uint) error {
	err := r.resolver.RemoveRole(roleId)
	if err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (r RoleEnforcer) ListPerms(option roleType.PageOption) ([]roleType.PermInfo, error) {
	perms, err := r.resolver.ListPerms(option)
	if err != nil {
		return []roleType.PermInfo{}, errs.DataBaseErr(err)
	}
	return perms, nil
}

func (r RoleEnforcer) CreatePerm(option roleType.CreatePermOption) error {
	err := r.resolver.CreatePerm(roleType.PermInfo{
		Name:   option.Name,
		Object: option.Object,
		Group:  option.Group,
		Action: option.Action,
		Tag:    option.Tag,
	})

	if err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (r RoleEnforcer) UpdatePerm(option roleType.UpdatePermOption) error {
	err := r.resolver.UpdatePerm(roleType.PermInfo{
		ID:   option.Id,
		Name: option.Name,
	})

	if err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (r RoleEnforcer) RemovePerm(permId uint) error {
	err := r.resolver.RemovePerm(permId)
	if err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}
