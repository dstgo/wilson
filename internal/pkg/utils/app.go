package utils

import (
	"github.com/dstgo/wilson/internal/sys/log"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func IsDebugMode() bool {
	return log.L().Level >= logrus.DebugLevel || gin.Mode() == gin.DebugMode
}

// PrintRouters
// use for debugging, print all the route has benn register
func PrintRouters(root *route.Router, printGroup bool) error {
	return root.Walk(func(info route.RouterInfo) error {
		if !printGroup && info.IsGroup {
			return nil
		}
		log.L().Debugf("Method:%s\tPath:%-20s\tMeta:%s", info.Method, info.FullPath, info.Meta)
		return nil
	})
}
