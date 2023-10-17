package user

import (
	"errors"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/api/user"
	"github.com/dstgo/wilson/internal/types/errs"
	"gorm.io/gorm"
)

func NewUserInfo(userData UserData) UserInfo {
	return UserInfo{
		userData: userData,
	}
}

type UserInfo struct {
	userData UserData
}

func (u UserInfo) GetUserInfoByEmail(email string) (user.Info, error) {
	userEntity, err := u.userData.GetUserByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user.Info{}, errs.ResourceNotFound(err).I18n("user.notfound")
	}
	var userInfo user.Info
	if err := cp.Copy(&userEntity, &userInfo); err != nil {
		return user.Info{}, errs.ProgramErr(err)
	}
	return userInfo, nil
}

func (u UserInfo) GetUserInfoByUUID(uuid string) (user.Info, error) {
	userEntity, err := u.userData.GetUserByUUID(uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user.Info{}, errs.ResourceNotFound(err).I18n("user.notfound")
	}
	var userInfo user.Info
	if err := cp.Copy(&userEntity, &userInfo); err != nil {
		return user.Info{}, errs.ProgramErr(err)
	}
	return userInfo, nil
}

func (u UserInfo) GetUserInfoByName(name string) (user.Info, error) {
	userEntity, err := u.userData.GetUserByName(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user.Info{}, errs.ResourceNotFound(err).I18n("user.notfound")
	}
	var userInfo user.Info
	if err := cp.Copy(&userEntity, &userInfo); err != nil {
		return user.Info{}, errs.ProgramErr(err)
	}
	return userInfo, nil
}

func (u UserInfo) GetUserInfoById(userId uint) (user.Info, error) {
	userEntity, err := u.userData.GetUserById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user.Info{}, errs.ResourceNotFound(err).I18n("user.notfound")
	}
	var userInfo user.Info
	if err := cp.Copy(&userEntity, &userInfo); err != nil {
		return user.Info{}, errs.ProgramErr(err)
	}
	return userInfo, nil
}

func (u UserInfo) GetUserInfoList(opt user.PageOption) ([]user.Info, error) {
	pageUser, err := u.userData.ListByPage(opt)
	if err != nil {
		return []user.Info{}, errs.DataBaseErr(err)
	}

	userInfoList := make([]user.Info, 0, len(pageUser))

	if err := cp.Copy(&pageUser, &userInfoList); err != nil {
		return []user.Info{}, errs.ProgramErr(err)
	}

	return userInfoList, err
}
