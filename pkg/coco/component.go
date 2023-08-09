package coco

import (
	"github.com/dstgo/wilson/pkg/coco/route"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type ComponentFn func(core *Core)

func WithConfig(cfg any) ComponentFn {
	return func(coco *Core) {
		coco.config = cfg
	}
}

func WithLogger(logger *logrus.Logger) ComponentFn {
	return func(coco *Core) {
		coco.logger = logger
	}
}

func WithRouter(router *route.Router) ComponentFn {
	return func(coco *Core) {
		coco.router = router
	}
}

func WithEngine(e *gin.Engine) ComponentFn {
	return func(coco *Core) {
		coco.engine = e
	}
}

func WithServer(server *http.Server) ComponentFn {
	return func(coco *Core) {
		coco.server = server
	}
}

type InterruptFn func(core *Core, signal os.Signal)
