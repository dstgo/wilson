package user

import (
	"errors"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/internal/types/user"
	"gorm.io/gorm"
)

func NewUserInfo(ds *data.DataSource, userData UserData) UserInfo {
	return UserInfo{
		ds:       ds,
		userData: userData,
	}
}

type UserInfo struct {
	userData UserData
	ds       *data.DataSource
}

func (u UserInfo) GetUserInfoByEmail(email string) (user.Info, error) {
	userEntity, err := u.userData.GetUserByEmail(u.ds.ORM(), email)
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
	userEntity, err := u.userData.GetUserByUUID(u.ds.ORM(), uuid)
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
	userEntity, err := u.userData.GetUserByName(u.ds.ORM(), name)
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
	userEntity, err := u.userData.GetUserById(u.ds.ORM(), userId)
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
	pageUser, err := u.userData.ListByPage(u.ds.ORM(), opt)
	if err != nil {
		return []user.Info{}, errs.DataBaseErr(err)
	}

	userInfoList := make([]user.Info, 0, len(pageUser))

	if err := cp.Copy(&pageUser, &userInfoList); err != nil {
		return []user.Info{}, errs.ProgramErr(err)
	}

	return userInfoList, err
}
