package user

import (
	"fmt"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/api/user"
	"gorm.io/gorm"
)

func NewUserData(source *data.DataSource) UserData {
	return UserData{DataSource: source}
}

type UserData struct {
	*data.DataSource
}

func (u UserData) GetUserById(id uint) (entity.User, error) {
	findUser := entity.User{}
	err := u.ORM().Where("id = ?", id).Find(&u).Error
	return findUser, err
}

func (u UserData) GetUserByName(username string) (entity.User, error) {
	findUser := entity.User{}
	err := u.ORM().Where("username = ?", username).First(&findUser).Error
	return findUser, err
}

func (u UserData) GetUserByUUID(uuid string) (entity.User, error) {
	findUser := entity.User{}
	err := u.ORM().Where("uuid =?", uuid).First(&findUser).Error
	return findUser, err
}

func (u UserData) GetUserByEmail(email string) (entity.User, error) {
	findUser := entity.User{}
	err := u.ORM().Model(findUser).Where("email =?", email).First(&findUser).Error
	return findUser, err
}

func (u UserData) DeleteByUUID(uuid string) error {
	return u.ORM().Delete(&entity.User{}, "uuid =?", uuid).Error
}

func (u UserData) CreateUser(user entity.User) error {
	return u.ORM().Create(&user).Error
}

func (u UserData) ListByPage(pageOpt user.PageOption) ([]entity.User, error) {

	pageDB := data.Page(u.ORM(), pageOpt.Page, pageOpt.Size)
	if len(pageOpt.Order) > 0 {
		if pageOpt.Desc {
			pageDB = pageDB.Order(fmt.Sprintf("%s", pageOpt.Order))
		} else {
			pageDB = pageDB.Order(fmt.Sprintf("%s DESC", pageOpt.Order))
		}
	}

	var (
		users []entity.User
	)

	if len(pageOpt.Search) > 0 {
		query := "%" + pageOpt.Search + "%"
		pageDB = pageDB.Where("username LIKE ? OR email LIKE ?", query, query)
	}

	err := pageDB.Find(&users).Error
	return users, err
}

// ListAllUsers
// in most time, you should use ListByPage
func (u UserData) ListAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := u.ORM().Find(&users).Error
	return users, err
}

func (u UserData) Count() (int64, error) {
	var count int64
	err := u.ORM().Model(entity.User{}).Count(&count).Error
	return count, err
}

func (u UserData) UpdateUserInfo(user entity.User) error {
	return u.ORM().Where("uuid = ?", user.UUID).Updates(&user).Error
}

func (u UserData) DisableUser(id uint) error {
	return u.ORM().Model(entity.User{}).Delete(entity.User{
		Model: gorm.Model{ID: id},
	}).Error
}

func (u UserData) RemoveUser(id uint) error {
	return u.ORM().Unscoped().Model(entity.User{}).Delete(entity.User{
		Model: gorm.Model{ID: id},
	}).Error
}

func (u UserData) RemoveByUUID(uuid string) error {
	return u.ORM().Unscoped().Model(entity.User{}).Delete("uuid = ?", uuid).Error
}
