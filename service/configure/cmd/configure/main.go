package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	thttp "github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/dstgo/wilson/api/rpc/configure"
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/kratosx/config"
	"github.com/dstgo/wilson/framework/kratosx/library/md"
	"github.com/dstgo/wilson/framework/pkg/webserver"

	"github.com/dstgo/wilson/service/configure/internal/app"
	"github.com/dstgo/wilson/service/configure/internal/conf"
)

var (
	Name      string
	Version   string
	BuildTime string
)

func main() {
	server := kratosx.New(
		kratosx.Config(configure.NewFromEnv()),
		kratosx.RegistrarServer(RegisterServer),
		kratosx.Options(
			kratos.Name(Name),
			kratos.Version(Version),
			kratos.Metadata(map[string]string{
				md.ServiceAppName:    Name,
				md.ServiceAppVersion: Version,
				md.ServiceBuildTime:  BuildTime,
			}),
			kratos.AfterStart(func(ctx context.Context) error {
				kt := kratosx.MustContext(ctx)
				fmt.Printf("hello %s !\n", kt.Name())
				return nil
			}),
		),
	)

	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func RegisterServer(c config.Config, hs *thttp.Server, gs *grpc.Server) {
	cfg := &conf.Config{}
	c.ScanWatch("business", func(value config.Value) {
		if err := value.Scan(cfg); err != nil {
			panic("business config format error: " + err.Error())
		}
	})
	go webserver.Run("web/dist/", cfg.WebUI.Addr, map[string]any{
		"Port": c.App().Server.Http.Port,
	})
	app.New(cfg, hs, gs)
}
