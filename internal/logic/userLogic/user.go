package userLogic

import "github.com/google/wire"

var UserLogicSet = wire.NewSet(NewUserLogic)
