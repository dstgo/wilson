package conf

import (
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/pkg/config"
)

// WilsonConf wilson config contains all needed configurations
type WilsonConf struct {
	ServerConf *ServerConf  `mapstructure:"app"`
	DataConf   *DataConf    `mapstructure:"data"`
	LogConf    *LogConf     `mapstructure:"log"`
	JwtConf    *JwtConf     `mapstructure:"jwt"`
	LocaleConf *locale.Conf `mapstructure:"locale"`
	EmailConf  *EmailConf   `mapstructure:"email"`
	BuildMeta  BuildInfo
}

func NewWilsonConf(config *config.Config, buildInfo BuildInfo) (*WilsonConf, error) {
	cfg := new(WilsonConf)
	if err := config.Viper().Unmarshal(cfg); err != nil {
		return nil, err
	}

	if len(buildInfo.Version) == 0 {
		buildInfo.Version = "none"
	}

	if len(buildInfo.Author) == 0 {
		buildInfo.Author = "none"
	}

	if len(buildInfo.BuildTime) == 0 {
		buildInfo.BuildTime = "none"
	}

	cfg.BuildMeta = buildInfo

	return cfg, nil
}

// WigfridConf wigfrid configuration
type WigfridConf struct {
	BuildMeta BuildInfo
	GrpcConf  *GrpcConf `mapstructure:"grpc"`
	DataConf  *DataConf `mapstructure:"data"`
	LogConf   *LogConf  `mapstructure:"log"`
	DstConf   *DstConf  `mapstructure:"dst"`
}

func NewWigfridConf(config *config.Config, buildInfo BuildInfo) (*WigfridConf, error) {
	cfg := new(WigfridConf)
	if err := config.Viper().Unmarshal(cfg); err != nil {
		return nil, err
	}

	if len(buildInfo.Version) == 0 {
		buildInfo.Version = "none"
	}

	if len(buildInfo.Author) == 0 {
		buildInfo.Author = "none"
	}

	if len(buildInfo.BuildTime) == 0 {
		buildInfo.BuildTime = "none"
	}

	cfg.BuildMeta = buildInfo

	return cfg, nil
}
