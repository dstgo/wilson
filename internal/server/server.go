package server

import (
	"context"
	"github.com/dstgo/filebox"
	"github.com/dstgo/wilson/internal/api"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/pkg/logw"
	"github.com/dstgo/wilson/pkg/coco"
	"github.com/dstgo/wilson/pkg/coco/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WilsonApp struct {
	Conf *conf.AppConf
	Api  api.ApiRouter
	*coco.Coco
}

func (a *WilsonApp) Run() error {
	tlsConf := a.Conf.App.Http.TlsConf
	a.Core().L().Infof("http server listening bind %s at %s mode", a.Core().Server().Addr, coco.Mode())
	a.Core().L().Infof("working directory at %s", filebox.GetCurrentRunningPath())
	if tlsConf.Enable {
		a.Core().L().Infof("tls is enable")
		return a.Coco.RunTLS(tlsConf.Cert, tlsConf.Pem)
	}
	return a.Coco.Run()
}

func NewApp(ctx context.Context, cfg *conf.AppConf) (*WilsonApp, error) {

	var (
		engine     *gin.Engine
		logger     *logw.LoggerW
		server     *http.Server
		router     *route.Router
		datasource *data.DataSource
		lang       *locale.Locale
	)

	// gin engine
	engine = newEngine(cfg.App)
	// http server
	server = newHttpServer(cfg.App)
	// root router
	router = route.NewRouter(engine)
	// logger
	newLogger, err := logw.NewLogger(cfg.Log)
	if err != nil {
		return nil, err
	}
	logger = newLogger

	// new coco
	c := coco.New(
		coco.WithCtx(ctx),
		coco.WithConfig(cfg),
		coco.WithLogger(logger.L()),
		// http part
		coco.WithEngine(engine),
		coco.WithServer(server),
		coco.WithRouter(router),
	)

	c.OnPanic = func(err any) {
		panic(err)
	}

	// boot task
	c.AddPreAsyncCs(
		LoadLangDir(cfg.Locale, &lang),
		LoadDataSource(cfg.Data, &datasource),
	)

	// graceful shutdown hook
	c.OnInterrupt(OnShutdown(datasource, logger))
	// attach api router
	apiRouter, err := attachRouter(cfg, c.Core(), lang, datasource)
	if err != nil {
		return nil, err
	}

	app := &WilsonApp{
		Conf: cfg,
		Api:  apiRouter,
		Coco: c,
	}

	return app, nil
}
