package userDao

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/google/wire"
)

var UserDaoSet = wire.NewSet(NewUserInfoDao)

func NewUserInfoDao(source *data.DataSource) *UserInfoDao {
	return &UserInfoDao{data: source}
}

type UserInfoDao struct {
	data *data.DataSource
}
