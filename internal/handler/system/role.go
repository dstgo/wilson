package system

import (
	"errors"
	"github.com/dstgo/wilson/internal/core/role"
	"github.com/dstgo/wilson/internal/data"
	roleType "github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
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
		return []roleType.RoleInfo{}, system.ErrDatabase.Wrap(err)
	}
	return listRole, nil
}

func (r RoleEnforcer) ListRolePerms(roleId uint) ([]roleType.PermGroup, error) {
	perms, err := r.resolver.ListRolePerms(roleId)
	if err != nil {
		return []roleType.PermGroup{}, system.ErrDatabase.Wrap(err)
	}
	return perms, nil
}

func (r RoleEnforcer) CreateRole(option roleType.CreateRoleOption) error {
	roleInfo, err := r.resolver.GetRoleByCode(option.Code)
	if err == nil && len(roleInfo.Code) > 0 {
		return roleType.ErrRoleConflict
	}

	info := roleType.RoleInfo{
		Name: option.Name,
		Code: option.Code,
	}

	err = r.resolver.CreateRole(info)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func (r RoleEnforcer) UpdateRole(option roleType.UpdateRoleOption) error {

	roleInfo, err := r.resolver.GetRole(option.Id)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	} else if roleInfo.Id == 0 {
		return roleType.ErrRoleNotFound
	}

	info := roleType.RoleInfo{
		Id:   option.Id,
		Name: option.Name,
	}

	err = r.resolver.UpdateRole(info)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func (r RoleEnforcer) UpdateRolePerms(option roleType.GrantOption) error {

	roleInfo, err := r.resolver.GetRole(option.RoleId)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	} else if roleInfo.Id == 0 {
		return roleType.ErrRoleNotFound
	}

	err = r.resolver.UpdateRolePermBatch(option.RoleId, option.Tag, option.PermIds)
	if errors.Is(err, role.ErrPermissionNotFound) {
		return role.ErrPermissionNotFound
	} else {
		return system.ErrDatabase.Wrap(err)
	}
}

func (r RoleEnforcer) RemoveRole(roleId uint) error {

	roleInfo, err := r.resolver.GetRole(roleId)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	} else if roleInfo.Id == 0 {
		return roleType.ErrRoleNotFound
	}

	err = r.resolver.RemoveRole(roleId)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func (r RoleEnforcer) ListPerms(option roleType.PageOption) ([]roleType.PermInfo, error) {
	perms, err := r.resolver.ListPerms(option)
	if err != nil {
		return []roleType.PermInfo{}, system.ErrDatabase.Wrap(err)
	}
	return perms, nil
}

func (r RoleEnforcer) CreatePerm(option roleType.CreatePermOption) error {

	perm, err := r.resolver.MatchPerm(option.Name, option.Object, option.Action, option.Group, option.Tag)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	} else if perm.Id > 0 {
		return roleType.ErrPermCojnflict
	}

	err = r.resolver.CreatePerm(roleType.PermInfo{
		Name:   option.Name,
		Object: option.Object,
		Group:  option.Group,
		Action: option.Action,
		Tag:    option.Tag,
	})

	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func (r RoleEnforcer) UpdatePerm(option roleType.UpdatePermOption) error {
	perm, err := r.resolver.GetPerm(option.Id)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	} else if perm.Id == 0 {
		return roleType.ErrPermNotFound
	}

	err = r.resolver.UpdatePerm(roleType.PermInfo{
		Id:   option.Id,
		Name: option.Name,
	})

	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func (r RoleEnforcer) RemovePerm(permId uint) error {
	perm, err := r.resolver.GetPerm(permId)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	} else if perm.Id == 0 {
		return roleType.ErrPermNotFound
	}

	err = r.resolver.RemovePerm(permId)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}
