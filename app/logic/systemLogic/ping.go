package systemLogic

import (
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/pkg/locale"
	"github.com/dstgo/wilson/app/pkg/sysinfo"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewPingLogic(conf *conf.AppConf, logger *logrus.Logger, lang *locale.Locale) *PingLogic {
	return &PingLogic{
		conf:   conf,
		logger: logger,
		locale: lang,
	}
}

type PingLogic struct {
	logger *logrus.Logger
	locale *locale.Locale
	conf   *conf.AppConf
}

func (Ping *PingLogic) Ping(ctx *gin.Context) sysinfo.HostInfo {
	return sysinfo.GetHostInfo()
}
