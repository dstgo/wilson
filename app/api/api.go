package api

import (
	"github.com/dstgo/wilson/app/api/systemApi"
	"github.com/dstgo/wilson/app/api/userApi"
	"github.com/google/wire"
)

// ApiSet api provider set
var ApiSet = wire.NewSet(
	systemApi.SystemApiSet,
	userApi.UserApiSet,
	wire.Struct(new(ApiRouter), "*"),
)

// ApiRouter
// combination of all router
type ApiRouter struct {
	SystemApi *systemApi.Router
	UserApi   *userApi.Router
}

// swagger declarative api comment

// @title Wilson API Documentation
// @version v1.0.0
// @description Wilson api documentation
// @BasePath /api/v1
//go:generate swag init -g api.go -d ./ --output ./swagger
