package system

import (
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/types"
	"github.com/dstgo/wilson/app/types/code"
	"github.com/dstgo/wilson/app/types/request"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var SystemProviderSet = wire.NewSet(
	NewPingHandler,
	NewPingLogic,
)

func NewPingHandler(logic PingLogic) PingHandler {
	return PingHandler{
		PingLogic: logic,
	}
}

type PingHandler struct {
	PingLogic PingLogic
}

// Ping
//
//	@Summary		app test ping api
//	@Description	to test app api if is ok
//	@Tags			system
//	@Accept			json
//	@Produce		json
//	@Param			name	query	string	true	"ping name"
//	@Router			/ping [GET]
func (p PingHandler) Ping(ctx *gin.Context) {
	pingReq := new(request.PingRequest)
	err := vax.BindAndResp(ctx,
		vax.Query(pingReq),
	)
	if err != nil {
		return
	}

	res := p.PingLogic.Ping(ctx, pingReq.Name)

	resp.Ok(ctx).Code(code.RequestOk).Msg("pong").
		Data(types.H{
			"reply": res,
		}).Send()
}
