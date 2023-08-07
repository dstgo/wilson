package conf

import (
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/pkg/coco"
)

// AppConf wilson config contains all needed configurations
type AppConf struct {
	App    *ServerConf  `mapstructure:"app"`
	Data   *DataConf    `mapstructure:"data"`
	Log    *LogConf     `mapstructure:"log"`
	Jwt    *JwtConf     `mapstructure:"jwt"`
	Locale *locale.Conf `mapstructure:"locale"`
}

func NewAppConf(config *coco.Config) (*AppConf, error) {
	cfg := new(AppConf)
	if err := config.Viper().Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
