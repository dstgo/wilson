package wilson

import (
	"context"
	"github.com/dstgo/wilson/internal/api"
	"github.com/dstgo/wilson/internal/core/log"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/handler"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/pkg/errors"
	"net/http"
	"sync"

	"github.com/dstgo/wilson/internal/conf"
	"github.com/gin-gonic/gin"
)

type Options func(app *App)

func (o Options) Apply(app *App) {
	o(app)
}

func WithCtx(ctx context.Context) Options {
	return func(app *App) {
		app.ctx = ctx
	}
}

func WithConf(appConf *conf.AppConf) Options {
	return func(app *App) {
		app.cfg = appConf
	}
}

func WithLogger(logger *log.Logger) Options {
	return func(app *App) {
		app.Logger = logger
	}
}

type App struct {
	ctx context.Context

	cfg    *conf.AppConf
	Logger *log.Logger
	Locale *locale.Locale
	server *http.Server
	once   sync.Once

	shutddownFn func()
}

func (a *App) run() error {
	appConf := a.cfg.ServerConf
	a.Logger.L().Infof("wilson app boot successfully, http server is listenning at %s, tls enable %t", a.server.Addr, appConf.HttpConf.TlsConf.Enable)
	tlsConf := a.cfg.ServerConf.HttpConf.TlsConf
	if tlsConf.Enable {
		return a.server.ListenAndServeTLS(tlsConf.Cert, tlsConf.Pem)
	}
	return a.server.ListenAndServe()
}

func (a *App) Run() error {
	err := a.run()
	if errors.Is(err, http.ErrServerClosed) {
		a.Logger.L().Infoln("http server closed successfully")
		return nil
	}
	return err
}

func (a *App) Shutdown() {
	a.once.Do(func() {
		a.Logger.L().Infof("wilson app ready to shutdown")
		a.server.Shutdown(context.Background())
		a.shutddownFn()
	})
}

func NewApp(options ...Options) (*App, error) {

	app := new(App)

	// apply options
	for _, option := range options {
		option.Apply(app)
	}

	if app.cfg == nil {
		return app, errors.New("empty app configuration")
	}

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

	if err = LogBanner(app.cfg, app.Logger.L()); err != nil {
		return nil, err
	}

	// datasource
	datasource, err = LoadDataSource(app.ctx, app.cfg.DataConf)
	if err != nil {
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
