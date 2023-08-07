package dao

import (
	"github.com/dstgo/wilson/internal/dao/userDao"
	"github.com/google/wire"
)

var DaoProviderSet = wire.NewSet(
	userDao.UserDaoSet,
)
