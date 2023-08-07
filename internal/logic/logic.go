package logic

import (
	"github.com/dstgo/wilson/internal/logic/systemLogic"
	"github.com/dstgo/wilson/internal/logic/userLogic"
	"github.com/google/wire"
)

var LogicSet = wire.NewSet(
	systemLogic.NewPingLogic,
	userLogic.NewUserLogic,
)
