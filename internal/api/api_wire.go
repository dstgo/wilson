//go:build wireinject
// +build wireinject

package api

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/dao"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/logic"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/pkg/coco"
	"github.com/google/wire"
)

//go:generate wire gen -output_file_prefix api_
func NewApiRouter(config *conf.AppConf, co *coco.Core, locale *locale.Locale, datasource *data.DataSource) (ApiRouter, error) {
	panic(wire.Build(dao.DaoProviderSet, logic.LogicSet, ApiSet, wire.FieldsOf(new(*conf.AppConf), "App")))
}
