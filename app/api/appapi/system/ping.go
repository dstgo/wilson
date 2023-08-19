package system

import (
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/logic/systemLogic"
	"github.com/dstgo/wilson/app/types/code"
	"github.com/dstgo/wilson/app/types/request"
	"github.com/gin-gonic/gin"
)

func NewPingApi(logic systemLogic.PingLogic) PingApi {
	api := PingApi{
		PingLogic: logic,
	}
	return api
}

type PingApi struct {
	PingLogic systemLogic.PingLogic
}

// Ping
//
//	@Summary		app test ping api
//	@Description	to test app api if is ok
//	@Tags			system
//	@Accept			json
//	@Produce		json
//	@Param			name	query	string	true	"ping name"
//	@Router			/ping [get]
func (p PingApi) Ping(ctx *gin.Context) {
	pingReq := new(request.PingRequest)
	err := vax.BindAndResp(ctx,
		vax.Query(pingReq),
	)
	if err != nil {
		return
	}

	res := p.PingLogic.Ping(ctx, pingReq.Name)

	resp.Ok(ctx).Code(code.RequestOk).Msg("pong").
		Data(gin.H{
			"say": res,
		}).Send()
}
