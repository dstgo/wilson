package server

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/pkg/logw"
	"github.com/dstgo/wilson/pkg/coco"
	"os"
)

func OnShutdown(datasource *data.DataSource, logger *logw.LoggerW) coco.InterruptFn {
	return func(core *coco.Core, signal os.Signal) {
		if datasource != nil {
			// close datasource
			if err := datasource.Close(); err != nil {
				core.L().Errorf("data source closed failed: %s", err)
			}
		}
		// close logger
		logger.Close()
	}
}

func LoadDataSource(dataConf *conf.DataConf, source **data.DataSource) coco.ComponentFn {
	return func(core *coco.Core) {
		dataSource, err := data.NewDataSource(core.Ctx(), dataConf)
		if err != nil {
			core.L().Panicf("data source load failed: %s", err.Error())
		}
		*source = dataSource
	}
}

func LoadLangDir(cfg *locale.Conf, lang **locale.Locale) coco.ComponentFn {
	return func(core *coco.Core) {
		l, err := locale.NewLocaleWithConf(cfg)
		if err != nil {
			core.L().Panicf("lang dir laod failed: %s", err.Error())
		}
		*lang = l
	}
}
