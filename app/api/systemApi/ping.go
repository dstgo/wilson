package systemApi

import (
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/logic/systemLogic"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/dstgo/wilson/app/types/request"
	"github.com/gin-gonic/gin"
)

func NewPingApi(appConf *conf.AppConf, logic *systemLogic.PingLogic) *PingApi {
	api := &PingApi{
		pingLogic: logic,
	}
	return api
}

type PingApi struct {
	pingLogic *systemLogic.PingLogic
}

// Ping
// @tags system
// @param name query string true "name"
// @router /ping [GET]
func (p *PingApi) Ping(ctx *gin.Context) {
	pingReq := new(request.PingRequest)
	err := vax.Binds(ctx,
		vax.Query(pingReq),
	)
	if err != nil {
		httpx.Failed(ctx, 4000, err)
		return
	}

	res := p.pingLogic.Ping(ctx, pingReq.Name)

	httpx.Ok(ctx, 2000, "pong", res)
}
