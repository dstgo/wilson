package user

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/api/user"
	"gorm.io/gorm"
)

func NewUserData() UserData {
	return UserData{}
}

type UserData struct{}

func (u UserData) GetUserById(db *gorm.DB, id uint) (entity.User, error) {
	findUser := entity.User{}
	err := db.Where("id = ?", id).Find(&findUser).Error
	return findUser, err
}

func (u UserData) GetUserByName(db *gorm.DB, username string) (entity.User, error) {
	findUser := entity.User{}
	err := db.Where("username = ?", username).First(&findUser).Error
	return findUser, err
}

func (u UserData) GetUserByUUID(db *gorm.DB, uuid string) (entity.User, error) {
	findUser := entity.User{}
	err := db.Where("uuid =?", uuid).First(&findUser).Error
	return findUser, err
}

func (u UserData) GetUserByEmail(db *gorm.DB, email string) (entity.User, error) {
	findUser := entity.User{}
	err := db.Model(findUser).Where("email =?", email).First(&findUser).Error
	return findUser, err
}

func (u UserData) DeleteByUUID(db *gorm.DB, uuid string) error {
	return db.Delete(&entity.User{}, "uuid =?", uuid).Error
}

func (u UserData) CreateUser(db *gorm.DB, user entity.User) error {
	return db.Create(&user).Error
}

func (u UserData) ListByPage(db *gorm.DB, pageOpt user.PageOption) ([]entity.User, error) {
	pageDB := db
	pageDB.Scopes(data.Pages(pageOpt.Page, pageOpt.Size))
	if len(pageOpt.Order) > 0 {
		pageDB.Scopes(data.Order(pageOpt.Order, pageOpt.Desc))
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
func (u UserData) ListAllUsers(db *gorm.DB) ([]entity.User, error) {
	var users []entity.User
	err := db.Find(&users).Error
	return users, err
}

func (u UserData) Count(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(entity.User{}).Count(&count).Error
	return count, err
}

func (u UserData) UpdateUserInfo(db *gorm.DB, user entity.User) error {
	return db.Where("uuid = ?", user.UUID).Updates(&user).Error
}

func (u UserData) DisableUser(db *gorm.DB, id uint) error {
	return db.Model(entity.User{}).Delete(entity.User{
		Model: gorm.Model{ID: id},
	}).Error
}

func (u UserData) RemoveUser(db *gorm.DB, id uint) error {
	return db.Unscoped().Model(entity.User{}).Delete(entity.User{
		Model: gorm.Model{ID: id},
	}).Error
}

func (u UserData) RemoveByUUID(db *gorm.DB, uuid string) error {
	return db.Unscoped().Model(entity.User{}).Delete("uuid = ?", uuid).Error
}
