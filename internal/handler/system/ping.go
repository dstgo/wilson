package system

import (
	"fmt"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/system"
	"time"
)

type PingApp struct {
	conf *conf.WilsonConf
}

func NewPingLogic(conf *conf.WilsonConf) PingApp {
	return PingApp{
		conf: conf,
	}
}

func (p PingApp) Ping(name string) system.PingReply {
	return system.PingReply{Reply: fmt.Sprintf("hello %s! Now is %s.", name, time.Now().Format(types.DateTimeFormat))}
}

func (p PingApp) Pong(name string) system.PingReply {
	return system.PingReply{Reply: fmt.Sprintf("goodbye %s! Now is %s.", name, time.Now().Format(types.DateTimeFormat))}
}
