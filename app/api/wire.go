//go:build wireinject
// +build wireinject

package api

import (
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/google/wire"
)

//go:generate wire gen
func SetupAPI(open *route.Router, datasource *data.DataSource) Router {
	panic(wire.Build(ApiProviderSet))
}
