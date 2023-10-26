package user

import (
	"errors"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/dstgo/wilson/internal/types/user"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewUserModify(ds *data.DataSource, userInfo UserInfo) UserModify {
	return UserModify{
		userInfo: userInfo,
		ds:       ds,
	}
}

type UserModify struct {
	userInfo UserInfo
	ds       *data.DataSource
}

func (u UserModify) Create(createOpt user.CreateUserOption) error {
	// try to find the user
	info, err := u.userInfo.GetUserInfoByName(createOpt.Username)
	if err == nil && info.UUID != "" {
		return user.ErrUserAlreadyExists
	} else if err != nil && !errors.Is(err, user.ErrUserNotFound) {
		return err
	}

	newUser := entity.User{
		UUID:     uuid.NewString(),
		Username: createOpt.Username,
		Password: cryptor.Sha512WithBase64(createOpt.Password),
		Email:    createOpt.Email,
	}

	// create new user
	if err := CreateUser(u.ds.ORM(), newUser); err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	// grant permission
	if err := u.SaveRolesByCode(newUser.UUID, createOpt.Roles); err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	return nil
}

func (u UserModify) Save(saveOpt user.SaveUserDetailOption) error {
	// try to find the user
	if _, err := u.userInfo.GetUserInfoByUUID(saveOpt.UUID); err != nil {
		return err
	}

	userEn := entity.User{
		UUID:     saveOpt.UUID,
		Username: saveOpt.Username,
		Password: cryptor.Sha512WithBase64(saveOpt.Password),
		Email:    saveOpt.Email,
	}

	// user user info
	err := UpdateUserInfo(u.ds.ORM(), userEn)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	// update roles
	err = u.SaveRolesByCode(userEn.UUID, saveOpt.Roles)
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	return nil
}

func (u UserModify) Update(updateOpt user.UpdateInfoOption) error {
	var userTable entity.User

	// try to find the user
	if _, err := u.userInfo.GetUserInfoByUUID(updateOpt.UUID); err != nil {
		return err
	}

	if err := cp.Copy(&updateOpt, &userTable); err != nil {
		return system.ErrProgram.Wrap(err)
	}

	if len(userTable.Password) > 0 {
		userTable.Password = cryptor.Sha512WithBase64(userTable.Password)
	}

	if err := UpdateUserInfo(u.ds.ORM(), userTable); err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	return nil
}

func (u UserModify) Remove(uuid string) error {
	db := u.ds.ORM()
	findUser, err := GetUserByUUID(db, uuid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return system.ErrDatabase.Wrap(err)
	} else if len(findUser.UUID) == 0 {
		return user.ErrUserNotFound
	}

	err = u.ds.ORM().Model(&findUser).Association("Roles").Clear()
	if err != nil {
		return system.ErrDatabase.Wrap(err)
	}

	if err := RemoveByUUID(u.ds.ORM(), uuid); err != nil {
		return system.ErrDatabase.Wrap(err)
	}
	return nil
}

func CreateUser(db *gorm.DB, user entity.User) error {
	return db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&user).Error
}

func UpdateUserInfo(db *gorm.DB, user entity.User) error {
	return db.Where("uuid = ?", user.UUID).Updates(&user).Error
}

func RemoveUser(db *gorm.DB, id uint) error {
	return db.Delete(entity.User{}, "id = ?", id).Error
}

func RemoveByUUID(db *gorm.DB, uuid string) error {
	return db.Delete(&entity.User{}, "uuid = ?", uuid).Error
}
