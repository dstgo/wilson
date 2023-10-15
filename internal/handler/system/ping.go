package system

import (
	"fmt"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/gin-gonic/gin"
)

type PingLogic struct {
	conf *conf.AppConf
}

func NewPingLogic(conf *conf.AppConf) PingLogic {
	return PingLogic{
		conf: conf,
	}
}

func (Ping PingLogic) Ping(ctx *gin.Context, name string) string {
	return fmt.Sprintf("hello wolrd! %s", name)
}
