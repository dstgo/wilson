package system

import (
	"github.com/dstgo/wilson/internal/pkg/resp"
	"github.com/dstgo/wilson/internal/pkg/valid"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/code"
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
//	@Summary		Ping
//	@Description	to test app api if is ok
//	@Tags			system
//	@Accept			json
//	@Produce		json
//	@Param			name	query	string	true	"ping name"
//	@Router			/ping [GET]
func (p PingHandler) Ping(ctx *gin.Context) {
	pingReq := new(PingRequest)
	err := valid.BindAndResp(ctx,
		valid.Query(pingReq),
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
