package api

import (
	_ "github.com/dstgo/wilson/internal/api/docs"
	"github.com/dstgo/wilson/internal/api/user"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/log"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/pkg/ginx"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"path"
)

const (
	BasePath = "/open/v1"
	DocPath  = "/open/v1/doc"
)

var ApiProviderSet = wire.NewSet(
	user.UserRouterSet,
	wire.Struct(new(Router), "*"),
)

type Router struct {
	User user.APIRouter
}

// SetupOpenAPI initializes app open api router configuration
func SetupOpenAPI(cfg *conf.AppConf, engine *gin.Engine, datasource *data.DataSource) Router {
	if !cfg.ServerConf.OpenAPI {
		return Router{}
	}
	if cfg.ServerConf.Swagger {
		engine.GET(path.Join(DocPath, "*any"), ginSwagger.CustomWrapHandler(Config, swaggerFiles.NewHandler()))
		log.L().Infof("visit OpenAPI Doc on http://%s%s", cfg.ServerConf.HttpConf.Address, path.Join(DocPath, "index.html"))
	}
	root := ginx.NewRouterGroup(engine.RouterGroup.Group(BasePath))

	return setupOpenAPIRouter(root, datasource)
}

var Config = &ginSwagger.Config{
	URL:                      "doc.json",
	DocExpansion:             "list",
	InstanceName:             "openapi",
	Title:                    "Wilson OpenAPI",
	DefaultModelsExpandDepth: 0,
	DeepLinking:              true,
	PersistAuthorization:     false,
	Oauth2DefaultClientID:    "",
}

// swagger declarative api comment

// @title		                    Wilson App Open API Documentation
// @version		                    v1.0.0
// @description                     Wilson open api documentation, to access these open api, you need to add apikey in query param named "key"
// @contact.name                    dstgo
// @contact.url                     https://github.com/dstgo
// @BasePath                        /open/v1
// @license.name                    MIT LICENSE
// @license.url                     https://mit-license.org/
// @securityDefinitions.apikey      ApiKeyAuth
// @in                              query
// @name                            key
//go:generate swag init --generatedTime --instanceName openapi -g api.go -d ./,../types,../core/resp --output ./docs && swag fmt -g api.go -d ./
