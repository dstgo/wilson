package system

import (
	"fmt"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/auth"
	"time"
)

type PingApp struct {
	conf *conf.AppConf
}

func NewPingLogic(conf *conf.AppConf) PingApp {
	return PingApp{
		conf: conf,
	}
}

func (p PingApp) Ping(name string) auth.PingReply {
	return auth.PingReply{Reply: fmt.Sprintf("hello %s! Now is %s.", name, time.Now().Format(types.DateTimeFormat))}
}

func (p PingApp) Pong(name string) auth.PingReply {
	return auth.PingReply{Reply: fmt.Sprintf("goodbye %s! Now is %s.", name, time.Now().Format(types.DateTimeFormat))}
}
