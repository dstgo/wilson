package node

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/ginx"
)

// HandlerRouter just for wire injection, no real influence
type HandlerRouter types.NopType

type Handler struct {
}

func SetupRouter(api *ginx.RouterGroup, handler Handler) Handler {
	return types.NopObj
}
