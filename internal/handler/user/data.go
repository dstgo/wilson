package user

import (
	"fmt"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/api/user"
	"gorm.io/gorm"
)

func NewInfoData(source *data.DataSource) InfoData {
	return InfoData{DataSource: source}
}

type InfoData struct {
	*data.DataSource
}

func (u InfoData) GetUserById(id uint) (entity.User, error) {
	findUser := entity.User{}
	err := u.ORM().Where("id = ?", id).Find(&u).Error
	return findUser, err
}

func (u InfoData) GetUserByName(username string) (entity.User, error) {
	findUser := entity.User{}
	err := u.ORM().Where("username = ?", username).First(&findUser).Error
	return findUser, err
}

func (u InfoData) GetUserByUUID(uuid string) (entity.User, error) {
	findUser := entity.User{}
	err := u.ORM().Where("uuid =?", uuid).First(&findUser).Error
	return findUser, err
}

func (u InfoData) GetUserByEmail(email string) (entity.User, error) {
	findUser := entity.User{}
	err := u.ORM().Model(findUser).Where("email =?", email).First(&findUser).Error
	return findUser, err
}

func (u InfoData) DeleteByUUID(uuid string) error {
	return u.ORM().Delete(&entity.User{}, "uuid =?", uuid).Error
}

func (u InfoData) CreateUser(user entity.User) error {
	return u.ORM().Create(&user).Error
}

func (u InfoData) ListByPage(pageOpt user.PageOption) ([]entity.User, error) {

	pageDB := data.Page(u.ORM(), pageOpt.Page, pageOpt.Size)
	if len(pageOpt.Order) > 0 {
		if pageOpt.Desc {
			pageDB = pageDB.Order(fmt.Sprintf("%s", pageOpt.Order))
		} else {
			pageDB = pageDB.Order(fmt.Sprintf("%s DESC", pageOpt.Order))
		}
	}

	var users []entity.User
	err := pageDB.Find(&users).Error
	return users, err
}

// ListAllUsers
// in most time, you should use ListByPage
func (u InfoData) ListAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := u.ORM().Find(&users).Error
	return users, err
}

func (u InfoData) Count() (int64, error) {
	var count int64
	err := u.ORM().Model(entity.User{}).Count(&count).Error
	return count, err
}

func (u InfoData) UpdateUserInfo(user entity.User) error {
	return u.ORM().Save(&user).Error
}

func (u InfoData) DisableUser(id uint) error {
	return u.ORM().Model(entity.User{}).Delete(entity.User{
		Model: gorm.Model{ID: id},
	}).Error
}

func (u InfoData) RemoveUser(id uint) error {
	return u.ORM().Unscoped().Model(entity.User{}).Delete(entity.User{
		Model: gorm.Model{ID: id},
	}).Error
}
