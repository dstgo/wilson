package config

import (
	"bytes"
	"github.com/spf13/viper"
	"io"
)

func NewConfigFile(file string) *Config {
	v := viper.New()
	v.SetConfigFile(file)
	return &Config{
		reader: nil,
		v:      v,
	}
}
func NewConfigReader(reader io.Reader, cfgType string) *Config {
	v := viper.New()
	v.SetConfigType(cfgType)
	return &Config{
		reader: reader,
		v:      v,
	}
}

func DefaultConfig() *Config {
	return &Config{
		reader: bytes.NewReader([]byte{}),
		v:      viper.GetViper(),
	}
}

type Config struct {
	reader io.Reader
	v      *viper.Viper
}

func (c *Config) SetCfgType(t string) {
	c.v.SetConfigType(t)
}

func (c *Config) Viper() *viper.Viper {
	return c.v
}

func (c *Config) ReadConfig() error {
	if len(c.v.ConfigFileUsed()) > 0 {
		return c.v.ReadInConfig()
	} else if c.reader != nil {
		return c.v.ReadConfig(c.reader)
	}
	return nil
}
