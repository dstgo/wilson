package user

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/dstgo/wilson/internal/types/user"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	user, err := u.userInfo.GetUserInfoByName(createOpt.Username)
	if user.UUID != "" {
		return errs.NewI18nError("user.alreadyExist")
	} else if err != nil {
		return err
	}

	err = CreateUser(u.ds.ORM(), entity.User{
		UUID:     uuid.NewString(),
		Username: createOpt.Username,
		Password: cryptor.Sha512WithBase64(createOpt.Password),
		Email:    createOpt.Email,
	})

	if err != nil {
		return errs.DataBaseErr(err)
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
		return errs.ProgramErr(err)
	}

	userTable.Password = cryptor.Sha512WithBase64(userTable.Password)

	if err := UpdateUserInfo(u.ds.ORM(), userTable); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}

func (u UserModify) Remove(uuid string) error {
	// try to find the user
	if _, err := u.userInfo.GetUserInfoByUUID(uuid); err != nil {
		return err
	}

	if err := RemoveByUUID(u.ds.ORM(), uuid); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}

func CreateUser(db *gorm.DB, user entity.User) error {
	return db.Create(&user).Error
}

func UpdateUserInfo(db *gorm.DB, user entity.User) error {
	return db.Where("uuid = ?", user.UUID).Updates(&user).Error
}

func DisableUser(db *gorm.DB, id uint) error {
	return db.Model(entity.User{}).Delete(entity.User{
		Model: gorm.Model{ID: id},
	}).Error
}

func RemoveUser(db *gorm.DB, id uint) error {
	return db.Unscoped().Model(entity.User{}).Delete(entity.User{
		Model: gorm.Model{ID: id},
	}).Error
}

func RemoveByUUID(db *gorm.DB, uuid string) error {
	return db.Unscoped().Model(entity.User{}).Delete("uuid = ?", uuid).Error
}
