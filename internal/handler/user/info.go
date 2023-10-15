package user

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
)

func NewInfoLogic(userDao InfoData) InfoLogic {
	return InfoLogic{
		userDao: userDao,
	}
}

type InfoLogic struct {
	userDao InfoData
}

func NewInfoData(source *data.DataSource) InfoData {
	return InfoData{DataSource: source}
}

type InfoData struct {
	*data.DataSource
}

func (u InfoData) GetUserByName(username string) (entity.User, error) {
	user := entity.User{}
	err := u.ORM().Model(user).Where("username = ?", username).First(&user).Error
	return user, err
}

func (u InfoData) GetUserByUUID(uuid string) (entity.User, error) {
	user := entity.User{}
	err := u.ORM().Model(user).Where("uuid =?", uuid).First(&user).Error
	return user, err
}

func (u InfoData) GetUserByEmail(email string) (entity.User, error) {
	user := entity.User{}
	err := u.ORM().Model(user).Where("email =?", email).First(&user).Error
	return user, err
}

func (u InfoData) DeleteByUUID(uuid string) error {
	return u.ORM().Delete(&entity.User{}, "uuid =?", uuid).Error
}

func (u InfoData) CreateUser(user entity.User) error {
	return u.ORM().Model(entity.User{}).Create(&user).Error
}

func (u InfoData) ListAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := u.ORM().Model(entity.User{}).Find(&users).Error
	return users, err
}

func (u InfoData) Count() (int64, error) {
	var count int64
	err := u.ORM().Model(entity.User{}).Count(&count).Error
	return count, err
}

func (u InfoData) UpdateUserInfo(user entity.User) error {
	return u.ORM().Model(user).Save(&user).Error
}
