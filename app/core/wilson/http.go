package wilson

import (
	"fmt"
	"github.com/dstgo/wilson/app/api/appapi"
	_ "github.com/dstgo/wilson/app/api/appapi/swagger"
	"github.com/dstgo/wilson/app/api/openapi"
	_ "github.com/dstgo/wilson/app/api/openapi/swagger"
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
	"github.com/jordan-wright/email"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path"
)

// NewHttpServer initializes http server configuration
func NewHttpServer(cfg *conf.AppConf, lang *locale.Locale, logger *logrus.Logger) (*gin.Engine, *http.Server) {

	serverConf := cfg.ServerConf

	engine := gin.New()
	gin.DisableConsoleColor()
	gin.DisableBindValidation()

	engine.MaxMultipartMemory = serverConf.HttpConf.MultipartMax

	engine.Use(
		middleware.UseLogger(logger, appapi.ApiDoc, openapi.ApiDoc),
		middleware.UseRecovery(logger),
		middleware.UseAcceptLanguage(lang.Default()),
	)

	engine.NoMethod(middleware.NoMethodHandler())
	engine.NoRoute(middleware.NotFoundHandler())

	server := &http.Server{
		Addr:              serverConf.HttpConf.Address,
		ReadTimeout:       serverConf.HttpConf.ReadTimeout,
		ReadHeaderTimeout: serverConf.HttpConf.ReadHeadTimeout,
		WriteTimeout:      serverConf.HttpConf.WriteTimeout,
		IdleTimeout:       serverConf.HttpConf.IdleTimeout,
		MaxHeaderBytes:    serverConf.HttpConf.MaxHeader,
	}

	// http request validate pkg
	vax.SetTranslator(locale.L())

	server.Handler = engine

	return engine, server
}

// NewAppApiRouter initializes app internal api router configuration
func NewAppApiRouter(cfg *conf.AppConf, lang *locale.Locale, engine *gin.Engine, datasource *data.DataSource, pool *email.Pool) appapi.ApiRouter {
	if cfg.ServerConf.Swagger {
		engine.GET(appapi.ApiDoc, ginSwagger.CustomWrapHandler(appapi.Config, swaggerFiles.NewHandler()))
		log.L().Infof("visit AppAPI Doc on http://%s%s", cfg.ServerConf.HttpConf.Address, path.Join(path.Dir(appapi.ApiDoc), "index.html"))
	}

	// jwt authenticator
	jwtAuthenticator := auth.NewJwtAuthenticator(cfg.JwtConf, lang, datasource.Redis())

	root := route.NewRouter(engine.RouterGroup.Group(appapi.BasePath))

	// attach middleware to gin router
	root.Attach(
		middleware.UseCors(cfg.ServerConf.HttpConf.CorsConf),
	)

	root.Use(
		middleware.UseJwtAuthenticate(jwtAuthenticator),
	)

	return appapi.NewApiRouter(cfg, root, datasource, jwtAuthenticator, pool)
}

// NewOpenApiRouter initializes app open api router configuration
func NewOpenApiRouter(cfg *conf.AppConf, lang *locale.Locale, engine *gin.Engine, datasource *data.DataSource) openapi.ApiRouter {
	if !cfg.ServerConf.OpenAPI {
		return openapi.ApiRouter{}
	}
	if cfg.ServerConf.Swagger {
		engine.GET(openapi.ApiDoc, ginSwagger.CustomWrapHandler(openapi.Config, swaggerFiles.NewHandler()))
		log.L().Infof("visit OpenAPI Doc on http://%s%s", cfg.ServerConf.HttpConf.Address, path.Join(path.Dir(openapi.ApiDoc), "index.html"))
	}
	root := route.NewRouter(engine.RouterGroup.Group(openapi.BasePath))

	return openapi.NewApiRouter(cfg, root, datasource)
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
func NewLogger(logConf *conf.LogConf) (*log.LoggerW, error) {
	logConf.TimeFormat = conf.DateTimeFormat
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
	log.Set(logger.L())
	return logger, nil
}
