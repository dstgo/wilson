package server

import (
	"context"
	"github.com/dstgo/wilson/app/api"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/pkg/locale"
	"github.com/dstgo/wilson/app/pkg/logw"
	"github.com/dstgo/wilson/pkg/coco"
	"github.com/dstgo/wilson/pkg/coco/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Components struct is component struct wrapper
type Components struct {
	Conf       *conf.AppConf
	Engine     *gin.Engine
	Logger     *logw.LoggerW
	Server     *http.Server
	Router     *route.Router
	Datasource *data.DataSource
	Lang       *locale.Locale
}

type WilsonApp struct {
	Api        api.ApiRouter
	Components *Components
	*coco.Coco
}

func (a *WilsonApp) Run() error {
	tlsConf := a.Components.Conf.AppConf.Http.TlsConf
	if tlsConf.Enable {
		return a.Coco.RunTLS(tlsConf.Cert, tlsConf.Pem)
	}
	return a.Coco.Run()
}

func NewApp(ctx context.Context, cfg *conf.AppConf) (*WilsonApp, error) {

	components := new(Components)

	// config
	components.Conf = cfg

	// logger
	newLogger, err := newLogger(cfg.LogConf)
	if err != nil {
		return nil, err
	}
	components.Logger = newLogger

	// locale
	l, err := newLocale(cfg.LocaleConf)
	if err != nil {
		return nil, err
	}
	components.Lang = l

	// gin engine
	components.Engine = newEngine(cfg.AppConf, l)
	// http server
	components.Server = newHttpServer(cfg.AppConf)
	// root router
	components.Router = route.NewRouter(components.Engine)

	// new coco
	c := coco.New(
		ctx,
		coco.WithConfig(cfg),
		coco.WithLogger(components.Logger.L()),
		// http part
		coco.WithEngine(components.Engine),
		coco.WithServer(components.Server),
		coco.WithRouter(components.Router),
	)

	c.OnPanic = func(err any) {
		panic(err)
	}

	// sync boot task
	c.AddPreSyncCs(
		LogBanner(cfg),
	)

	// async boot task
	c.AddPreAsyncCs(
		LoadDataSource(cfg.DataConf, components),
	)

	// after boot task
	c.AddPostAsyncCs(
		BootLog(cfg),
		DebugPrintRouter(),
	)

	// graceful shutdown hook
	c.OnInterrupt(
		ShutdownWithInfo(),
		CloseDataSource(components.Datasource),
		CloseLogger(components.Logger),
		ShutdownWithCloseHttp(),
	)

	// attach api router
	apiRouter, err := attachRouter(components)
	if err != nil {
		return nil, err
	}

	app := &WilsonApp{
		Components: components,
		Api:        apiRouter,
		Coco:       c,
	}

	return app, nil
}
