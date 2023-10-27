package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/dstgo/task"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/log"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/utils"
	"github.com/dstgo/wilson/internal/types/errs"
	"github.com/go-redis/redis/v8"
	errors2 "github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
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

// NewDBClient return a new db client
// due to gorm not support open with cancel context, so we need to implement it by self
func NewDBClient(ctx context.Context, conf *conf.DatabaseConf) (*gorm.DB, error) {

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

	ormDB, err := gorm.Open(dial, &gorm.Config{
		Logger:          ormLogger(log.L()),
		PrepareStmt:     true,
		CreateBatchSize: 50,
	})
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
	if utils.IsDebugMode() {
		return d.orm.Debug()
	}
	return d.orm
}

func (d *DataSource) Close() error {
	db, err := d.orm.DB()
	return errs.Join(
		err,
		db.Close(),
		d.redis.Close(),
	).Err()
}

func NewDataSource(ctx context.Context, databaseConf *conf.DataConf) (*DataSource, error) {
	datasource := new(DataSource)

	// async task to load each db
	dataTask, cancel := task.WithTimeout(ctx, time.Second*10)
	defer cancel(nil)

	// connect to gorm db
	gormDataWorker := task.NewWorker(func(ctx context.Context) error {
		db, err := NewDBClient(ctx, databaseConf.DatabaseConf)
		if err != nil {
			return errors2.Wrap(err, "gorm db connected failed")
		}
		log.L().Infof("gorm db connected:(%s) ok √", databaseConf.DatabaseConf.Address)
		datasource.orm = db
		return nil
	})

	// connect to redis db
	redisDataWorker := task.NewWorker(func(ctx context.Context) error {
		redisClient, err := NewRedisClient(ctx, databaseConf.RedisConf)
		if err != nil {
			return errors2.Wrap(err, "redis client connected failed")
		}
		log.L().Infof("redis server connected:(%s) ok √", redisClient.Options().Addr)
		datasource.redis = redisClient
		return nil
	})

	dataTask.Add(gormDataWorker, redisDataWorker)

	err := dataTask.Run()
	if err != nil {
		return datasource, err
	}

	return datasource, nil
}
