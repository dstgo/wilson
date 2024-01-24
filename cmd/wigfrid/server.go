package main

import (
	"context"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/server"
	"github.com/dstgo/wilson/pkg/config"
	"github.com/go-kratos/kratos/v2"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

var configFile string

var serverCmd = &cobra.Command{
	Use:          "server",
	Short:        "run wigfrid daemon",
	Long:         "run wigfrid daemon",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serve(configFile, Author, Version, BuildTime)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	serverCmd.Flags().StringVarP(&configFile, "config", "f", "config.yaml", "specified config file")
}

func serve(configFile string, author string, version string, buildTime string) error {
	// listen signal
	signalCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGABRT, syscall.SIGTERM)
	defer cancel()
	app, err := newServer(signalCtx, configFile, author, version, buildTime)
	if err != nil {
		return err
	}
	return app.Run()
}

func newServer(ctx context.Context, configFile string, author string, version string, buildTime string) (*kratos.App, error) {
	// read configuration
	appConfig := config.NewConfigFile(configFile)
	if err := appConfig.ReadConfig(); err != nil {
		return nil, err
	}

	wigfridConf, err := conf.NewWigfridConf(appConfig, conf.BuildInfo{
		Author:    author,
		Version:   version,
		BuildTime: buildTime,
	})
	if err != nil {
		return nil, err
	}

	logger, err := server.NewLogger(wigfridConf.LogConf)
	if err != nil {
		return nil, err
	}

	app, err := server.NewGrpcApp(ctx, wigfridConf, logger)
	if err != nil {
		return nil, err
	}

	return app, nil
}
