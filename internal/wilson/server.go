package wilson

import (
	"context"
	"github.com/dstgo/wilson/internal/api"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/handler"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/pkg/log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pkg/errors"

	"github.com/dstgo/wilson/pkg/task"

	"github.com/dstgo/wilson/internal/conf"
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

func NewApp(ctx context.Context, cfg *conf.AppConf, loggerw *log.Logger) (*App, error) {

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
	locale.Setup(lang)

	if err = LogBanner(cfg, logger); err != nil {
		return nil, err
	}

	// datasource
	datasource, err = LoadDataSource(ctx, cfg.DataConf)
	if err != nil {
		return nil, err
	}

	// http server
	engine, server = NewHttpServer(cfg, lang, logger)

	// setup app handler router
	_, cleanupHandler, err := handler.SetupHandler(cfg, engine, datasource)
	if err != nil {
		return nil, err
	}

	// setup app open api router
	_ = api.SetupOpenAPI(cfg, engine, datasource)

	// execute on server shutdown
	shutdownFn := func() {
		cleanupHandler()
		CloseDataSource(datasource)
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
