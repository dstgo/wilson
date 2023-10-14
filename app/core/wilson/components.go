package wilson

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dstgo/size"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/log"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/pkg/sysinfo"
	"github.com/dstgo/wilson/assets"
	"github.com/dstgo/wilson/pkg/route"
	"github.com/jordan-wright/email"
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
	memInfo := sysinfo.GetMemInfo()

	bannerData := map[string]any{
		"author":     cfg.ServerConf.Author,
		"timezone":   time.Now().Format("MST -07"),
		"appName":    cfg.ServerConf.Name,
		"goVersion":  runtime.Version(),
		"appMode":    strings.ToUpper(cfg.ServerConf.Mode),
		"appVersion": cfg.ServerConf.Version,
		"osInfo":     fmt.Sprintf("%s %s %s", hostInfo.Os, hostInfo.Platform, hostInfo.Version),
		"archInfo":   runtime.GOARCH,
		"cpuInfo":    fmt.Sprintf("%s %d Cores", cpuInfo.Name, cpuInfo.Count),
		"memInfo":    size.ParseTargetSize(memInfo.Virtual.Total.String(), size.GB),
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

func LoadEmailPool(emailConf *conf.EmailConf) (*email.Pool, error) {
	pool, err := email.NewPool(emailConf.Address(), emailConf.MaxPoolSize, emailConf.SmtpAuth())
	if err != nil {
		log.L().Errorf("load email pool failed: %s", emailConf.Address())
		return nil, err
	}
	log.L().Infof("load email pool ok: %s", emailConf.Address())
	return pool, nil
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
