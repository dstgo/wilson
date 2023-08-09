//go:build wireinject
// +build wireinject

package api

import (
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/dao"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/logic"
	"github.com/dstgo/wilson/app/pkg/locale"
	"github.com/dstgo/wilson/pkg/coco/route"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

//go:generate wire gen -output_file_prefix api_
func NewApiRouter(appConf *conf.AppConf, rootRouter *route.Router, logger *logrus.Logger,
	lang *locale.Locale, datasource *data.DataSource) (ApiRouter, error) {
	panic(wire.Build(dao.DaoProviderSet, logic.LogicSet, ApiSet))
}
