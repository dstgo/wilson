package role

import (
	"errors"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/alg/collection"
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
	CreatePerm(permInfo role.PermInfo) error
	CreatePermInBatch(permInfo []role.PermInfo) error
	ListPerms(option role.PageOption) ([]role.PermInfo, error)
	ListAllPerms(tag string) ([]role.PermInfo, error)
	UpdatePerm(permInfo role.PermInfo) error
	RemovePerm(permId uint) error

	GetRole(roleId uint) (role.RoleInfo, error)
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
	var roleIds []uint
	// find all roles
	err := g.db.Model(entity.Role{}).Select("id").Where("code IN ?", roles).Find(&roleIds).Error
	if err != nil {
		return err
	}

	if len(roleIds) < len(roles) {
		return ErrRoleNotFound
	}

	var perm entity.Permission

	err = g.db.Model(entity.Permission{}).Where("object = ? AND action = ?", permObj, permAct).Find(&perm).Error
	if err != nil {
		return err
	}

	var rolePerms []entity.RolePermission
	err = g.db.Model(entity.RolePermission{}).Where("role_id IN ? AND perm_id = ?", roleIds, perm.ID).Find(&rolePerms).Error
	if err != nil {
		return err
	}

	if len(rolePerms) == 0 {
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

	// get all the tag permissions
	tagPerms, err := g.ListAllPerms(tag)
	if err != nil {
		return err
	}

	allPermIds := make([]uint, 0, len(tagPerms))
	for _, perm := range tagPerms {
		allPermIds = append(allPermIds, perm.ID)
	}

	// there has unexpected permId in params if is not subset
	if !collection.IsSubset(newPermIds, allPermIds) {
		return ErrPermissionNotFound
	}

	// find the complement set
	deletedSet := collection.ComplementSet(newPermIds, allPermIds)

	// start transaction
	tx := g.db.Begin()

	// insert new permIds
	if err = insertRolePermBatch(tx, roleId, newPermIds); err != nil {
		tx.Rollback()
		return err
	}

	// remove the complement set
	if err = removeRolePermBatch(tx, roleId, deletedSet); err != nil {
		tx.Rollback()
		return err
	}

	// commit
	return tx.Commit().Error
}

func (g GormResolver) CreateRolePermBatch(roleInfo role.RoleInfo, perms []role.PermInfo) error {

	tx := g.db.Begin()

	// try to create role if code conflict do nothing
	_, err := createRole(tx, role.MakeRoleRecord(roleInfo))
	if err != nil {
		tx.Rollback()
		return err
	}

	ens := role.MakePermRecordList(perms)
	// try to create perm if conflict do nothing
	_, err = createPermInBatch(tx, ens)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	// query roles
	queryRole, err := getRoleByCode(g.db, roleInfo.Code)
	if err != nil {
		return err
	}

	// query the permIds
	queryEns, err := listAllPermsByPerms(g.db, ens)
	if err != nil {
		return err
	}

	var permIds []uint
	for _, en := range queryEns {
		permIds = append(permIds, en.ID)
	}

	// create the relation
	err = insertRolePermBatch(g.db, queryRole.ID, permIds)
	if err != nil {
		return err
	}

	return nil
}
