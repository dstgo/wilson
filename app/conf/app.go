package conf

import (
	"time"
)

// ServerConf app config
type ServerConf struct {
	Mode       string `mapstructure:"mode"`
	Author     string
	Repository string
	Swagger    bool     `mapstructure:"swagger"`
	Version    string   `mapstructure:"version"`
	Name       string   `mapstructure:"name"`
	Http       HttpConf `mapstructure:"http"`
	Rpc        RpcConf  `mapstructure:"rpc"`
}

// HttpConf http server config
type HttpConf struct {
	Address string `mapstructure:"address"`

	TlsConf *TlsConf `mapstructure:"tls"`

	ReadTimeout     time.Duration `mapstructure:"readTimeout"`
	WriteTimeout    time.Duration `mapstructure:"writeTimeout"`
	ReadHeadTimeout time.Duration `mapstructure:"readHeaderTimeout"`
	IdleTimeout     time.Duration `mapstructure:"idleTimeout"`
	MultipartMax    int64         `mapstructure:"multipartMax"`
	MaxHeader       int           `mapstructure:"maxHeader"`
}

// TlsConf tls config
type TlsConf struct {
	Enable bool   `mapstructure:"enable"`
	Cert   string `mapstructure:"cert"`
	Pem    string `mapstructure:"pem"`
}

// RpcConf Rpc client config
type RpcConf struct {
}

// JwtConf jwt config
type JwtConf struct {
	Sig string        `mapstructure:"sig"`
	Isu string        `mapstructure:"isu"`
	Exp time.Duration `mapstructure:"exp"`
}

// LogConf app logger config
type LogConf struct {
	Level      string `mapstructure:"level"`
	InfoLog    string `mapstructure:"infoLog"`
	ErrorLog   string `mapstructure:"errorLog"`
	TimeFormat string
	Order      []string
}
