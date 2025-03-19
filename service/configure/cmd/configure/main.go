package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	thttp "github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/kratosx/cli"
	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/framework/kratosx/library/logger"
	"github.com/dstgo/wilson/framework/pkg/webserver"
	"github.com/dstgo/wilson/service/configure/internal/app"
	"github.com/dstgo/wilson/service/configure/internal/conf"
)

const AppName = "configure"

var (
	AppVersion string
)

var service = cli.NewCLI(&cli.Options{
	AppName:     AppName,
	AppVersion:  AppVersion,
	Description: "configure service for wilson framework",
	StartFn:     Start,
})

func init() {
	service.Parse()
}

func main() {
	service.Start()
}

func Start(opts *cli.StartOptions) error {
	server := kratosx.New(
		kratosx.ID(opts.ServiceID),
		kratosx.Name(opts.AppName),
		kratosx.Version(opts.AppVersion),
		kratosx.Config(opts.Loader()),
		kratosx.RegistrarServer(RegisterServer),
		kratosx.Options(
			kratos.AfterStart(func(ctx context.Context) error {
				kt := kratosx.MustContext(ctx)
				kt.Logger().Infof("service %s started successfully!", kt.ID())
				return nil
			}),
		),
	)

	return server.Run()
}

func RegisterServer(c config.Config, hs *thttp.Server, gs *grpc.Server) {
	cfg := &conf.Config{}

	// watch config
	c.ScanWatch("configure", func(value config.Value) {
		if err := value.Scan(cfg); err != nil {
			log.Printf("configure config format error: %s", err.Error())
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
