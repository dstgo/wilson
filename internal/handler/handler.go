package handler

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/handler/auth"
	"github.com/dstgo/wilson/internal/handler/email"
	"github.com/dstgo/wilson/internal/handler/middleware"
	"github.com/dstgo/wilson/internal/handler/system"
	"github.com/dstgo/wilson/internal/handler/user"
	"github.com/dstgo/wilson/internal/pkg/log"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"path"
)

const (
	BasePath = "/api"
	DocPath  = "/api/doc"
)

var HandlerProviderSet = wire.NewSet(
	auth.AuthRouterSet,
	email.EmailRouterSet,
	system.SystemRouterSet,
	user.UserRouterSet,
	wire.Struct(new(Router), "*"),
)

// Router has no influence, just for wire injection
type Router struct {
	Auth   auth.HandlerRouter
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

	// static swagger documentation
	if swaggerEnabled {
		httpserver.GET(path.Join(DocPath, "*any"), ginSwagger.CustomWrapHandler(Config, swaggerFiles.NewHandler()))
		log.L().Infof("visit AppAPI Doc on http://%s%s", cfg.ServerConf.HttpConf.Address, path.Join(DocPath, "index.html"))
	}

	authenticator := auth.NewCacheAuthor(cfg.JwtConf, auth.NewTokenRedisCache(datasource))

	// wrap http router
	handlerRouter := route.NewRouter(httpserver.RouterGroup.Group(BasePath))

	// attach global router components
	handlerRouter.Attach(
		middleware.UseCors(cfg.ServerConf.HttpConf.CorsConf),
	)

	// add middleware chains
	handlerRouter.Use(
		middleware.UseAuthenticate(authenticator),
	)

	return setupHandlerRouter(cfg, handlerRouter, datasource)
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

//	@title			Wilson App Internal API Documentation
//	@version		v1.0.0
//	@description	Wilson api documentation
//	@BasePath		/api
//go:generate swag init --generatedTime --instanceName appapi -g handler.go -d ./,../types,../core/resp --output ./swagger && swag fmt -g handler.go -d ./
