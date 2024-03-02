//go:build wireinject
// +build wireinject

package service

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func setupService(cfg *conf.WigfridConf, datasource *data.DataSource, dockerClient *client.Client, steamClient *steamapi.Client, logger log.Logger) RegisteredService {
	panic(wire.Build())
}
