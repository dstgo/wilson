package role

import (
	"errors"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/role"
	"gorm.io/gorm"
)

var (
	ErrPermissionNotFound = errors.New("permission not found")
	ErrRoleNotFound       = errors.New("role not found")
	ErrHasNoPermission    = errors.New("role has not specified permission")
)

type Resolver interface {
	GetPerm(permId uint) (role.PermInfo, error)
	GetPermInBatch(permIds []uint) ([]role.PermInfo, error)
	MatchPerm(name, obj, act, group, tag string) (role.PermInfo, error)
	CreatePerm(permInfo role.PermInfo) error
	CreatePermInBatch(permInfo []role.PermInfo) error
	ListPerms(option role.PageOption) ([]role.PermInfo, error)
	ListAllPerms(tag string) ([]role.PermInfo, error)
	UpdatePerm(permInfo role.PermInfo) error
	RemovePerm(permId uint) error

	GetRole(roleId uint) (role.RoleInfo, error)
	GetRoleInBatch(roleIds []uint) ([]role.RoleInfo, error)
	GetRoleByCode(code string) (role.RoleInfo, error)
	ListRole(option role.PageOption) ([]role.RoleInfo, error)
	ListAllRole() ([]role.RoleInfo, error)
	CreateRole(roleInfo role.RoleInfo) error
	CreateRoleInBatch(roleInfo []role.RoleInfo) error
	UpdateRole(roleInfo role.RoleInfo) error
	RemoveRole(roleId uint) error

	// ListRolePerms list all permissions belonging to specified role
	ListRolePerms(roleId uint) ([]role.PermGroup, error)
	// AddRolePerm add new permission to specified role
	AddRolePerm(roleId uint, permId uint) error
	// RemoveRolePerm remove specified permission from specified role
	RemoveRolePerm(roleId uint, permId uint) error
	// UpdateRolePermBatch update specified role permissions in batch
	// it will add and remove some permissions to make the records sync with the incoming permIds
	// if there are some permissions not in permission record which is result of ListAllPerms, it will return error
	UpdateRolePermBatch(roleId uint, tag string, permIds []uint) error
	CreateRolePermBatch(roles role.RoleInfo, perms []role.PermInfo) error

	// ResolveAny judge specified role if any role is able to access permObj by the way of permAct
	ResolveAny(permObj string, permAct string, roles ...string) error
}

var gormResolver Resolver = GormResolver{}

func NewGormResolver(db *gorm.DB) GormResolver {
	return GormResolver{db: db}
}

type GormResolver struct {
	db *gorm.DB
}

func (g GormResolver) ResolveAny(permObj string, permAct string, roles ...string) error {

	db := g.db
	var findRoles []entity.Role

	found, err := data.HasRecordFound(db.Find(&findRoles, "code IN ?", roles))
	if err != nil {
		return err
	} else if !found {
		return ErrRoleNotFound
	}

	var perms []entity.Permission

	err = db.Model(&findRoles).Association("Perms").Find(&perms, "object = ? AND action = ?", permObj, permAct)
	if err != nil {
		return err
	} else if len(perms) == 0 {
		return ErrHasNoPermission
	}

	return nil
}

func (g GormResolver) ListRolePerms(roleId uint) ([]role.PermGroup, error) {
	permRecords, err := getRolePerms(g.db, roleId)
	if err != nil {
		return []role.PermGroup{}, err
	}
	return role.MakePermGroup(permRecords), nil
}

func (g GormResolver) AddRolePerm(roleId uint, permId uint) error {
	return insertRolePerm(g.db, roleId, permId)
}

func (g GormResolver) RemoveRolePerm(roleId uint, permId uint) error {
	return removeRolePerm(g.db, roleId, permId)
}

func (g GormResolver) UpdateRolePermBatch(roleId uint, tag string, newPermIds []uint) error {

	var findRole entity.Role

	found, err := data.HasRecordFound(g.db.First(&findRole, "id = ?", roleId))
	if err != nil {
		return err
	} else if !found {
		return ErrRoleNotFound
	}

	var findPerms []entity.Permission

	hasFound, err := data.HasRecordFound(g.db.Find(&findPerms, "id IN ?", newPermIds))
	if err != nil {
		return err
	} else if !hasFound || len(findPerms) < len(newPermIds) {
		return ErrPermissionNotFound
	}

	err = g.db.Model(&findRole).Association("Perms").Replace(findPerms)
	if err != nil {
		return err
	}
	return nil
}

func (g GormResolver) CreateRolePermBatch(roleInfo role.RoleInfo, perms []role.PermInfo) error {

	var findRole entity.Role

	found, err := data.HasRecordFound(g.db.Find(&findRole, "code = ?", roleInfo.Code))
	if err != nil {
		return err
	} else if !found {
		return ErrRoleNotFound
	}

	permsList := role.MakePermRecordList(perms)

	_, err = createPermInBatch(g.db, permsList)
	if err != nil {
		return err
	}

	permRecords, err := listAllPermsByPerms(g.db, permsList)
	if err != nil {
		return err
	}

	err = g.db.Model(&findRole).Association("Perms").Append(&permRecords)
	if err != nil {
		return err
	}
	return nil
}
