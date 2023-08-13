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

func NewAppConf(config *config.Config) (*AppConf, error) {
	cfg := new(AppConf)
	if err := config.Viper().Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
