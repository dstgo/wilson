package systemLogic

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/pkg/coco"
)

func NewPingLogic(coco *coco.Core, conf *conf.AppConf) *PingLogic {
	return &PingLogic{conf: conf}
}

type PingLogic struct {
	conf *conf.AppConf
}

func (Ping *PingLogic) Ping() {

}
