package dao

import (
	"github.com/dstgo/wilson/app/dao/userDao"
	"github.com/google/wire"
)

var DaoProviderSet = wire.NewSet(
	userDao.UserDaoSet,
)
