package data

import (
	"fmt"
	"github.com/dstgo/wilson/app/conf"
	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"net/netip"
)

type Dsn func(conf *conf.DatabaseConf) string

// MysqlDsn
// eg: user:password@tcp(ip:port)/db?params
func MysqlDsn(cfg *conf.DatabaseConf) string {
	return fmt.Sprintf(
		`%s:%s@%s(%s)/%s`,
		cfg.User, cfg.Password, cfg.Network, cfg.Address, cfg.Params)
}

// SqlLiteDsn
// eg: file:test.db?cache=shared&mode=memory
func SqlLiteDsn(cfg *conf.DatabaseConf) string {
	return fmt.Sprintf("%s", cfg.Params)
}

// SqlServerDsn
// eg: sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm
func SqlServerDsn(cfg *conf.DatabaseConf) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s?%s",
		cfg.User, cfg.Password, cfg.Address, cfg.Params)
}

// PostGreDsn
// eg: host=localhost port=9920 user=gorm password=gorm dbname=gorm sslmode=disable TimeZone=Asia/Shanghai
func PostGreDsn(cfg *conf.DatabaseConf) string {
	addrPort := netip.MustParseAddrPort(cfg.Address)
	return fmt.Sprintf("host=%s port=%d user=%s password=%s %s",
		addrPort.Addr(), addrPort.Port(), cfg.User, cfg.Password, cfg.Params)
}

// ClickHouseDsn
// eg: tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20
func ClickHouseDsn(cfg *conf.DatabaseConf) string {
	return fmt.Sprintf("%s://%s?username=%s&password=%s&%s",
		cfg.Network, cfg.Address, cfg.User, cfg.Password, cfg.Params)
}

type DBDriver func(cfg *conf.DatabaseConf) (gorm.Dialector, error)

var SupportDrivers = map[string]DBDriver{
	"mysql":      MysqlDriver,
	"postgres":   PostGreDriver,
	"sqlite":     SqlLiteDriver,
	"sqlserver":  SqlServerDriver,
	"clickhouse": ClickHouseDriver,
	"tidb":       TiDBDriver,
}

func MysqlDriver(cfg *conf.DatabaseConf) (gorm.Dialector, error) {
	dsn := MysqlDsn(cfg)
	dsnConf, err := mysqldriver.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	client := mysql.New(mysql.Config{
		DriverName: cfg.Driver,
		DSN:        dsn,
		DSNConfig:  dsnConf,
	})
	return client, nil
}

func SqlLiteDriver(cfg *conf.DatabaseConf) (gorm.Dialector, error) {
	return sqlite.Open(SqlLiteDsn(cfg)), nil
}

func PostGreDriver(cfg *conf.DatabaseConf) (gorm.Dialector, error) {
	dsn := PostGreDsn(cfg)
	return postgres.Open(dsn), nil
}

func SqlServerDriver(cfg *conf.DatabaseConf) (gorm.Dialector, error) {
	dsn := SqlServerDsn(cfg)
	return sqlserver.Open(dsn), nil
}

// TiDBDriver
// TiDB is compatible with Mysql protocol
func TiDBDriver(cfg *conf.DatabaseConf) (gorm.Dialector, error) {
	return MysqlDriver(cfg)
}

func ClickHouseDriver(cfg *conf.DatabaseConf) (gorm.Dialector, error) {
	return clickhouse.Open(ClickHouseDsn(cfg)), nil
}
