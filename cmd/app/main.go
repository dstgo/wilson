package main

import (
	"context"
	"flag"
	"github.com/dstgo/filebox"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/server"
	"github.com/dstgo/wilson/pkg/coco"
	"path"
)

var (
	// specified config file path
	configFile string
)

func init() {
	flag.StringVar(&configFile, "conf", path.Join(filebox.GetCurrentRunningPath(), "config.yaml"), "config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()
	// read configuration
	confFile := coco.NewConfigFile(configFile)
	if err := confFile.ReadConfig(); err != nil {
		panic(err)
	}
	// map configuration struct
	appConf, err := conf.NewAppConf(confFile)
	if err != nil {
		panic(err)
	}
	coco.SetMode(appConf.AppConf.Mode)
	// new app
	app, err := server.NewApp(ctx, appConf)
	if err != nil {
		panic(err)
	}
	// run app
	app.Core().L().Infoln(app.Run())
}
