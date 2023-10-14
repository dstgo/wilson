//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
	"github.com/jordan-wright/email"
)

//go:generate wire gen
func SetupHandler(appConf *conf.AppConf, api *route.Router, datasource *data.DataSource, issue auth.Issuer, pool *email.Pool) Router {
	panic(wire.Build(HandlerProviderSet))
}
