package main

import (
	"context"
	"flag"
	"github.com/dstgo/wilson/app/core/log"
	"path"

	"github.com/dstgo/filebox"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/wilson"
	"github.com/dstgo/wilson/pkg/config"
	"github.com/gin-gonic/gin"
)

var (
	// specified config file path
	configFile string
	Author     string
	Version    string
)

func init() {
	flag.StringVar(&configFile, "conf", path.Join(filebox.GetCurrentRunningPath(), "config.yaml"), "app configuration file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// read configuration
	appConfig := config.NewConfigFile(configFile)
	if err := appConfig.ReadConfig(); err != nil {
		panic(err)
	}

	// map configuration struct
	appConf, err := conf.NewAppConf(appConfig, Author, Version)
	if err != nil {
		panic(err)
	}

	// ini logger
	logger, err := wilson.NewLogger(appConf.LogConf)
	if err != nil {
		panic(err)
	}
	log.Setup(logger.L())

	// set app mode
	gin.SetMode(appConf.ServerConf.Mode)

	// initialize app server
	app, err := wilson.NewApp(ctx, appConf, logger)
	if err != nil {
		panic(err)
	}

	app.Run(ctx)
}
