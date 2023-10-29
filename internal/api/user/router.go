package user

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/ginx"
	"github.com/google/wire"
)

// APIRouter just for wire injection, no real influence
type APIRouter types.NopType

var UserRouterSet = wire.NewSet(
	UserApiProviderSet,
	wire.Struct(new(API), "*"),
	SetupRouter,
)

type API struct {
	Info InfoApi
}

func SetupRouter(open *ginx.RouterGroup, api API) APIRouter {
	open.GET("/info", ginx.M(), api.Info.KeyInfo)
	return types.NopObj
}
