package openapi

import (
	"github.com/dstgo/wilson/app/api/openapi/user"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	BasePath = "/openapi/v1/"
	ApiDoc   = "/opendoc/*any"
)

var ApiSet = wire.NewSet(
	user.UserApiSet,
	wire.Struct(new(ApiRouter), "*"),
)

type ApiRouter struct {
	UserApi user.UserRouter
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

//	@title			Wilson App Open API Documentation
//	@version		v1.0.0
//	@description	Wilson api documentation
//	@BasePath		/openapi/v1/
//go:generate swag init --generatedTime --instanceName openapi -g api.go -d ./,../../types --output ./swagger && && swag fmt -g api.go -d ./
