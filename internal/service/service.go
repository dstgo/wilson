package service

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/service/archive"
	"github.com/dstgo/wilson/internal/service/control"
	"github.com/dstgo/wilson/internal/service/cron"
	"github.com/dstgo/wilson/internal/service/daemon"
	"github.com/dstgo/wilson/internal/service/mod"
	"github.com/dstgo/wilson/internal/service/player"
	"github.com/dstgo/wilson/internal/service/setting"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
)

type RegisteredService struct {
	daemon  *daemon.DaemonService
	control *control.ControlService
	player  *player.PlayerService
	setting *setting.SettingService
	archive *archive.ArchiveService
	cron    *cron.CronService
	mod     *mod.ModService
}

// Register registers the grpc service to the  gpc-server
func Register(grpcServer *grpc.Server, cfg *conf.WigfridConf, datasource *data.DataSource, logger log.Logger) {

}
