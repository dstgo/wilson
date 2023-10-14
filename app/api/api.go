package api

import (
	"github.com/dstgo/wilson/app/api/user"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	BasePath = "/open"
	DocPath  = "/open/doc"
)

var ApiProviderSet = wire.NewSet(
	user.UserRouterSet,
	wire.Struct(new(Router), "*"),
)

type Router struct {
	User user.APIRouter
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

//	@title			Wilson App Internal API Documentation
//	@version		v1.0.0
//	@description	Wilson api documentation
//	@BasePath		/openapi/
//go:generate swag init --generatedTime --instanceName openapi -g api.go -d ./,../types,../core/resp --output ./swagger && swag fmt -g api.go -d ./
