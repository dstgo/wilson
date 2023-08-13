package wilson

import (
	"fmt"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/core/logw"
	"github.com/dstgo/wilson/app/middleware"
	"github.com/dstgo/wilson/app/types"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func newHttpServer(conf *conf.ServerConf, lang *locale.Locale, logger *logrus.Logger) (*gin.Engine, *http.Server) {

	engine := gin.New()
	gin.DisableConsoleColor()
	gin.DisableBindValidation()

	engine.MaxMultipartMemory = conf.Http.MultipartMax

	engine.Use(
		middleware.UseLogger(logger),
		middleware.UseAcceptLanguage(lang.Default()),
	)

	engine.NoMethod(middleware.NoMethodHandler(lang))
	engine.NoRoute(middleware.NotFoundHandler(lang))

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", conf.Http.Port),
		ReadTimeout:       conf.Http.ReadTimeout,
		ReadHeaderTimeout: conf.Http.ReadHeadTimeout,
		WriteTimeout:      conf.Http.WriteTimeout,
		IdleTimeout:       conf.Http.IdleTimeout,
		MaxHeaderBytes:    conf.Http.MaxHeader,
	}

	server.Handler = engine

	return engine, server
}

func newLocale(cfg *locale.Conf) (*locale.Locale, error) {
	l, err := locale.NewLocaleWithConf(cfg)
	if err != nil {
		return nil, fmt.Errorf("load language directory failed: %s", err.Error())
	}
	return l, nil
}

// NewLogger config logrus middleware
// param logConf *conf.LogConf
// return *logw.LoggerW
// return error
func NewLogger(logConf *conf.LogConf) (*logw.LoggerW, error) {
	logConf.TimeFormat = conf.DateTimeFormat
	logConf.Order = []string{types.LogIpKey, types.LogHttpMethodKey, types.LogRequestPathKey,
		types.LogHttpStatusKey, types.LogRequestIdKey, types.LogRequestCostKey, types.LogHttpContentLength, types.LogHttpResponseLength}
	logger, err := logw.NewLogger(logConf)
	if err != nil {
		return nil, errors.Wrap(err, "load logger failed")
	}
	return logger, nil
}
