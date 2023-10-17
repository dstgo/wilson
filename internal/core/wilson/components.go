package wilson

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dstgo/wilson/assets"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/log"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/dstgo/wilson/pkg/sysinfo"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"text/template"
	"time"
)

// on server boot hooks

func LogBanner(cfg *conf.AppConf, logger *logrus.Logger) error {
	bannerTemplate := bytes.NewBuffer(nil)

	banner, err := template.ParseFS(assets.Fs, "banner.txt")
	if err != nil {
		return err
	}

	hostInfo := sysinfo.GetHostInfo()
	cpuInfo := sysinfo.GetCpuInfo()

	bannerData := map[string]any{
		"author":    cfg.ServerConf.Author,
		"appName":   cfg.ServerConf.Name,
		"appMode":   strings.ToUpper(cfg.ServerConf.Mode),
		"logMode":   strings.ToUpper(cfg.LogConf.Level),
		"goVersion": runtime.Version(),
		"version":   cfg.ServerConf.Version,
		"osInfo":    fmt.Sprintf("%s %s", hostInfo.Os, hostInfo.Version),
		"timezone":  time.Now().Format("MST -07"),
		"archInfo":  runtime.GOARCH,
		"cpuInfo":   cpuInfo.Name,
	}

	if err := banner.Execute(bannerTemplate, bannerData); err != nil {
		return err
	}

	logger.Infoln(fmt.Sprintf("\n\n%s", bannerTemplate.String()))

	return nil
}

func LoadDataSource(ctx context.Context, dataConf *conf.DataConf) (*data.DataSource, error) {
	log.L().Infoln("attempt to load wilson datasource...")
	datasource, err := data.NewDataSource(ctx, dataConf)
	if err != nil {
		log.L().Errorf("load data datasource failed: %s", err.Error())
		return datasource, err
	}
	log.L().Infof("load data datasource ok √")
	return datasource, nil
}

func DebugPrintRouter(router *route.Router) error {
	return router.Walk(func(info route.RouterInfo) error {
		if !info.IsGroup {
			log.L().
				WithField("method", info.Method).
				WithField("path", info.FullPath).
				Debugln()
		}
		return nil
	})
}

// on server shutdown hooks

func CloseDataSource(datasource *data.DataSource) {
	if datasource != nil {
		// close datasource
		if err := datasource.Close(); err != nil {
			log.L().Errorf("data source closed failed: %s", err)
			return
		}
	}
	log.L().Infoln("data source closed successfully")
}
