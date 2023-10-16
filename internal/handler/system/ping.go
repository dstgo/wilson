package system

import (
	"fmt"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/types"
	"github.com/gin-gonic/gin"
	"time"
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
	return fmt.Sprintf("hello %s! Now is %s.", name, time.Now().Format(types.DateTimeFormat))
}
