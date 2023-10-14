package wilson

import (
	"fmt"
	"github.com/dstgo/wilson/app/api"
	_ "github.com/dstgo/wilson/app/api/swagger"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/core/log"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/handler"
	_ "github.com/dstgo/wilson/app/handler/swagger"
	"github.com/dstgo/wilson/app/middleware"
	"github.com/dstgo/wilson/app/pkg/vax"
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
		middleware.UseLogger(logger, handler.DocPath, api.DocPath),
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

// SetupHandler initializes app internal api router configuration
func SetupHandler(cfg *conf.AppConf, engine *gin.Engine, datasource *data.DataSource, pool *email.Pool) handler.Router {
	if cfg.ServerConf.Swagger {
		engine.GET(path.Join(handler.DocPath, "*any"), ginSwagger.CustomWrapHandler(handler.Config, swaggerFiles.NewHandler()))
		log.L().Infof("visit AppAPI Doc on http://%s%s", cfg.ServerConf.HttpConf.Address, path.Join(handler.DocPath, "index.html"))
	}

	// jwt authenticator
	jwtAuthenticator := auth.NewJwtAuthenticator(cfg.JwtConf, datasource.Redis())

	root := route.NewRouter(engine.RouterGroup.Group(handler.BasePath))

	// attach middleware to gin router
	root.Attach(
		middleware.UseCors(cfg.ServerConf.HttpConf.CorsConf),
	)

	root.Use(
		middleware.UseJwtAuthenticate(jwtAuthenticator),
	)

	return handler.SetupHandler(cfg, root, datasource, jwtAuthenticator, pool)
}

// SetupOpenAPI initializes app open api router configuration
func SetupOpenAPI(cfg *conf.AppConf, engine *gin.Engine, datasource *data.DataSource) api.Router {
	if !cfg.ServerConf.OpenAPI {
		return api.Router{}
	}
	if cfg.ServerConf.Swagger {
		engine.GET(path.Join(api.DocPath, "*any"), ginSwagger.CustomWrapHandler(api.Config, swaggerFiles.NewHandler()))
		log.L().Infof("visit OpenAPI Doc on http://%s%s", cfg.ServerConf.HttpConf.Address, path.Join(api.DocPath, "index.html"))
	}
	root := route.NewRouter(engine.RouterGroup.Group(api.BasePath))

	return api.SetupAPI(root, datasource)
}

func NewLocale(cfg *locale.Conf) (*locale.Locale, error) {
	l, err := locale.NewLocaleWithConf(cfg)
	if err != nil {
		return nil, fmt.Errorf("load language directory failed: %s", err.Error())
	}
	locale.Setup(l)
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
