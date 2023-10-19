package role

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/api/role"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewRoleData() RoleData {
	return RoleData{}
}

type RoleData struct{}

func (r RoleData) GetRoleById(db *gorm.DB, id uint) (entity.Role, error) {
	var roleEntity entity.Role
	err := db.Where("id = ?", id).Find(&roleEntity).Error
	return roleEntity, err
}

func (r RoleData) GetRoleByCode(db *gorm.DB, code string) (entity.Role, error) {
	var roleEntity entity.Role
	err := db.Where("code = ?", code).Find(&roleEntity).Error
	return roleEntity, err
}

func (r RoleData) GetAllRoleList(db *gorm.DB) ([]entity.Role, error) {
	var entities []entity.Role
	err := db.Find(&entities).Error
	return entities, err
}

func (r RoleData) GetPageRoleList(db *gorm.DB, pageOpt role.PageOption) ([]entity.Role, error) {
	db.Scopes(data.Pages(pageOpt.Page, pageOpt.Size))
	if len(pageOpt.Search) > 0 {
		like := data.Like(pageOpt.Search)
		db = db.Where("name LIKE ? OR code LIKE ?", like, like)
	}
	var roleList []entity.Role
	err := db.Find(&roleList).Error
	return roleList, err
}

func (r RoleData) CreateRole(db *gorm.DB, role entity.Role) error {
	return db.Create(&role).Error
}

func (r RoleData) RemoveRole(db *gorm.DB, roleId uint) error {
	db = db.Begin()
	// remove the permission record belonging to the role firstly
	err := db.Model(entity.RolePermission{}).Where("role_id = ?", roleId).Delete(nil).Error
	if err != nil {
		db.Rollback()
		return err
	}

	// then remove the role
	if err = db.Unscoped().Model(entity.Role{}).Delete("id = ?", roleId).Error; err != nil {
		db.Rollback()
		return err
	}

	return db.Commit().Error
}

func (r RoleData) UpdateRole(db *gorm.DB, role entity.Role) error {
	return db.Model(role).Where("id = ?", role.ID).Update("name = ?", role.Name).Error
}

// GetRolePermIds get all id of the permissions belonging to specified role
func (r RoleData) GetRolePermIds(db *gorm.DB, id uint) ([]uint, error) {
	var permIds []uint
	err := db.Model(entity.Permission{}).Where("role_id = ?", id).Find(&permIds).Error
	return permIds, err
}

// GetRolePerms get all permissions belonging to specified role
func (r RoleData) GetRolePerms(db *gorm.DB, id uint) ([]entity.Permission, error) {
	var (
		perms     []entity.Permission
		rolePerms []entity.RolePermission
	)

	if err := db.Model(entity.RolePermission{}).Find(&rolePerms, "role_id = ?", id).Error; err != nil {
		return []entity.Permission{}, err
	}

	if err := db.Model(rolePerms).Association("Permission").Find(&perms); err != nil {
		return []entity.Permission{}, err
	}

	return perms, nil
}

// InsertRolePerms insert new records if key has no conflicts,or do nothing
func (r RoleData) InsertRolePerms(db *gorm.DB, roleId uint, permIds []uint) error {

	rolePermList := makeRolePerms(roleId, permIds)

	// create the new relation, if existed, do nothing
	err := db.Clauses(clause.OnConflict{
		Columns:     []clause.Column{{Name: "role_id"}, {Name: "perm_id"}},
		Where:       clause.Where{},
		TargetWhere: clause.Where{},
		DoNothing:   true,
	}).Create(&rolePermList).Error

	return err
}

func (r RoleData) RemoveRolePerms(db *gorm.DB, roleId uint, permIds []uint) error {
	if permIds == nil || len(permIds) == 0 {
		return nil
	}
	rolePermList := makeRolePerms(roleId, permIds)
	return db.Unscoped().Delete(&rolePermList).Error
}

func makeRolePerms(roleId uint, permIds []uint) []entity.RolePermission {
	rolePermList := make([]entity.RolePermission, 0, len(permIds))
	for _, permId := range permIds {
		rolePermList = append(rolePermList, entity.RolePermission{RoleId: roleId, PermId: permId})
	}
	return rolePermList
}

func NewPermData() PermData {
	return PermData{}
}

type PermData struct{}

// GetIds gets the list of permissions matching the given tag
func (p PermData) GetIds(db *gorm.DB, tag string) ([]uint, error) {
	var ids []uint
	err := db.Model(entity.Permission{}).Select("id").Find(&ids, "tag = ?", tag).Error
	return ids, err
}

func (p PermData) GetById(db *gorm.DB, id uint) (entity.Permission, error) {
	var perm entity.Permission
	err := db.Find(&perm, "id = ?", id).Error
	return perm, err
}

func (p PermData) GetByName(db *gorm.DB, name string) (entity.Permission, error) {
	var perm entity.Permission
	err := db.Find(&perm, "name = ?", name).Error
	return perm, err
}

func (p PermData) GetAllPermList(db *gorm.DB) ([]entity.Permission, error) {
	var perms []entity.Permission
	err := db.Find(&perms).Error
	return perms,

		err
}

func (p PermData) GetPagePermListBy(db *gorm.DB, pageOpt role.PageOption) ([]entity.Permission, error) {
	pageDB := db
	pageDB.Scopes(data.Pages(pageOpt.Page, pageOpt.Size))
	if len(pageOpt.Search) > 0 {
		pageDB = pageDB.Where("name LIKE ?", data.Like(pageOpt.Search))
	}
	var perms []entity.Permission
	err := pageDB.Find(&perms).Error
	return perms, err
}

func (p PermData) FindPerm(db *gorm.DB, obj, act, g, t string) (entity.Permission, error) {
	var perm entity.Permission
	err := db.Model(entity.Permission{}).
		Where("object = ? AND action = ? AND `group` = ? AND tag = ?", obj, act, g, t).
		Find(&perm).Error
	return perm, err
}

func (p PermData) CreatePerm(db *gorm.DB, role entity.Permission) error {
	return db.Create(&role).Error
}

func (p PermData) RemoveRole(db *gorm.DB, permId uint) error {
	return db.Unscoped().Model(entity.Permission{}).Delete("id = ?", permId).Error
}

func (p PermData) UpdateRole(db *gorm.DB, perm entity.Permission) error {
	return db.Model(perm).Where("id = ?", perm.ID).UpdateColumn("name", perm.Name).Error
}
