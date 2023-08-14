package userLogic

import (
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/dao/userDao"
	"github.com/sirupsen/logrus"
)

func NewUserLogic(userdao *userDao.UserInfoDao) *UserInfoLogic {
	return &UserInfoLogic{
		userDao: userdao,
	}
}

type UserInfoLogic struct {
	logger  *logrus.Logger
	locale  *locale.Locale
	userDao *userDao.UserInfoDao
}
