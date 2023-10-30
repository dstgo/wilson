package user

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/dstgo/wilson/internal/types/user"
	"gorm.io/gorm"
)

func (u UserInfo) GetUserRoles(uuid string) ([]role.RoleInfo, error) {
	db := u.ds.ORM()
	roleInfos := make([]role.RoleInfo, 0)
	queryUser, found, err := GetUserByUUID(db, uuid)
	if err != nil {
		return roleInfos, err
	} else if !found {
		return roleInfos, user.ErrUserNotFound
	}

	roles, err := ListAllUserRoles(u.ds.ORM(), queryUser.Id)
	if err != nil {
		return roleInfos, err
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

	queryUser, found, err := GetUserByUUID(db, uuid)
	if err != nil {
		return err
	} else if !found {
		return user.ErrUserNotFound
	}

	// confirm roles had been exists in db
	var findRoles []entity.Role
	result := db.Model(entity.Role{}).Where("id IN ?", saveRoleIds).Find(&findRoles)
	recordFound, err := data.HasRecordFound(result)
	if err != nil {
		return err
	} else if !recordFound {
		return role.ErrInvalidRoles
	}

	// replace user-role association
	err = db.Model(&queryUser).Association("Roles").Replace(&findRoles)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func ListAllUserRoles(db *gorm.DB, userId uint) ([]entity.Role, error) {
	var userRoles []entity.UserRole
	var roles []entity.Role

	err := db.Where("user_id = ?", userId).Find(&userRoles).Error
	if err != nil {
		return roles, system.ErrDatabase.Wrap(err)
	}

	err = db.Model(userRoles).Association("Role").Find(&roles)
	if err != nil {
		return roles, system.ErrDatabase.Wrap(err)
	}

	return roles, nil
}
