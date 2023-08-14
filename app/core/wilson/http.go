package wilson

import (
	"fmt"
	_ "github.com/dstgo/wilson/app/api/swagger"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/core/log"
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/middleware"
	"github.com/dstgo/wilson/app/types"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewHttpServer(cfg *conf.AppConf, lang *locale.Locale, logger *logrus.Logger, datasource *data.DataSource) (*gin.Engine, *http.Server) {

	serverConf := cfg.ServerConf

	engine := gin.New()
	gin.DisableConsoleColor()
	gin.DisableBindValidation()

	engine.MaxMultipartMemory = serverConf.Http.MultipartMax

	engine.Use(
		middleware.UseLogger(logger),
		middleware.UseAcceptLanguage(lang.Default()),
	)

	if cfg.ServerConf.Swagger {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	engine.NoMethod(middleware.NoMethodHandler(lang))
	engine.NoRoute(middleware.NotFoundHandler(lang))

	server := &http.Server{
		Addr:              serverConf.Http.Address,
		ReadTimeout:       serverConf.Http.ReadTimeout,
		ReadHeaderTimeout: serverConf.Http.ReadHeadTimeout,
		WriteTimeout:      serverConf.Http.WriteTimeout,
		IdleTimeout:       serverConf.Http.IdleTimeout,
		MaxHeaderBytes:    serverConf.Http.MaxHeader,
	}

	// http request validate pkg
	vax.SetTranslator(locale.L())

	server.Handler = engine

	return engine, server
}

func NewRouter(cfg *conf.AppConf, lang *locale.Locale, engine *gin.Engine, datasource *data.DataSource) *route.Router {
	router := route.NewRouter(engine.RouterGroup.Group("/api/v1/"))
	// jwt authenticator
	jwtAuthenticator := auth.NewJwtAuthenticator(cfg.JwtConf, lang, datasource.Redis)

	router.Use(
		middleware.UseJwtAuthenticate(jwtAuthenticator, lang),
	)

	return router
}

func NewLocale(cfg *locale.Conf) (*locale.Locale, error) {
	l, err := locale.NewLocaleWithConf(cfg)
	if err != nil {
		return nil, fmt.Errorf("load language directory failed: %s", err.Error())
	}
	locale.Set(l)
	return l, nil
}

// NewLogger config logrus middleware
// param logConf *conf.LogConf
// return *log.LoggerW
// return error
func NewLogger(logConf *conf.LogConf) (*log.LoggerW, error) {
	logConf.TimeFormat = conf.DateTimeFormat
	logConf.Order = []string{types.LogIpKey, types.LogHttpMethodKey, types.LogRequestPathKey, types.LogRequestUrlKey,
		types.LogHttpStatusKey, types.LogRequestCostKey, types.LogHttpContentLength, types.LogHttpResponseLength, types.LogRequestIdKey}
	logger, err := log.NewLogger(logConf)
	if err != nil {
		return nil, errors.Wrap(err, "load logger failed")
	}
	log.Set(logger.L())
	return logger, nil
}
