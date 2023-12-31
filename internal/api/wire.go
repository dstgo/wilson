//go:build wireinject
// +build wireinject

package api

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/pkg/ginx"
	"github.com/google/wire"
)

//go:generate wire gen
func setupOpenAPIRouter(open *ginx.RouterGroup, datasource *data.DataSource) Router {
	panic(wire.Build(ApiProviderSet))
}
