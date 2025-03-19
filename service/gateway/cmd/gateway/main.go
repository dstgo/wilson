package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	_ "go.uber.org/automaxprocs"

	"github.com/dstgo/wilson/framework/kratosx/cli"
	"github.com/dstgo/wilson/service/gateway/client"
	"github.com/dstgo/wilson/service/gateway/config"
	"github.com/dstgo/wilson/service/gateway/discovery"
	"github.com/dstgo/wilson/service/gateway/middleware"
	"github.com/dstgo/wilson/service/gateway/middleware/circuitbreaker"
	"github.com/dstgo/wilson/service/gateway/proxy"
	"github.com/dstgo/wilson/service/gateway/proxy/debug"
	"github.com/dstgo/wilson/service/gateway/server"

	_ "github.com/dstgo/wilson/service/gateway/discovery/consul"
	_ "github.com/dstgo/wilson/service/gateway/middleware/auth"
	_ "github.com/dstgo/wilson/service/gateway/middleware/bbr"
	_ "github.com/dstgo/wilson/service/gateway/middleware/cors"
	_ "github.com/dstgo/wilson/service/gateway/middleware/logging"
	_ "github.com/dstgo/wilson/service/gateway/middleware/rewrite"
	_ "github.com/dstgo/wilson/service/gateway/middleware/tracing"
	_ "github.com/dstgo/wilson/service/gateway/middleware/transcoder"
)

const (
	AppName = "gateway"
)

var (
	AppVersion string
)

var service = cli.NewCLI(&cli.Options{
	AppName:     AppName,
	AppVersion:  AppVersion,
	Description: "api gateway service for wilson framework",
	StartFn:     Start,
})

func init() {
	service.Parse()
}

func main() {
	service.Start()
}

func Start(opts *cli.StartOptions) error {
	conf, err := config.New(opts.Loader())
	if err != nil {
		return err
	}

	srv, err := NewServer(conf)
	if err != nil {
		return err
	}

	app := kratos.New(
		kratos.Server(srv),
		kratos.ID(opts.ServiceID),
		kratos.Name(opts.AppName),
		kratos.Version(opts.AppVersion),
		kratos.AfterStart(func(ctx context.Context) error {
			log.Infof("gateway started successfully!")
			return nil
		}),
	)

	return app.Run()
}

func NewServer(conf *config.Config) (transport.Server, error) {
	clientFactory := client.NewFactory(makeDiscovery(conf.Discovery))

	pxy, err := proxy.New(clientFactory, middleware.Create)
	if err != nil {
		return nil, fmt.Errorf("failed to new proxy: %v", err)
	}

	circuitbreaker.Init(clientFactory)

	if err = pxy.Update(conf); err != nil {
		return nil, fmt.Errorf("failed to update gateway config: %v", err)
	}
	// 监听配置变化
	conf.WatchEndpoints(func(c *config.Config) {
		if er := pxy.Update(c); er != nil {
			log.Errorf("failed to update gateway config: %v", err)
		}
		log.Infof("watch endpoints config change updated")
	})

	handler := http.Handler(pxy)
	if conf.Debug {
		debug.Register("proxy", pxy)
		handler = debug.MashupWithDebugHandler(pxy)
	}

	return server.NewProxy(handler, conf.Addr), nil
}

func makeDiscovery(dsn string) registry.Discovery {
	if dsn == "" {
		return nil
	}
	d, err := discovery.Create(dsn)
	if err != nil {
		log.Fatalf("failed to create discovery: %v", err)
	}
	return d
}
