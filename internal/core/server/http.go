package server

import (
	"context"
	"github.com/dstgo/wilson/internal/api"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/log"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/handler"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

type Options func(app *HttpServer)

func (o Options) Apply(app *HttpServer) {
	o(app)
}

func WithCtx(ctx context.Context) Options {
	return func(app *HttpServer) {
		app.ctx = ctx
	}
}

func WithConf(appConf *conf.WilsonConf) Options {
	return func(app *HttpServer) {
		app.cfg = appConf
	}
}

func WithLogger(logger *log.Logger) Options {
	return func(app *HttpServer) {
		app.Logger = logger
	}
}

type HttpServer struct {
	ctx context.Context

	cfg    *conf.WilsonConf
	Logger *log.Logger
	Locale *locale.Locale
	server *http.Server
	once   sync.Once

	shutddownFn func()
}

func (a *HttpServer) run() error {
	appConf := a.cfg.ServerConf
	a.Logger.L().Infof("wilson app boot successfully, http server is listenning at %s, tls enable %t", a.server.Addr, appConf.HttpConf.TlsConf.Enable)
	tlsConf := a.cfg.ServerConf.HttpConf.TlsConf
	if tlsConf.Enable {
		return a.server.ListenAndServeTLS(tlsConf.Cert, tlsConf.Pem)
	}
	return a.server.ListenAndServe()
}

func (a *HttpServer) Run() error {
	err := a.run()
	if errors.Is(err, http.ErrServerClosed) {
		a.Logger.L().Infoln("http server closed successfully")
		return nil
	}
	return err
}

func (a *HttpServer) Shutdown() {
	a.once.Do(func() {
		a.Logger.L().Infof("wilson app ready to shutdown")
		a.server.Shutdown(context.Background())
		a.shutddownFn()
	})
}

func NewHTTPApp(options ...Options) (*HttpServer, error) {

	app := new(HttpServer)

	// apply options
	for _, option := range options {
		option.Apply(app)
	}

	if app.cfg == nil {
		return app, errors.New("empty app configuration")
	}

	log.Setup(app.Logger.L())

	var (
		engine     *gin.Engine
		server     *http.Server
		datasource *data.DataSource

		lang = app.Locale

		err error
	)

	if app.Locale == nil {
		// locale
		lang, err = NewLocale(app.cfg.LocaleConf)
		if err != nil {
			return nil, err
		}
		locale.Setup(lang)
		// set validation translator
		vax.SetTranslator(locale.L())
	}

	if err = LogBanner(app.cfg.BuildMeta, app.cfg.LogConf, app.Logger.L(), "wilson.txt"); err != nil {
		return nil, err
	}

	// datasource
	datasource, err = LoadDataSource(app.ctx, app.cfg.DataConf)
	if err != nil {
		return nil, err
	}

	// migrate tables
	if err := entity.Migrate(datasource.ORM()); err != nil {
		return nil, err
	}

	// http server
	engine, server = NewHttpServer(app.cfg, lang, app.Logger.L())

	// setup app handler router
	_, cleanup, err := handler.SetupHandler(app.cfg, engine, datasource)
	if err != nil {
		return nil, err
	}

	// setup app open api router
	_, err = api.SetupOpenAPI(app.cfg, engine, datasource)
	if err != nil {
		return nil, err
	}

	// execute on server shutdown
	shutdownFn := func() {
		if cleanup != nil {
			cleanup()
		}
		CloseDataSource(datasource)
		app.Logger.Close()
	}

	app.server = server
	app.shutddownFn = shutdownFn

	return app, nil
}
