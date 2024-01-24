package conf

import (
	"fmt"
	"github.com/dstgo/wilson/pkg/ginx/httpx"
	"net/smtp"
	"time"
)

type BuildInfo struct {
	Author    string
	Version   string
	BuildTime string
}

// ServerConf app config
type ServerConf struct {
	Swagger  bool     `mapstructure:"swagger"`
	OpenAPI  bool     `mapstructure:"openapi"`
	HttpConf HttpConf `mapstructure:"http"`
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

type GrpcConf struct {
	Address       string `mapstructure:"address"`
	MaxRecv       int    `mapstructure:"maxRecv"`
	MaxSend       int    `mapstructure:"maxSend"`
	ReadBuffer    int    `mapstructure:"readBuffer"`
	WriteBuffer   int    `mapstructure:"writeBuffer"`
	MaxHeaderSize uint32 `mapstructure:"maxHeaderSize"`
}

// TlsConf tls config
type TlsConf struct {
	Enable bool   `mapstructure:"enable"`
	Cert   string `mapstructure:"cert"`
	Pem    string `mapstructure:"pem"`
}

// JwtConf jwt config
type JwtConf struct {
	Sig   string        `mapstructure:"sig"`
	Isu   string        `mapstructure:"isu"`
	Exp   time.Duration `mapstructure:"exp"`
	RExp  time.Duration `mapstructure:"rexp"`
	Delay time.Duration `mapstructure:"delay"`
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
