package wilson

import (
	"fmt"
	"github.com/dstgo/wilson/internal/api"
	_ "github.com/dstgo/wilson/internal/api/swagger"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/handler"
	"github.com/dstgo/wilson/internal/handler/middleware"
	_ "github.com/dstgo/wilson/internal/handler/swagger"
	locale2 "github.com/dstgo/wilson/internal/sys/locale"
	"github.com/dstgo/wilson/internal/sys/log"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"path"
)

// NewHttpServer initializes http server configuration
func NewHttpServer(cfg *conf.AppConf, lang *locale2.Locale, logger *logrus.Logger) (*gin.Engine, *http.Server) {

	serverConf := cfg.ServerConf

	gin.DisableConsoleColor()
	gin.DisableBindValidation()
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard

	engine := gin.New()

	engine.MaxMultipartMemory = serverConf.HttpConf.MultipartMax

	engine.Use(
		middleware.UseLogger(logger, path.Join(handler.DocPath, "*any"), path.Join(api.DocPath, "*any")),
		middleware.UseRecovery(logger),
		middleware.UseAcceptLanguage(lang.Default()),
	)

	engine.NoMethod(middleware.NoMethodHandler())
	engine.NoRoute(middleware.NotFoundHandler())

	gin.ErrorLogger()

	server := &http.Server{
		Addr:              serverConf.HttpConf.Address,
		ReadTimeout:       serverConf.HttpConf.ReadTimeout,
		ReadHeaderTimeout: serverConf.HttpConf.ReadHeadTimeout,
		WriteTimeout:      serverConf.HttpConf.WriteTimeout,
		IdleTimeout:       serverConf.HttpConf.IdleTimeout,
		MaxHeaderBytes:    serverConf.HttpConf.MaxHeader,
	}

	// http request validate pkg
	vax.SetTranslator(locale2.L())

	server.Handler = engine

	return engine, server
}

func NewLocale(cfg *locale2.Conf) (*locale2.Locale, error) {
	l, err := locale2.NewLocaleWithConf(cfg)
	if err != nil {
		return nil, fmt.Errorf("load language directory failed: %s", err.Error())
	}
	locale2.Setup(l)
	return l, nil
}

// NewLogger config logrus middleware
func NewLogger(logConf *conf.LogConf) (*log.Logger, error) {
	logConf.TimeFormat = types.DateTimeFormat
	logConf.Order = []string{
		types.LogIpKey, types.LogHttpMethodKey,
		types.LogHttpStatusKey, types.LogRequestPathKey,
		types.LogRequestUrlKey, types.LogRequestCostKey,
		types.LogRequestContentType, types.LogHttpContentLength,
		types.LogResponseContentType, types.LogHttpResponseLength,
		types.LogRecoverRequestKey, types.LogRecoverErrorKey,
		types.LogRecoverStackKey, types.LogRequestIdKey}
	logger, err := log.NewLogger(logConf)
	if err != nil {
		return nil, errors.Wrap(err, "load logger failed")
	}
	return logger, nil
}
