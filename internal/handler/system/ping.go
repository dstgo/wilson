package system

import (
	"fmt"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/api/auth"
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

func (p PingLogic) Ping(name string) auth.PingReply {
	return auth.PingReply{Reply: fmt.Sprintf("hello %s! Now is %s.", name, time.Now().Format(types.DateTimeFormat))}
}

func (p PingLogic) Pong(name string) auth.PingReply {
	return auth.PingReply{Reply: fmt.Sprintf("goodbye %s! Now is %s.", name, time.Now().Format(types.DateTimeFormat))}
}
