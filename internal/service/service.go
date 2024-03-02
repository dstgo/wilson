package service

import (
	"github.com/246859/steamapi"
	"github.com/docker/docker/client"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/service/archive"
	"github.com/dstgo/wilson/internal/service/cronjob"
	"github.com/dstgo/wilson/internal/service/daemon"
	"github.com/dstgo/wilson/internal/service/mod"
	"github.com/dstgo/wilson/internal/service/player"
	"github.com/dstgo/wilson/internal/service/server"
	"github.com/dstgo/wilson/internal/service/setting"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

var serviceProviderSet = wire.NewSet(
	wire.Struct(new(RegisteredService), "*"),
)

type RegisteredService struct {
	daemon  *daemon.Service
	server  *server.Service
	player  *player.Service
	setting *setting.Service
	archive *archive.Service
	cron    *cronjob.Service
	mod     *mod.Service
}

// Register registers the grpc service to the gpc-server
func Register(grpcServer *grpc.Server, cfg *conf.WigfridConf, datasource *data.DataSource, dockerClient *client.Client, steamClient *steamapi.Client, logger log.Logger) {

}
