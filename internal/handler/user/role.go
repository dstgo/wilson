package user

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/alg/collection"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/dstgo/wilson/internal/types/role"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewUserRole(ds *data.DataSource) UserRole {
	return UserRole{ds: ds}
}

type UserRole struct {
	ds *data.DataSource
}

func (u UserRole) GetUserRoles(uuid string) ([]role.RoleInfo, error) {
	db := u.ds.ORM()
	roleInfos := make([]role.RoleInfo, 0)
	queryUser, err := GetUserByUUID(db, uuid)
	if err != nil {
		return roleInfos, errs.DataBaseErr(err)
	} else if queryUser.ID == 0 {
		return roleInfos, errs.NewI18nError("user.notfound")
	}

	roles, err := ListAllUserRoles(u.ds.ORM(), queryUser.ID)
	if err != nil {
		return roleInfos, errs.DataBaseErr(err)
	}

	roleInfos = role.MakeRoleInfoList(roles)

	return roleInfos, nil
}

func (u UserRole) GetUserRoleCodes(uuid string) ([]string, error) {
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

func (u UserRole) SaveRoles(uuid string, saveRoleIds []uint) error {
	db := u.ds.ORM()
	queryUser, err := GetUserByUUID(db, uuid)
	if err != nil {
		return errs.DataBaseErr(err)
	} else if queryUser.ID == 0 {
		return errs.NewI18nError("user.notfound")
	}

	queryRoles, err := ListAllUserRoles(db, queryUser.ID)
	if err != nil {
		return errs.DataBaseErr(err)
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
		return errs.DataBaseErr(err)
	}

	// remove obsolete roles
	if err := RemoveUserRoleInBatch(tx, queryUser.ID, obsoleteRoleIds); err != nil {
		tx.Rollback()
		return errs.DataBaseErr(err)
	}

	if tx.Commit().Error != nil {
		return errs.DataBaseErr(tx.Commit().Error)
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
