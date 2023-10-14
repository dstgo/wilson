package conf

import (
	"fmt"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"net/smtp"
	"time"
)

// ServerConf app config
type ServerConf struct {
	Mode      string `mapstructure:"mode"`
	Author    string
	GoVersion string
	Swagger   bool     `mapstructure:"swagger"`
	OpenAPI   bool     `mapstructure:"openapi"`
	Version   string   `mapstructure:"version"`
	Name      string   `mapstructure:"name"`
	HttpConf  HttpConf `mapstructure:"http"`
	Rpc       RpcConf  `mapstructure:"rpc"`
}

// HttpConf http server config
type HttpConf struct {
	Address string `mapstructure:"address"`

	TlsConf  *TlsConf    `mapstructure:"tls"`
	CorsConf *httpx.Cors `mapstructure:"cors"`

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
	Format     string `mapstructure:"format"`
	Level      string `mapstructure:"level"`
	InfoLog    string `mapstructure:"infoLog"`
	ErrorLog   string `mapstructure:"errorLog"`
	TimeFormat string
	Order      []string
}

type EmailConf struct {
	Host        string        `mapstructure:"host"`
	Port        int           `mapstructure:"port"`
	User        string        `mapstructure:"user"`
	Password    string        `mapstructure:"password"`
	SendTimeout time.Duration `mapstructure:"timeout"`
	MaxPoolSize int           `mapstructure:"maxPoolSize"`
	Exp         int           `mapstructure:"exp"`
}

func (e EmailConf) Expire() time.Duration {
	return time.Duration(e.Exp) * time.Minute
}

func (e EmailConf) Address() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

func (e EmailConf) SmtpAuth() smtp.Auth {
	return smtp.PlainAuth("", e.User, e.Password, e.Host)
}
