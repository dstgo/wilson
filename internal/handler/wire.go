//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

//go:generate wire gen
func setupHandlerRouter(appConf *conf.AppConf, api *route.Router, datasource *data.DataSource) (Router, func(), error) {
	panic(wire.Build(HandlerProviderSet))
}
