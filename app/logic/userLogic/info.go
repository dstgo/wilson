package userLogic

import (
	"github.com/dstgo/wilson/app/repo/dao/userDao"
)

func NewUserLogic(userDao userDao.UserInfoDao) UserInfoLogic {
	return UserInfoLogic{
		userDao: userDao,
	}
}

type UserInfoLogic struct {
	userDao userDao.UserInfoDao
}
