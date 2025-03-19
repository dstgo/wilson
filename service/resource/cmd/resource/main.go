package main

import (
	"context"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/kratosx/cli"
	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/framework/kratosx/library/logger"
	"github.com/dstgo/wilson/framework/pkg/filex"
	"github.com/dstgo/wilson/service/resource/internal/app"
	"github.com/dstgo/wilson/service/resource/internal/conf"
)

const (
	AppName = "resource"
)

var (
	AppVersion string
)

var service = cli.NewCLI(&cli.Options{
	AppName:     AppName,
	AppVersion:  AppVersion,
	Description: "resource service for wilson framework",
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
	c.ScanWatch("resource", func(value config.Value) {
		if err := value.Scan(&cfg); err != nil {
			panic("watch resource config format error: " + err.Error())
		}
		log.Infof("watch resource config change updated")
	})

	if cfg.Export.LocalDir != "" {

		if !filex.IsDirExist(cfg.Export.LocalDir) {
			_ = os.MkdirAll(cfg.Export.LocalDir, 0655)
			_ = os.MkdirAll(cfg.Export.LocalDir+"/tmp", 0655)
		}

		logger.Helper().Infof("export local dir at %s", cfg.Export.LocalDir)
	}

	app.New(cfg, hs, gs)
}
