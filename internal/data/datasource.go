package data

import (
	"context"
	"errors"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/pkg/errorw"
	"github.com/go-redis/redis/v8"
	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewRedisClient(ctx context.Context, conf *conf.RedisConf) (*redis.Client, error) {
	if conf == nil {
		return nil, errors.New("redis conf is empty")
	}
	client := redis.NewClient(conf.Options())

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}

func NewDBClient(ctx context.Context, conf *conf.DatabaseConf) (*gorm.DB, error) {
	if conf == nil {
		return nil, errors.New("db conf is empty")
	}
	var dial gorm.Dialector

	// choose driver
	switch conf.Driver {
	case "mysql":
		client, err := newMysqlClient(conf)
		if err != nil {
			return nil, err
		}
		dial = client
	}

	ormDB, err := gorm.Open(dial)
	if err != nil {
		return nil, err
	}

	db, err := ormDB.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conf.MaxOpenCons)
	db.SetMaxIdleConns(conf.MaxIdleCons)

	// ping server
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return ormDB, nil
}

func newMysqlClient(conf *conf.DatabaseConf) (gorm.Dialector, error) {
	dsnConf, err := mysqldriver.ParseDSN(conf.Dsn)
	if err != nil {
		return nil, err
	}
	client := mysql.New(mysql.Config{
		DriverName: conf.Driver,
		DSN:        conf.Dsn,
		DSNConfig:  dsnConf,
	})
	return client, nil
}

var DataSourceSet = wire.NewSet(NewDataSource)

type DataSource struct {
	Redis *redis.Client
	OrmDB *gorm.DB
}

func (d *DataSource) Close() error {
	db, err := d.OrmDB.DB()
	return errorw.Join(
		err,
		db.Close(),
		d.Redis.Close(),
	).Err()
}

func NewDataSource(ctx context.Context, databaseConf *conf.DataConf) (*DataSource, error) {
	db, err := NewDBClient(ctx, databaseConf.DatabaseConf)
	if err != nil {
		return nil, err
	}
	redisClient, err := NewRedisClient(ctx, databaseConf.RedisConf)
	if err != nil {
		return nil, err
	}
	return &DataSource{
		Redis: redisClient,
		OrmDB: db,
	}, nil
}
