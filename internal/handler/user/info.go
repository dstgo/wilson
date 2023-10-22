package user

import (
	"errors"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/internal/types/user"
	"gorm.io/gorm"
)

func NewUserInfo(ds *data.DataSource) UserInfo {
	return UserInfo{
		ds: ds,
	}
}

type UserInfo struct {
	ds *data.DataSource
}

func (u UserInfo) GetUserInfoByEmail(email string) (user.Info, error) {
	userEntity, err := GetUserByEmail(u.ds.ORM(), email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user.Info{}, errs.ResourceNotFound(err).I18n("user.notfound")
	}
	userInfo := user.Info{
		UUID:      userEntity.UUID,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		CreatedAt: helper.CreatedAt{CreatedAt: userEntity.CreatedAt},
	}
	return userInfo, nil
}

func (u UserInfo) GetUserInfoByUUID(uuid string) (user.Info, error) {
	userEntity, err := GetUserByUUID(u.ds.ORM(), uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user.Info{}, errs.ResourceNotFound(err).I18n("user.notfound")
	}
	userInfo := user.Info{
		UUID:      userEntity.UUID,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		CreatedAt: helper.CreatedAt{CreatedAt: userEntity.CreatedAt},
	}
	return userInfo, nil
}

func (u UserInfo) GetUserInfoByName(name string) (user.Info, error) {
	userEntity, err := GetUserByName(u.ds.ORM(), name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user.Info{}, errs.ResourceNotFound(err).I18n("user.notfound")
	}
	userInfo := user.Info{
		UUID:      userEntity.UUID,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		CreatedAt: helper.CreatedAt{CreatedAt: userEntity.CreatedAt},
	}
	return userInfo, nil
}

func (u UserInfo) GetUserInfoById(userId uint) (user.Info, error) {
	userEntity, err := GetUserById(u.ds.ORM(), userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user.Info{}, errs.ResourceNotFound(err).I18n("user.notfound")
	}
	userInfo := user.Info{
		UUID:      userEntity.UUID,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		CreatedAt: helper.CreatedAt{CreatedAt: userEntity.CreatedAt},
	}
	return userInfo, nil
}

func (u UserInfo) GetUserInfoList(opt user.PageOption) ([]user.Info, error) {
	pageUser, err := ListByPage(u.ds.ORM(), opt)
	if err != nil {
		return []user.Info{}, errs.DataBaseErr(err)
	}

	userInfoList := make([]user.Info, 0, len(pageUser))

	if err := cp.Copy(&pageUser, &userInfoList); err != nil {
		return []user.Info{}, errs.ProgramErr(err)
	}

	return userInfoList, err
}

func GetUserById(db *gorm.DB, id uint) (entity.User, error) {
	findUser := entity.User{}
	err := db.Where("id = ?", id).Find(&findUser).Error
	return findUser, err
}

func GetUserByName(db *gorm.DB, username string) (entity.User, error) {
	findUser := entity.User{}
	err := db.Where("username = ?", username).First(&findUser).Error
	return findUser, err
}

func GetUserByUUID(db *gorm.DB, uuid string) (entity.User, error) {
	findUser := entity.User{}
	err := db.Where("uuid =?", uuid).First(&findUser).Error
	return findUser, err
}

func GetUserByEmail(db *gorm.DB, email string) (entity.User, error) {
	findUser := entity.User{}
	err := db.Model(findUser).Where("email =?", email).First(&findUser).Error
	return findUser, err
}

func ListByPage(db *gorm.DB, pageOpt user.PageOption) ([]entity.User, error) {
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
func ListAllUsers(db *gorm.DB) ([]entity.User, error) {
	var users []entity.User
	err := db.Find(&users).Error
	return users, err
}

func Count(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(entity.User{}).Count(&count).Error
	return count, err
}
