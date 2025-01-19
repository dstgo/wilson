package main

import (
	"log"

	"github.com/go-kratos/kratos/v2"
	configfile "github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	thttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/samber/lo"
	_ "go.uber.org/automaxprocs"

	"github.com/dstgo/wilson/framework/cli"
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/framework/kratosx/library/logger"
	"github.com/dstgo/wilson/framework/kratosx/library/md"
	"github.com/dstgo/wilson/framework/pkg/webserver"
	"github.com/dstgo/wilson/service/configure/internal/app"
	"github.com/dstgo/wilson/service/configure/internal/conf"
)

var (
	AppName      string
	AppVersion   string
	AppBuildTime string
)

var service cli.Service

func init() {
	opts := &cli.Options{
		AppName:      lo.Ternary(AppName != "", AppName, "service-configure"),
		AppVersion:   AppVersion,
		AppBuildTime: AppBuildTime,
		Description:  "service for wilson configure",
		StartFn:      Start,
	}

	service = cli.NewCLI(opts)

	service.Parse()
}

func main() {
	service.Start()
}

func Start(opts *cli.Options) error {
	server := kratosx.New(
		kratosx.Config(configfile.NewSource(opts.ConfigFile)),
		kratosx.RegistrarServer(RegisterServer),
		kratosx.Options(
			kratos.Metadata(map[string]string{
				md.ServiceAppName:    opts.AppName,
				md.ServiceAppVersion: opts.AppVersion,
				md.ServiceBuildTime:  opts.AppBuildTime,
			}),
		),
	)

	return server.Run()
}

func RegisterServer(c config.Config, hs *thttp.Server, gs *grpc.Server) {
	cfg := &conf.Config{}

	// watch config
	c.ScanWatch("business", func(value config.Value) {
		if err := value.Scan(cfg); err != nil {
			log.Printf("business config format error: %s", err.Error())
		}
	})

	if cfg.WebUI.Enable {
		go func() {
			err := webserver.ServeDir(cfg.WebUI.Dist, cfg.WebUI.Addr, map[string]any{
				"Port": c.App().Server.Http.Port,
			})

			if err != nil {
				logger.Helper().Error(err)
			}
		}()
	}

	app.New(cfg, hs, gs)
}
