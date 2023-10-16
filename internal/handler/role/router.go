package role

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/route"
)

// HandlerRouter just for wire injection, no real influence
type HandlerRouter types.NopType

type Handler struct {
	Role RoleHandler
}

func SetupRouter(root *route.Router, handler Handler) HandlerRouter {

	return types.NopObj
}
