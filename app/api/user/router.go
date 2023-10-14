package user

import (
	"github.com/dstgo/wilson/app/types"
	"github.com/dstgo/wilson/pkg/route"
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

func SetupRouter(open *route.Router, api API) APIRouter {
	// register open user info api
	return types.NopObj
}
