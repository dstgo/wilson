package user

import (
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/api/user"
	"github.com/dstgo/wilson/internal/types/errs"
)

func NewUserModify(userdata UserData, userInfo UserInfo) UserModify {
	return UserModify{
		userData: userdata,
		userInfo: userInfo,
	}
}

type UserModify struct {
	userData UserData
	userInfo UserInfo
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

	if err := u.userData.UpdateUserInfo(userTable); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}

func (u UserModify) Remove(uuid string) error {
	// try to find the user
	if _, err := u.userInfo.GetUserInfoByUUID(uuid); err != nil {
		return err
	}

	if err := u.userData.RemoveByUUID(uuid); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}
