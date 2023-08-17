package conf

import (
	"github.com/go-redis/redis/v8"
	"time"
)

// DataConf Here are the data related configurations
type DataConf struct {
	DatabaseConf *DatabaseConf `mapstructure:"database"`
	RedisConf    *RedisConf    `mapstructure:"redis"`
}

// RedisConf redis client configuration
type RedisConf struct {
	Address      string        `mapstructure:"addr"`
	Auth         string        `mapstructure:"auth"`
	Retry        int           `mapstructure:"retry"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
}

func (r RedisConf) Options() *redis.Options {
	return &redis.Options{
		Addr:         r.Address,
		Password:     r.Auth,
		MaxRetries:   r.Retry,
		ReadTimeout:  r.ReadTimeout,
		WriteTimeout: r.WriteTimeout,
	}
}

// DatabaseConf relational database config
type DatabaseConf struct {
	Driver      string        `mapstructure:"driver"`
	Network     string        `mapstructure:"network"`
	Address     string        `mapstructure:"addr"`
	User        string        `mapstructure:"user"`
	Password    string        `mapstructure:"pswd"`
	Params      string        `mapstructure:"params"`
	MaxOpenCons int           `mapstructure:"maxOpenCons"`
	MaxIdleCons int           `mapstructure:"maxIdleCons"`
	MaxIdleTime time.Duration `mapstructure:"maxIdleTime"`
	MaxLifetime time.Duration `mapstructure:"maxLifetime"`
}
