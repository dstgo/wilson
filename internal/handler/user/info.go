package user

import (
	"errors"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/api/user"
	"github.com/dstgo/wilson/internal/types/errs"
	"gorm.io/gorm"
)

func NewUserInfo(userData InfoData) UserInfo {
	return UserInfo{
		userData: userData,
	}
}

type UserInfo struct {
	userData InfoData
}

func (u UserInfo) GetUserInfo(userId uint) (user.Info, error) {
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

func (u UserInfo) UpdateUserInfo(info user.UpdateInfoOption) error {
	var userTable entity.User

	// try to find the user
	if _, err := u.GetUserInfo(info.Id); err != nil {
		return err
	}

	if err := cp.Copy(&info, &userTable); err != nil {
		return errs.ProgramErr(err)
	}

	if err := u.userData.UpdateUserInfo(userTable); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}

func (u UserInfo) RemoveUser(userId uint) error {

	// try to find the user
	if _, err := u.GetUserInfo(userId); err != nil {
		return err
	}

	if err := u.userData.RemoveUser(userId); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}
