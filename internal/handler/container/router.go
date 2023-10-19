package container

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/route"
)

// HandlerRouter just for wire injection, no real influence
type HandlerRouter types.NopType

type Handler struct {
}

func SetupRouter(api *route.Router, handler Handler) Handler {
	return types.NopObj
}
