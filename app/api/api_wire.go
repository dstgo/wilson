//go:build wireinject
// +build wireinject

package api

import (
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/dao"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/logic"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

//go:generate wire gen -output_file_prefix api_
func NewApiRouter(appConf *conf.AppConf, rootRouter *route.Router, datasource *data.DataSource) ApiRouter {
	panic(wire.Build(dao.DaoProviderSet, logic.LogicSet, ApiSet))
}
