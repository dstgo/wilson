package userDao

import (
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/data/entity"
	"github.com/google/wire"
)

var DaoSet = wire.NewSet(NewUserInfoDao)

func NewUserInfoDao(source *data.DataSource) UserInfoDao {
	return UserInfoDao{DataSource: source}
}

type UserInfoDao struct {
	*data.DataSource
}

func (u UserInfoDao) GetUserByName(username string) (entity.User, error) {
	user := entity.User{}
	err := u.ORM().Model(user).Where("username = ?", username).First(&user).Error
	return user, err
}

func (u UserInfoDao) GetUserByUUID(uuid string) (entity.User, error) {
	user := entity.User{}
	err := u.ORM().Model(user).Where("uuid =?", uuid).First(&user).Error
	return user, err
}

func (u UserInfoDao) GetUserByEmail(email string) (entity.User, error) {
	user := entity.User{}
	err := u.ORM().Model(user).Where("email =?", email).First(&user).Error
	return user, err
}

func (u UserInfoDao) DeleteByUUID(uuid string) error {
	return u.ORM().Delete(&entity.User{}, "uuid =?", uuid).Error
}

func (u UserInfoDao) CreateUser(user entity.User) error {
	return u.ORM().Model(entity.User{}).Create(&user).Error
}

func (u UserInfoDao) ListAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := u.ORM().Model(entity.User{}).Find(&users).Error
	return users, err
}

func (u UserInfoDao) Count() (int64, error) {
	var count int64
	err := u.ORM().Model(entity.User{}).Count(&count).Error
	return count, err
}

func (u UserInfoDao) UpdateUser(user entity.User) error {
	return u.ORM().Model(user).Save(&user).Error
}
