package data

import (
	"fmt"
	"github.com/dstgo/wilson/app/conf"
)

type Dsn func(conf *conf.DatabaseConf) string

// MysqlDsn eg: user:password@tcp(ip:port)/db?params
// param conf *conf.DatabaseConf
// return string
func MysqlDsn(cfg *conf.DatabaseConf) string {
	return fmt.Sprintf(
		`%s:%s@%s(%s)/%s`,
		cfg.User, cfg.Password, cfg.Network, cfg.Address, cfg.Params)
}
