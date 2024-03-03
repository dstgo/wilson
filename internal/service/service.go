package service

import (
	"github.com/246859/steamapi"
	"github.com/docker/docker/client"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data"
	v1 "github.com/dstgo/wilson/internal/proto/api/v1"
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
	archive.ArchiveProvider,
	cronjob.CronJobProvider,
	daemon.DaemonProvider,
	mod.ModProvider,
	player.PlayerProvider,
	server.ServerProvider,
	setting.SettingProvider,
)

// RegisteredService includes all daemon services
type RegisteredService struct {
	daemon  *daemon.Service
	server  *server.Service
	player  *player.Service
	setting *setting.Service
	archive *archive.Service
	cron    *cronjob.Service
	mod     *mod.Service
}

func registerServiceV1(grpcServer *grpc.Server, services RegisteredService) {
	v1.RegisterArchiveServiceServer(grpcServer, services.archive)
	v1.RegisterCronJobServiceServer(grpcServer, services.cron)
	v1.RegisterDaemonServiceServer(grpcServer, services.daemon)
	v1.RegisterModServiceServer(grpcServer, services.mod)
	v1.RegisterPlayerServiceServer(grpcServer, services.player)
	v1.RegisterDstServerServiceServer(grpcServer, services.server)
	v1.RegisterSettingServiceServer(grpcServer, services.setting)
}

// Register registers the grpc service to the gpc-server
func Register(grpcServer *grpc.Server, cfg *conf.WigfridConf, datasource *data.DataSource, dockerClient *client.Client, steamClient *steamapi.Client, logger log.Logger) error {
	// inject the service provider
	services := setupService(cfg, datasource, dockerClient, steamClient, logger)
	// register grpc service
	registerServiceV1(grpcServer, services)

	return nil
}
