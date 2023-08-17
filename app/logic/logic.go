package logic

import (
	"github.com/dstgo/wilson/app/logic/systemLogic"
	"github.com/dstgo/wilson/app/logic/userLogic"
	"github.com/google/wire"
)

// The role of different sets is only for convenience
// and to avoid unused problems during wire injection
// and they are logically consistent

var AppLogicSet = wire.NewSet(
	systemLogic.ProviderSet,
	userLogic.ProviderSet,
)

var OpenLogicSet = wire.NewSet(
	userLogic.NewUserLogic,
)
