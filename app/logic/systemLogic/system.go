package systemLogic

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewPingLogic,
	NewAuthLogic,
	NewEmailLogic,
)
