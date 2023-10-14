package handler

import (
	"github.com/dstgo/wilson/app/handler/auth"
	"github.com/dstgo/wilson/app/handler/email"
	"github.com/dstgo/wilson/app/handler/system"
	"github.com/dstgo/wilson/app/handler/user"
	"github.com/google/wire"
	ginSwagger "github.com/swaggo/gin-swagger"
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
