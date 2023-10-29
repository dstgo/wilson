package handler

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/authen"
	"github.com/dstgo/wilson/internal/core/log"
	roleSo "github.com/dstgo/wilson/internal/core/role"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/cache"
	_ "github.com/dstgo/wilson/internal/handler/docs"
	"github.com/dstgo/wilson/internal/handler/email"
	"github.com/dstgo/wilson/internal/handler/middleware"
	"github.com/dstgo/wilson/internal/handler/system"
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/dstgo/wilson/internal/pkg/utils"
	"github.com/dstgo/wilson/pkg/ginx"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"path"
)

const (
	BasePath = "/api/v1/"
	DocPath  = "/api/v1/doc"
)

var HandlerProviderSet = wire.NewSet(
	email.EmailRouterSet,
	system.SystemRouterSet,
	user.UserRouterSet,
	wire.Struct(new(Router), "*"),
)

// Router has no influence, just for wire injection
type Router struct {
	Email  email.HandlerRouter
	System system.HandlerRouter
	User   user.HandlerRouter
}

// SetupHandler wilson http handlers
func SetupHandler(cfg *conf.AppConf, httpserver *gin.Engine, datasource *data.DataSource) (Router, func(), error) {
	var (
		serverConf     = cfg.ServerConf
		swaggerEnabled = serverConf.Swagger
	)

	var (
		accessTokenCache   = cache.NewAccessTokenCache(datasource)
		refreshTokenCache  = cache.NewRefreshTokenCache(datasource)
		tokenAuthenticator = authen.NewRefreshTokenAuthor(cfg.JwtConf, accessTokenCache, refreshTokenCache)
		roleResolver       = roleSo.NewGormResolver(datasource.ORM())
		userInfo           = user.NewUserInfo(datasource)
	)

	// wrap http router
	handlerRouter := ginx.NewRouterGroup(httpserver.RouterGroup.Group(BasePath))

	// attach global router components
	handlerRouter.Attach(
		middleware.UseCors(cfg.ServerConf.HttpConf.CorsConf),
	)

	// add middleware chains
	handlerRouter.Use(
		middleware.UseAuthenticate(tokenAuthenticator),
		middleware.UseRoleAuthorize(roleResolver, userInfo),
	)

	router, cleanup, err := setupHandlerRouter(cfg, handlerRouter, datasource)
	if err != nil {
		return Router{}, nil, err
	}

	// initialize router role access
	err = initHandlerData(handlerRouter, datasource, roleResolver)
	if err != nil {
		return Router{}, nil, err
	}

	log.L().Debugln("print api router tree...")
	utils.PrintRouters(handlerRouter, false)

	// static swagger documentation
	if swaggerEnabled {
		httpserver.GET(path.Join(DocPath, "*any"), ginSwagger.CustomWrapHandler(Config, swaggerFiles.NewHandler()))
		log.L().Infof("visit AppAPI Doc on http://%s%s", cfg.ServerConf.HttpConf.Address, path.Join(DocPath, "index.html"))
	}

	return router, cleanup, nil
}

var Config = &ginSwagger.Config{
	URL:                      "doc.json",
	DocExpansion:             "list",
	InstanceName:             "appapi",
	Title:                    "Wilson AppAPI",
	DefaultModelsExpandDepth: 0,
	DeepLinking:              true,
	PersistAuthorization:     false,
	Oauth2DefaultClientID:    "",
}

// swagger declarative api comment

// @title	                        Wilson App Internal API Documentation
// @version		                    v1.0.0
// @description                     Wilson app http api documentation, use the Bearer Token to authenticate
// @description                     It should be noted that when using swagger doc for API debugging, the Token needs to be manually prefixed with Bearer.
// @contact.name                    dstgo
// @contact.url                     https://github.com/dstgo
// @BasePath	                    /api/v1
// @license.name                    MIT LICENSE
// @license.url                     https://mit-license.org/
// @securityDefinitions.apikey      BearerAuth
// @in                              header
// @name                            Authorization
//go:generate swag init --generatedTime --instanceName appapi -g handler.go -d ./,../types,../core/resp --output ./docs && swag fmt -g handler.go -d ./
