package data

import (
	"context"
	"errors"
	"github.com/dstgo/task"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/pkg/errorx"
	"github.com/go-redis/redis/v8"
	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
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

func NewDBClient(ctx context.Context, conf *conf.DatabaseConf, logger *logrus.Logger) (*gorm.DB, error) {
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

	ormDB, err := gorm.Open(dial, ormOption(ormLogger(logger)))
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
	dsn := MysqlDsn(conf)
	dsnConf, err := mysqldriver.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	client := mysql.New(mysql.Config{
		DriverName: conf.Driver,
		DSN:        dsn,
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
	return errorx.Join(
		err,
		db.Close(),
		d.Redis.Close(),
	).Err()
}

func NewDataSource(ctx context.Context, databaseConf *conf.DataConf, logger *logrus.Logger) (*DataSource, error) {
	datasource := new(DataSource)

	var errs error
	// async task to load each db
	dataTask := task.NewTask(func(err any) {
		logger.Panicln(err)
	})

	// connect to gorm db
	dataTask.AddJobs(func() {
		db, err := NewDBClient(ctx, databaseConf.DatabaseConf, logger)
		if err != nil {
			errs = errorx.Join(errs, err).Err()
			logger.Errorf("gorm db connected failed: %s", err)
			return
		}
		logger.Infof("gorm db connected:(%s) ok √", databaseConf.DatabaseConf.Address)
		datasource.OrmDB = db
	})

	// connect to redis db
	dataTask.AddJobs(func() {
		redisClient, err := NewRedisClient(ctx, databaseConf.RedisConf)
		if err != nil {
			errs = errorx.Join(errs, err).Err()
			logger.Errorf("redis client connected failed: %s", err)
			return
		}
		logger.Infof("redis server connected:(%s) ok √", redisClient.Options().Addr)
		datasource.Redis = redisClient
	})

	dataTask.Run()

	return datasource, errs
}
