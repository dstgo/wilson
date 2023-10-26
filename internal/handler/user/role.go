package user

import (
	"errors"
	"github.com/dstgo/wilson/internal/data/entity"
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
	} else if queryUser.Id == 0 {
		return roleInfos, user.ErrUserNotFound
	}

	roles, err := ListAllUserRoles(u.ds.ORM(), queryUser.Id)
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
		roleIds = append(roleIds, e.Id)
	}

	return u.SaveRoles(uuid, roleIds)
}

func (u UserModify) SaveRoles(uuid string, saveRoleIds []uint) error {
	db := u.ds.ORM()

	queryUser, err := GetUserByUUID(db, uuid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return system.ErrDatabase.Wrap(err)
	} else if queryUser.Id == 0 {
		return user.ErrUserNotFound
	}

	// confirm roles had been exists in db
	var findRoles []entity.Role
	err = db.Model(entity.Role{}).Where("id IN ?", saveRoleIds).Find(&findRoles).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return system.ErrDatabase.Wrap(err)
	} else if len(findRoles) != len(saveRoleIds) {
		return role.ErrInvalidRoles
	}

	// replace user-role association
	err = db.Model(&queryUser).Association("Roles").Replace(&findRoles)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
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
