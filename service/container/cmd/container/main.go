package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	thttp "github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/dstgo/wilson/framework/cli"
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/service/container/internal/app"
	"github.com/dstgo/wilson/service/container/internal/conf"
)

const AppName = "container"

var (
	AppVersion string
)

var service = cli.NewCLI(&cli.Options{
	AppName:     AppName,
	AppVersion:  AppVersion,
	Description: "container service for wilson framework",
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
	c.ScanWatch("container", func(value config.Value) {
		if err := value.Scan(cfg); err != nil {
			log.Printf("container config format error: %s", err.Error())
		}
	})

	app.New(cfg, hs, gs)
}
