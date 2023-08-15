package conf

import (
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/pkg/config"
)

// AppConf wilson config contains all needed configurations
type AppConf struct {
	ServerConf *ServerConf  `mapstructure:"app"`
	DataConf   *DataConf    `mapstructure:"data"`
	LogConf    *LogConf     `mapstructure:"log"`
	JwtConf    *JwtConf     `mapstructure:"jwt"`
	LocaleConf *locale.Conf `mapstructure:"locale"`
}

func NewAppConf(config *config.Config, author string, version string, goVersion string) (*AppConf, error) {
	cfg := new(AppConf)
	if err := config.Viper().Unmarshal(cfg); err != nil {
		return nil, err
	}
	cfg.ServerConf.Author = author
	cfg.ServerConf.Version = version
	cfg.ServerConf.GoVersion = goVersion
	return cfg, nil
}
