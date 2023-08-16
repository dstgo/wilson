package wilson

import (
	"context"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pkg/errors"

	"github.com/dstgo/wilson/pkg/task"

	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/core/log"
	"github.com/dstgo/wilson/app/data"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	server      *http.Server
	logger      *logrus.Logger
	cfg         *conf.AppConf
	once        sync.Once
	shutddownFn func()
}

func (a *App) run() error {
	appConf := a.cfg.ServerConf
	a.logger.Infof("wilson app boot successfully, http server is listenning at %s, tls enable %t", a.server.Addr, appConf.HttpConf.TlsConf.Enable)
	tlsConf := a.cfg.ServerConf.HttpConf.TlsConf
	if tlsConf.Enable {
		return a.server.ListenAndServeTLS(tlsConf.Cert, tlsConf.Pem)
	}
	return a.server.ListenAndServe()
}

func (a *App) Run(ctx context.Context) error {
	c, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGKILL, syscall.SIGABRT, syscall.SIGTERM)

	bootTask := task.NewTask(ctx)

	bootTask.AddJobs(func(ctx context.Context) error {
		err := a.run()
		stop()
		if errors.Is(err, http.ErrServerClosed) {
			a.logger.Infoln("http server closed successfully")
			return nil
		}
		return err
	})

	bootTask.AddJobs(func(ctx context.Context) error {
		select {
		case <-c.Done():
			a.Shutdown()
		}
		return nil
	})

	return bootTask.Run()
}

func (a *App) Shutdown() {
	a.once.Do(func() {
		a.logger.Infof("wilson app ready to shutdown")
		a.server.Shutdown(context.Background())
		a.shutddownFn()
	})
}

func NewApp(ctx context.Context, cfg *conf.AppConf, loggerw *log.LoggerW) (*App, error) {

	var (
		lang       *locale.Locale
		engine     *gin.Engine
		server     *http.Server
		datasource *data.DataSource
		err        error
		logger     = loggerw.L()
	)

	// locale
	lang, err = NewLocale(cfg.LocaleConf)
	if err != nil {
		return nil, err
	}

	if err = LogBanner(cfg, logger); err != nil {
		return nil, err
	}

	// datasource
	datasource, err = LoadDataSource(ctx, cfg.DataConf, logger)
	if err != nil {
		return nil, err
	}

	// http server
	engine, server = NewHttpServer(cfg, lang, logger)

	// register app api router
	_ = NewAppApiRouter(cfg, lang, engine, datasource)

	// register open api router
	_ = NewOpenApiRouter(cfg, lang, engine, datasource)

	// execute on server shutdown
	shutdownFn := func() {
		CloseDataSource(datasource, logger)
		loggerw.Close()
	}

	app := &App{
		server:      server,
		logger:      logger,
		cfg:         cfg,
		shutddownFn: shutdownFn,
	}

	return app, nil
}
