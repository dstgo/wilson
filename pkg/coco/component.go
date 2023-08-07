package coco

import (
	"context"
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

func WithCtx(ctx context.Context) ComponentFn {
	return func(coco *Core) {
		coco.ctx = ctx
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

func ShutdownWithInfo() InterruptFn {
	return func(coco *Core, signal os.Signal) {
		coco.logger.Infof("received os signal: %s, ready to graceful shutdown\n", signal.String())
	}
}

func ShutdownWithCloseHttp() InterruptFn {
	return func(coco *Core, signal os.Signal) {
		err := coco.server.Shutdown(coco.ctx)
		coco.logger.Infof("ready to close http server: %s", err)
	}
}
