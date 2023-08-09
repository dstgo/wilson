package logic

import (
	"github.com/dstgo/wilson/app/logic/systemLogic"
	"github.com/dstgo/wilson/app/logic/userLogic"
	"github.com/google/wire"
)

var LogicSet = wire.NewSet(
	systemLogic.NewPingLogic,
	userLogic.NewUserLogic,
)
