package role

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
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
	result := db.Where("id = ?", id).Find(&roleEntity)
	_, err := data.HasRecordFound(result)
	if err != nil {
		return roleEntity, system.ErrDatabase.Wrap(err)
	}
	return roleEntity, nil
}

func getRoleByCode(db *gorm.DB, code string) (entity.Role, error) {
	var roleEntity entity.Role
	_, err := data.HasRecordFound(db.Where("code = ?", code).Find(&roleEntity))
	if err != nil {
		return roleEntity, system.ErrDatabase.Wrap(err)
	}
	return roleEntity, nil
}

func listAllRoles(db *gorm.DB) ([]entity.Role, error) {
	var entities []entity.Role
	_, err := data.HasRecordFound(db.Find(&entities))
	if err != nil {
		return entities, system.ErrDatabase.Wrap(err)
	}
	return entities, nil
}

func listRoles(db *gorm.DB, pageOpt role.PageOption) ([]entity.Role, error) {
	db.Scopes(data.Pages(pageOpt.Page, pageOpt.Size))
	if len(pageOpt.Search) > 0 {
		like := data.Like(pageOpt.Search)
		db = db.Where("name LIKE ? OR code LIKE ?", like, like)
	}
	var roleList []entity.Role
	err := db.Find(&roleList).Error
	if err != nil {
		return roleList, system.ErrDatabase.Wrap(err)
	}
	return roleList, err
}

func createRole(db *gorm.DB, role entity.Role) (entity.Role, error) {
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "code"}},
		DoNothing: true,
	}).Create(&role).Error
	if err != nil {
		return entity.Role{}, system.ErrDatabase.Wrap(err)
	}
	return role, nil
}

func createRoleInBatch(db *gorm.DB, roles []entity.Role) ([]entity.Role, error) {
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "code"}},
		DoNothing: true,
	}).Create(&roles).Error
	if err != nil {
		return nil, system.ErrDatabase.Wrap(err)
	}
	return roles, err
}

func removeRole(db *gorm.DB, roleId uint) error {
	db = db.Begin()
	// remove the permission record belonging to the role firstly
	err := db.Model(entity.RolePermission{}).Where("role_id = ?", roleId).Delete(nil).Error
	if err != nil {
		db.Rollback()
		return system.ErrDatabase.Wrap(err)
	}

	// then remove the role
	if err = db.Unscoped().Model(entity.Role{}).Delete("id = ?", roleId).Error; err != nil {
		db.Rollback()
		return system.ErrDatabase.Wrap(err)
	}

	err = db.Commit().Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	return nil
}

func updateRole(db *gorm.DB, role entity.Role) error {
	err := db.Model(role).Where("id = ?", role.Id).Update("name = ?", role.Name).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

// GetRolePermIds get all id of the permissions belonging to specified role
func getRolePermIds(db *gorm.DB, id uint) ([]uint, error) {
	var permIds []uint
	err := db.Model(entity.Permission{}).Where("role_id = ?", id).Find(&permIds).Error
	if err != nil {
		return permIds, system.ErrDatabase.Wrap(err)
	}
	return permIds, nil
}

// GetRolePerms get all permissions belonging to specified role
func getRolePerms(db *gorm.DB, id uint) ([]entity.Permission, error) {
	var (
		perms     []entity.Permission
		rolePerms []entity.RolePermission
	)

	if err := db.Model(entity.RolePermission{}).Find(&rolePerms, "role_id = ?", id).Error; err != nil {
		return []entity.Permission{}, system.ErrDatabase.Wrap(err)
	}

	if err := db.Model(rolePerms).Association("Permission").Find(&perms); err != nil {
		return []entity.Permission{}, system.ErrDatabase.Wrap(err)
	}

	return perms, nil
}

func insertRolePerm(db *gorm.DB, roleId uint, permId uint) error {
	// create the new relation, if existed, do nothing
	err := db.Clauses(clause.OnConflict{
		Columns:     []clause.Column{{Name: "role_id"}, {Name: "permission_id"}},
		Where:       clause.Where{},
		TargetWhere: clause.Where{},
		DoNothing:   true,
	}).Create(&entity.RolePermission{RoleId: roleId, PermissionId: permId}).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

// insertRolePermBatch insert new records if key has no conflicts,or do nothing
func insertRolePermBatch(db *gorm.DB, roleId uint, permIds []uint) error {

	rolePermList := role.MakeRolePerms(roleId, permIds)

	// create the new relation, if existed, do nothing
	err := db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&rolePermList).Error

	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	return nil
}

func removeRolePerm(db *gorm.DB, roleId uint, permId uint) error {
	err := db.Model(entity.RolePermission{}).Where("role_id = ? AND permission_id = ?", roleId, permId).Delete(nil).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func removeRolePermBatch(db *gorm.DB, roleId uint, permIds []uint) error {
	if permIds == nil || len(permIds) == 0 {
		return nil
	}
	rolePermList := role.MakeRolePerms(roleId, permIds)
	err := db.Unscoped().Model(entity.RolePermission{}).Delete(&rolePermList).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}
