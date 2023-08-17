package appapi

import (
	"github.com/dstgo/wilson/app/api/appapi/system"
	"github.com/dstgo/wilson/app/api/appapi/user"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	BasePath = "/api/v1/"
	ApiDoc   = "/appdoc/*any"
)

// ApiSet api provider set
var ApiSet = wire.NewSet(
	system.SystemApiSet,
	user.UserApiSet,
	wire.Struct(new(ApiRouter), "*"),
)

// ApiRouter
// combination of all router
type ApiRouter struct {
	SystemApi system.Router
	UserApi   user.Router
}

var Config = &ginSwagger.Config{
	URL:                      "doc.json",
	DocExpansion:             "list",
	InstanceName:             "appapi",
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
//	@BasePath		/api/v1
//go:generate swag init --generatedTime --instanceName appapi -g api.go -d ./,../../types --output ./swagger && swag fmt -g api.go -d ./
