package server

import (
	"fmt"
	"github.com/dstgo/wilson/internal/api"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/middleware"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/pkg/coco"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newEngine(conf *conf.ServerConf) *gin.Engine {

	engine := gin.New()
	gin.DisableConsoleColor()
	gin.DisableBindValidation()

	engine.MaxMultipartMemory = conf.Http.MultipartMax

	return engine
}

func newHttpServer(conf *conf.ServerConf) *http.Server {
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", conf.Http.Port),
		ReadTimeout:       conf.Http.ReadTimeout,
		ReadHeaderTimeout: conf.Http.ReadHeadTimeout,
		WriteTimeout:      conf.Http.WriteTimeout,
		IdleTimeout:       conf.Http.IdleTimeout,
		MaxHeaderBytes:    conf.Http.MaxHeader,
	}
	return server
}

func attachRouter(config *conf.AppConf, co *coco.Core, lang *locale.Locale, datasource *data.DataSource) (api.ApiRouter, error) {
	root := co.RootRouter()

	root.Use(
		middleware.UseLogger(),
		middleware.UseCors(),
		middleware.UseJwt(),
		middleware.UseCasbin(),
	)

	return api.NewApiRouter(config, co, lang, datasource)
}
