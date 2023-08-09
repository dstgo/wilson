package userLogic

import (
	"github.com/dstgo/wilson/app/dao/userDao"
	"github.com/dstgo/wilson/app/pkg/locale"
	"github.com/sirupsen/logrus"
)

func NewUserLogic(logger *logrus.Logger, lang *locale.Locale, userdao *userDao.UserInfoDao) *UserInfoLogic {
	return &UserInfoLogic{
		locale:  lang,
		logger:  logger,
		userDao: userdao,
	}
}

type UserInfoLogic struct {
	logger  *logrus.Logger
	locale  *locale.Locale
	userDao *userDao.UserInfoDao
}
