package dao

import (
	"github.com/dstgo/wilson/app/dao/userDao"
	"github.com/google/wire"
)

// The role of different sets is only for convenience
// and to avoid unused problems during wire injection
// and they are logically consistent

var AppDaoSet = wire.NewSet(
	userDao.DaoSet,
)

var OpenDaoSet = wire.NewSet(
	userDao.DaoSet,
)
