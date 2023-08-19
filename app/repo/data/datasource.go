package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/dstgo/task"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/pkg/errorx"
	"github.com/dstgo/wilson/app/repo/data/entity"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
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
	var (
		dial gorm.Dialector
		err  error
	)

	driverFn, exist := SupportDrivers[conf.Driver]
	if !exist {
		return nil, fmt.Errorf("driverFn is unsupport: %s", conf.Driver)
	}

	dial, err = driverFn(conf)
	if err != nil {
		return nil, err
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
	db.SetConnMaxIdleTime(conf.MaxIdleTime)
	db.SetConnMaxLifetime(conf.MaxLifetime)

	// ping server
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	// migrate table
	if err := entity.Migrate(ormDB); err != nil {
		return ormDB, err
	}

	return ormDB, nil
}

type DataSource struct {
	redis *redis.Client
	orm   *gorm.DB
}

func (d *DataSource) Redis() *redis.Client {
	return d.redis
}

func (d *DataSource) ORM() *gorm.DB {
	return d.orm
}

func (d *DataSource) Close() error {
	db, err := d.orm.DB()
	return errorx.Join(
		err,
		db.Close(),
		d.redis.Close(),
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
		datasource.orm = db
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
		datasource.redis = redisClient
	})

	dataTask.Run()

	return datasource, errs
}
