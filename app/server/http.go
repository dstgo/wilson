package server

import (
	"fmt"
	"github.com/dstgo/wilson/app/api"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/middleware"
	"github.com/dstgo/wilson/app/pkg/locale"
	"github.com/dstgo/wilson/app/pkg/logw"
	"github.com/dstgo/wilson/app/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

// newEngine config gin engine and create new instance
// param conf *conf.ServerConf
// return *gin.Engine
func newEngine(conf *conf.ServerConf, locale *locale.Locale) *gin.Engine {

	engine := gin.New()
	gin.DisableConsoleColor()
	gin.DisableBindValidation()

	engine.MaxMultipartMemory = conf.Http.MultipartMax
	engine.NoMethod(middleware.NoMethodHandler(locale))
	engine.NoRoute(middleware.NotFoundHandler(locale))

	return engine
}

// newHttpServer config http server and create new instance
// param conf *conf.ServerConf
// return *http.Server
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

func newLocale(cfg *locale.Conf) (*locale.Locale, error) {
	l, err := locale.NewLocaleWithConf(cfg)
	if err != nil {
		return nil, fmt.Errorf("load language directory failed: %s", err.Error())
	}
	return l, nil
}

// newLogger config logrus middleware
// param logConf *conf.LogConf
// return *logw.LoggerW
// return error
func newLogger(logConf *conf.LogConf) (*logw.LoggerW, error) {
	logConf.TimeFormat = conf.DateTimeFormat
	logConf.Order = []string{types.LogIpKey, types.LogHttpMethodKey, types.LogRequestPathKey,
		types.LogHttpStatusKey, types.LogRequestIdKey, types.LogRequestCostKey, types.LogHttpContentLength, types.LogHttpResponseLength}
	logger, err := logw.NewLogger(logConf)
	if err != nil {
		return nil, errors.Wrap(err, "load logger failed")
	}
	return logger, nil
}

// attachRouter attach router to engine and install some middleware
// param comps *Components
// return api.ApiRouter
// return error
func attachRouter(comps *Components) (api.ApiRouter, error) {
	comps.Engine.Use(
		middleware.UseLogger(comps.Logger.L()),
	)
	root := comps.Router

	root.Attach()

	return api.NewApiRouter(comps.Conf, comps.Router, comps.Logger.L(), comps.Lang, comps.Datasource)
}
