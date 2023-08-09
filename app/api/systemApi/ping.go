package systemApi

import (
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/logic/systemLogic"
	"github.com/dstgo/wilson/app/pkg/httputil"
	"github.com/gin-gonic/gin"
)

func NewPingApi(appConf *conf.AppConf, logic *systemLogic.PingLogic) *PingApi {
	api := &PingApi{
		p: logic,
	}
	return api
}

type PingApi struct {
	p *systemLogic.PingLogic
}

func (p *PingApi) Ping(ctx *gin.Context) {
	httputil.Ok(ctx, 2000, "pong", p.p.Ping(ctx))
}
