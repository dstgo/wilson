package config

import (
	kc "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/dstgo/wilson/framework/kratosx/config"
)

type Config struct {
	conf        config.Config
	Debug       bool
	Addr        string
	Discovery   string
	Endpoints   []Endpoint
	Middlewares []Middleware
}

type Watch func(*Config)

// New 新建并初始化配置
func New(source kc.Source) (*Config, error) {
	ins := config.New(source)
	if err := ins.Load(); err != nil {
		return nil, err
	}

	conf := &Config{
		conf: ins,
	}
	return conf, ins.Scan(conf)
}

// WatchEndpoints 监听配置
func (c *Config) WatchEndpoints(fn Watch) {
	c.conf.Watch("endpoints", func(value config.Value) {
		var ends []Endpoint
		if err := value.Scan(&ends); err != nil {
			log.Error("watch endpoints change error:" + err.Error())
			return
		}
		c.Endpoints = ends
		fn(c)
	})
}
