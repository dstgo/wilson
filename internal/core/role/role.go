package role

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/role"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormResolver) GetRole(roleId uint) (role.RoleInfo, error) {
	rocord, err := getRoleById(g.db, roleId)
	if err != nil {
		return role.RoleInfo{}, err
	}
	return role.MakeRoleInfo(rocord), nil
}

func (g GormResolver) GetRoleInBatch(roleIds []uint) ([]role.RoleInfo, error) {
	var roles []role.RoleInfo
	err := g.db.Model(entity.Role{}).Find(&roles, "id IN ?", roleIds).Error
	return roles, err
}

func (g GormResolver) GetRoleByCode(code string) (role.RoleInfo, error) {
	byCode, err := getRoleByCode(g.db, code)
	return role.MakeRoleInfo(byCode), err
}

func (g GormResolver) ListRole(option role.PageOption) ([]role.RoleInfo, error) {
	records, err := listRoles(g.db, option)
	if err != nil {
		return []role.RoleInfo{}, err
	}
	return role.MakeRoleInfoList(records), nil
}

func (g GormResolver) ListAllRole() ([]role.RoleInfo, error) {
	records, err := listAllRoles(g.db)
	if err != nil {
		return []role.RoleInfo{}, err
	}
	return role.MakeRoleInfoList(records), nil
}

func (g GormResolver) CreateRole(roleInfo role.RoleInfo) error {
	_, err := createRole(g.db, role.MakeRoleRecord(roleInfo))
	return err
}

func (g GormResolver) CreateRoleInBatch(roles []role.RoleInfo) error {
	records := role.MakeRoleRecordList(roles)
	_, err := createRoleInBatch(g.db, records)
	return err
}

func (g GormResolver) UpdateRole(roleInfo role.RoleInfo) error {
	return updateRole(g.db, role.MakeRoleRecord(roleInfo))
}

func (g GormResolver) RemoveRole(roleId uint) error {
	return removeRole(g.db, roleId)
}

func getRoleById(db *gorm.DB, id uint) (entity.Role, error) {
	var roleEntity entity.Role
	err := db.Where("id = ?", id).Find(&roleEntity).Error
	return roleEntity, err
}

func getRoleByCode(db *gorm.DB, code string) (entity.Role, error) {
	var roleEntity entity.Role
	err := db.Where("code = ?", code).Find(&roleEntity).Error
	return roleEntity, err
}

func listAllRoles(db *gorm.DB) ([]entity.Role, error) {
	var entities []entity.Role
	err := db.Find(&entities).Error
	return entities, err
}

func listRoles(db *gorm.DB, pageOpt role.PageOption) ([]entity.Role, error) {
	db.Scopes(data.Pages(pageOpt.Page, pageOpt.Size))
	if len(pageOpt.Search) > 0 {
		like := data.Like(pageOpt.Search)
		db = db.Where("name LIKE ? OR code LIKE ?", like, like)
	}
	var roleList []entity.Role
	err := db.Find(&roleList).Error
	return roleList, err
}

func createRole(db *gorm.DB, role entity.Role) (entity.Role, error) {
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "code"}},
		DoNothing: true,
	}).Create(&role).Error
	return role, err
}

func createRoleInBatch(db *gorm.DB, roles []entity.Role) ([]entity.Role, error) {
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "code"}},
		DoNothing: true,
	}).Create(&roles).Error
	return roles, err
}

func removeRole(db *gorm.DB, roleId uint) error {
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

func updateRole(db *gorm.DB, role entity.Role) error {
	return db.Model(role).Where("id = ?", role.Id).Update("name = ?", role.Name).Error
}

// GetRolePermIds get all id of the permissions belonging to specified role
func getRolePermIds(db *gorm.DB, id uint) ([]uint, error) {
	var permIds []uint
	err := db.Model(entity.Permission{}).Where("role_id = ?", id).Find(&permIds).Error
	return permIds, err
}

// GetRolePerms get all permissions belonging to specified role
func getRolePerms(db *gorm.DB, id uint) ([]entity.Permission, error) {
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

func insertRolePerm(db *gorm.DB, roleId uint, permId uint) error {
	// create the new relation, if existed, do nothing
	return db.Clauses(clause.OnConflict{
		Columns:     []clause.Column{{Name: "role_id"}, {Name: "permission_id"}},
		Where:       clause.Where{},
		TargetWhere: clause.Where{},
		DoNothing:   true,
	}).Create(&entity.RolePermission{RoleId: roleId, PermissionId: permId}).Error
}

// insertRolePermBatch insert new records if key has no conflicts,or do nothing
func insertRolePermBatch(db *gorm.DB, roleId uint, permIds []uint) error {

	rolePermList := MakeRolePerms(roleId, permIds)

	// create the new relation, if existed, do nothing
	err := db.Debug().Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&rolePermList).Error

	return err
}

func removeRolePerm(db *gorm.DB, roleId uint, permId uint) error {
	return db.Model(entity.RolePermission{}).
		Where("role_id = ? AND permission_id = ?", roleId, permId).Delete(nil).Error
}

func removeRolePermBatch(db *gorm.DB, roleId uint, permIds []uint) error {
	if permIds == nil || len(permIds) == 0 {
		return nil
	}
	rolePermList := MakeRolePerms(roleId, permIds)
	return db.Unscoped().Model(entity.RolePermission{}).Delete(&rolePermList).Error
}
