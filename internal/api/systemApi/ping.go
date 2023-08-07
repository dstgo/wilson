package systemApi

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/logic/systemLogic"
	"github.com/gin-gonic/gin"
)

func NewPingApi(appcfg *conf.ServerConf, logic *systemLogic.PingLogic) *PingApi {
	api := &PingApi{
		p: logic,
	}
	return api
}

type PingApi struct {
	p *systemLogic.PingLogic
}

func (p *PingApi) Ping(ctx *gin.Context) {

}
