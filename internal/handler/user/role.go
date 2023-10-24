package user

import (
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/alg/collection"
	"github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/dstgo/wilson/internal/types/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (u UserInfo) GetUserRoles(uuid string) ([]role.RoleInfo, error) {
	db := u.ds.ORM()
	roleInfos := make([]role.RoleInfo, 0)
	queryUser, err := GetUserByUUID(db, uuid)
	if err != nil {
		return roleInfos, system.ErrDatabase.Wrap(err)
	} else if queryUser.ID == 0 {
		return roleInfos, user.ErrUserNotFound
	}

	roles, err := ListAllUserRoles(u.ds.ORM(), queryUser.ID)
	if err != nil {
		return roleInfos, system.ErrDatabase.Wrap(err)
	}

	roleInfos = role.MakeRoleInfoList(roles)

	return roleInfos, nil
}

func (u UserInfo) GetUserRoleCodes(uuid string) ([]string, error) {
	var codes []string
	roles, err := u.GetUserRoles(uuid)
	if err != nil {
		return codes, err
	}

	for _, info := range roles {
		codes = append(codes, info.Code)
	}

	return codes, err
}

func (u UserModify) SaveRolesByCode(uuid string, codes []string) error {
	var roles []entity.Role

	db := u.ds.ORM()
	err := db.Model(entity.Role{}).Where("code IN ?", codes).Find(&roles).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	var roleIds []uint
	for _, e := range roles {
		roleIds = append(roleIds, e.ID)
	}

	return u.SaveRoles(uuid, roleIds)
}

func (u UserModify) SaveRoles(uuid string, saveRoleIds []uint) error {
	if len(saveRoleIds) == 0 {
		return nil
	}

	db := u.ds.ORM()
	queryUser, err := GetUserByUUID(db, uuid)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	} else if queryUser.ID == 0 {
		return user.ErrUserNotFound
	}

	// confirm roles had been exists in db
	var findRoles []entity.Role
	err = db.Model(entity.Role{}).Where("id IN ?", saveRoleIds).Find(&findRoles).Error
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	} else if len(findRoles) != len(saveRoleIds) {
		return role.ErrInvalidRoles
	}

	queryRoles, err := ListAllUserRoles(db, queryUser.ID)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	// convert ids
	var queryRoleIds []uint
	for _, queryRole := range queryRoles {
		queryRoleIds = append(queryRoleIds, queryRole.ID)
	}

	extraRoleIds := collection.DifferenceSet(queryRoleIds, saveRoleIds)
	obsoleteRoleIds := collection.DifferenceSet(saveRoleIds, queryRoleIds)

	tx := db.Begin()

	createdRecordList := role.MakeUserRoleRecordList(queryUser.ID, extraRoleIds)
	// insert extra roles
	if err := CreateUserRoleInBatch(tx, createdRecordList); err != nil {
		tx.Rollback()
		return system.ErrDatabase.Wrap(err)
	}

	// remove obsolete roles
	if err := RemoveUserRoleInBatch(tx, queryUser.ID, obsoleteRoleIds); err != nil {
		tx.Rollback()
		return system.ErrDatabase.Wrap(err)
	}

	if tx.Commit().Error != nil {
		return system.ErrDatabase.Wrap(tx.Commit().Error)
	}

	return nil
}

func CreateUserRoleInBatch(db *gorm.DB, userRoles []entity.UserRole) error {
	if len(userRoles) == 0 {
		return nil
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&userRoles).Error
}

func RemoveUserRoleInBatch(db *gorm.DB, userId uint, roleIds []uint) error {
	if len(roleIds) == 0 {
		return nil
	}
	return db.Where("user_id = ? AND role_id IN ?", userId, roleIds).Error
}

func ListAllUserRoles(db *gorm.DB, userId uint) ([]entity.Role, error) {
	var userRoles []entity.UserRole
	var roles []entity.Role

	err := db.Where("user_id = ?", userId).Find(&userRoles).Error
	if err != nil {
		return roles, err
	}

	err = db.Model(userRoles).Association("Role").Find(&roles)
	if err != nil {
		return roles, err
	}

	return roles, nil
}
