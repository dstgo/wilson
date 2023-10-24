package cmd

import (
	"context"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/log"
	"github.com/dstgo/wilson/internal/core/wilson"
	"github.com/dstgo/wilson/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log/slog"
	"net/http"
	"path"
)

var serverCmd = &cobra.Command{
	Use:     "server [--f filename]",
	Short:   "Run wilson backend server",
	Example: "wilson server --f /etc/wilson/config.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		if err := server(configFile, Author, Version); err != nil {
			slog.Error(errors.Wrap(err, "wilson server running failed").Error())
		}
	},
}

func init() {
	serverCmd.Flags().StringVar(&configFile, "f", path.Join(DefaultDir, "config.yaml"), "specified wilson server config file")
}

func server(configFile string, author string, version string) error {
	ctx := context.Background()

	// read configuration
	appConfig := config.NewConfigFile(configFile)
	if err := appConfig.ReadConfig(); err != nil {
		return err
	}

	// map configuration struct
	appConf, err := conf.NewAppConf(appConfig, author, version)
	if err != nil {
		return err
	}

	// ini logger
	logger, err := wilson.NewLogger(appConf.LogConf)
	if err != nil {
		return err
	}
	log.Setup(logger.L())

	// set app mode
	gin.SetMode(appConf.ServerConf.Mode)

	// initialize app server
	app, err := wilson.NewApp(ctx, appConf, logger)
	if err != nil {
		return err
	}

	if err := app.Run(ctx); errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}
