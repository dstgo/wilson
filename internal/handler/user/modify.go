package user

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/dstgo/wilson/internal/types/user"
	"github.com/duke-git/lancet/v2/cryptor"
)

func NewUserModify(ds *data.DataSource, userdata UserData, userInfo UserInfo) UserModify {
	return UserModify{
		userData: userdata,
		userInfo: userInfo,
		ds:       ds,
	}
}

type UserModify struct {
	userData UserData
	userInfo UserInfo
	ds       *data.DataSource
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

	userTable.Password = cryptor.Sha512WithBase64(userTable.Password)

	if err := u.userData.UpdateUserInfo(u.ds.ORM(), userTable); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}

func (u UserModify) Remove(uuid string) error {
	// try to find the user
	if _, err := u.userInfo.GetUserInfoByUUID(uuid); err != nil {
		return err
	}

	if err := u.userData.RemoveByUUID(u.ds.ORM(), uuid); err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}
