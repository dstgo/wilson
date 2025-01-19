package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/dstgo/wilson/framework/cli"
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/service/manager/internal/app"
	"github.com/dstgo/wilson/service/manager/internal/conf"
)

const AppName = "manager"

var (
	AppVersion string
)

var service = cli.NewCLI(&cli.Options{
	AppName:     AppName,
	AppVersion:  AppVersion,
	Description: "manager service for wilson framework",
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

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	cfg := &conf.Config{}
	c.ScanWatch("manager", func(value config.Value) {
		if err := value.Scan(cfg); err != nil {
			panic("manager config format error: " + err.Error())
		}
	})

	app.New(cfg, hs, gs)
}
