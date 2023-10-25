package cmd

import (
	"context"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/log"
	"github.com/dstgo/wilson/internal/core/wilson"
	"github.com/dstgo/wilson/pkg/config"
	"github.com/dstgo/wilson/pkg/task"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log/slog"
	"os/signal"
	"path"
	"syscall"
)

var (
	configFile string
	initial    bool
)

var serverCmd = &cobra.Command{
	Use:     "server [--f filename]",
	Short:   "Run wilson backend server",
	Example: "wilson server --f /etc/wilson/config.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serve(configFile, Author, Version); err != nil {
			slog.Error(errors.Wrap(err, "wilson server running failed").Error())
		}
	},
}

func init() {
	serverCmd.Flags().StringVar(&configFile, "f", path.Join(DefaultDir, "config.yaml"), "specified wilson server config file")
	serverCmd.Flags().BoolVar(&initial, "i", false, "only initial server data, not run web server")
}

func newServer(ctx context.Context, configFile string, author string, version string) (*wilson.App, error) {

	// read configuration
	appConfig := config.NewConfigFile(configFile)
	if err := appConfig.ReadConfig(); err != nil {
		return nil, err
	}

	// map configuration struct
	appConf, err := conf.NewAppConf(appConfig, author, version)
	if err != nil {
		return nil, err
	}

	// ini logger
	logger, err := wilson.NewLogger(appConf.LogConf)
	if err != nil {
		return nil, err
	}
	log.Setup(logger.L())

	// set app mode
	gin.SetMode(appConf.ServerConf.Mode)

	// initialize app server
	app, err := wilson.NewApp(
		wilson.WithCtx(ctx),
		wilson.WithConf(appConf),
		wilson.WithLogger(logger),
	)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func serve(configFile string, author string, version string) error {

	// listen signal
	signalCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGABRT, syscall.SIGTERM)

	// create app
	app, err := newServer(signalCtx, configFile, author, version)
	if err != nil {
		return err
	}

	serverTask, causeFunc := task.New(signalCtx)
	defer causeFunc(nil)

	serverTask.Add(func(ctx context.Context) error {
		defer cancel()
		if initial {
			return nil
		}
		// run the http server
		return app.Run()
	})

	serverTask.Add(func(ctx context.Context) error {
		select {
		case <-signalCtx.Done():
			app.Shutdown()
		}
		return nil
	})

	return serverTask.Run()
}
