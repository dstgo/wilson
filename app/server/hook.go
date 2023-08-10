package server

import (
	"github.com/dstgo/size"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/pkg/logw"
	"github.com/dstgo/wilson/app/pkg/sysinfo"
	"github.com/dstgo/wilson/pkg/coco"
	"github.com/dstgo/wilson/pkg/coco/route"
	"os"
	"strings"
)

// on server boot hooks

func LoadDataSource(dataConf *conf.DataConf, components *Components) coco.ComponentFn {
	return func(core *coco.Core) {
		datasource, err := data.NewDataSource(core.Ctx(), dataConf, core)
		if err != nil {
			core.L().Panicf("load data datasource failed: %s", err.Error())
		}
		core.L().Infof("load data datasource ok √")
		components.Datasource = datasource
	}
}

// after server booted hooks

func LogBanner(cfg *conf.AppConf) coco.ComponentFn {
	return func(core *coco.Core) {
		banner := `

             ██  ██                                                                           
           ░░  ░██                                                                           
 ███     ██ ██ ░██  ██████  ██████  ███████     ██████  █████  ██████ ██    ██  █████  ██████
░░██  █ ░██░██ ░██ ██░░░░  ██░░░░██░░██░░░██   ██░░░░  ██░░░██░░██░░█░██   ░██ ██░░░██░░██░░█
 ░██ ███░██░██ ░██░░█████ ░██   ░██ ░██  ░██  ░░█████ ░███████ ░██ ░ ░░██ ░██ ░███████ ░██ ░ 
 ░████░████░██ ░██ ░░░░░██░██   ░██ ░██  ░██   ░░░░░██░██░░░░  ░██    ░░████  ░██░░░░  ░██   
 ███░ ░░░██░██ ███ ██████ ░░██████  ███  ░██   ██████ ░░██████░███     ░░██   ░░██████░███   
░░░    ░░░ ░░ ░░░ ░░░░░░   ░░░░░░  ░░░   ░░   ░░░░░░   ░░░░░░ ░░░       ░░     ░░░░░░ ░░░

Author=%s AppName=%s Mode=%s Version=%s
Os=%s %s %s Arch=%s
Cpu=%s %d Cores Mem=%s`
		hostInfo := sysinfo.GetHostInfo()
		cpuInfo := sysinfo.GetCpuInfo()
		memInfo := sysinfo.GetMemInfo()

		appConf := cfg.AppConf
		core.L().Infof(banner,
			appConf.Author, appConf.Name, strings.ToUpper(appConf.Mode), appConf.Version,
			hostInfo.Os, hostInfo.Platform, hostInfo.Version, hostInfo.Arch,
			cpuInfo.Name, cpuInfo.Count, size.ParseTargetSize(memInfo.Virtual.Total.String(), size.GB),
		)
	}
}

func BootLog(cfg *conf.AppConf) coco.ComponentFn {
	return func(core *coco.Core) {
		appConf := cfg.AppConf
		core.L().Infof("wilson app boot successfully, http server is listenning at %s, tls enable %t", core.Server().Addr, appConf.Http.TlsConf.Enable)
	}
}

func DebugPrintRouter() coco.ComponentFn {
	return func(core *coco.Core) {
		core.RootRouter().Walk(func(info route.RouterInfo) error {
			if !info.IsGroup {
				core.L().
					WithField("method", info.Method).
					WithField("path", info.FullPath).
					Debugln()
			}
			return nil
		})
	}
}

// on server shutdown hooks

func CloseDataSource(datasource *data.DataSource) coco.InterruptFn {
	return func(core *coco.Core, signal os.Signal) {
		if datasource != nil {
			// close datasource
			if err := datasource.Close(); err != nil {
				core.L().Errorf("data source closed failed: %s", err)
				return
			}
		}
		core.L().Infoln("data source closed successfully")
	}
}

func CloseLogger(w *logw.LoggerW) coco.InterruptFn {
	return func(core *coco.Core, signal os.Signal) {
		// close logger
		w.Close()
	}
}

func ShutdownWithInfo() coco.InterruptFn {
	return func(core *coco.Core, signal os.Signal) {
		core.L().Infof("received os signal: %s, ready to graceful shutdown", signal.String())
	}
}

func ShutdownWithCloseHttp() coco.InterruptFn {
	return func(core *coco.Core, signal os.Signal) {
		core.Server().Shutdown(core.Ctx())
	}
}
