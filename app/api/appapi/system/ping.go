package system

import (
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/logic/systemLogic"
	"github.com/dstgo/wilson/app/types/request"
	"github.com/gin-gonic/gin"
)

func NewPingApi(appConf *conf.AppConf, logic systemLogic.PingLogic) PingApi {
	api := PingApi{
		pingLogic: logic,
	}
	return api
}

type PingApi struct {
	pingLogic systemLogic.PingLogic
}

// Ping
//
//	@Summary		app test ping api
//	@Description	to test app api if is ok
//	@Tags			system
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	true	"comment"
//	@Success		200		{object}	{"code":2000,"msg":"pong","data":"hello wolrd! wilson"}	"success"
//	@Failure		400		{object}	{"code":4000,"error":"名称: 是必需值，不能为空"}	"name param empty"
//	@Failure		400		{object}	{"code":4000,"error":"名称: 是必需值，不能为空"}	"name param empty"
//	@Router			/ping [get]
func (p PingApi) Ping(ctx *gin.Context) {
	pingReq := new(request.PingRequest)
	err := vax.Binds(ctx,
		vax.Query(pingReq),
	)
	if err != nil {
		resp.Fail(ctx, 4000, err)
		return
	}

	res := p.pingLogic.Ping(ctx, pingReq.Name)

	resp.Ok(ctx, 2000, "pong", res)
}
