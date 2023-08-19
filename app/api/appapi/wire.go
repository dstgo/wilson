//go:build wireinject
// +build wireinject

package appapi

import (
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/auth"
	"github.com/dstgo/wilson/app/dao"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/logic"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
	"github.com/jordan-wright/email"
)

//go:generate wire gen
func NewApiRouter(appConf *conf.AppConf, rootRouter *route.Router, datasource *data.DataSource, issue auth.Issuer, pool *email.Pool) ApiRouter {
	panic(wire.Build(dao.AppDaoSet, logic.AppLogicSet, ApiSet))
}
