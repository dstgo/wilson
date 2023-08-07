package userLogic

import (
	"github.com/dstgo/wilson/internal/dao/userDao"
	"github.com/dstgo/wilson/pkg/coco"
)

func NewUserLogic(core *coco.Core, userdao *userDao.UserInfoDao) *UserInfoLogic {
	return &UserInfoLogic{userDao: userdao}
}

type UserInfoLogic struct {
	*coco.Core
	userDao *userDao.UserInfoDao
}
