package server

import (
	"context"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/log"
	"github.com/go-kratos/kratos/contrib/log/logrus/v2"
	"github.com/go-kratos/kratos/v2"
)

// NewGrpcApp initializes grpc app server
func NewGrpcApp(ctx context.Context, cfg *conf.WigfridConf, logger *log.Logger) (*kratos.App, error) {
	// set the global logger
	log.Setup(logger.L())
	// log banner
	if err := LogBanner(cfg.BuildMeta, cfg.LogConf, logger.L(), "wigfrid.txt"); err != nil {
		return nil, err
	}
	// load datasource
	datasource, err := LoadDataSource(ctx, cfg.DataConf)
	if err != nil {
		return nil, err
	}
	// use the kratos logrus logger adaptor
	kratosLogger := logrus.NewLogger(logger.L())

	// create the grpc server
	grpcServer := NewGRPCServer(cfg.GrpcConf, kratosLogger)
	// app cleanup func
	cleanup := func(ctx context.Context) error {
		if err := datasource.Close(); err != nil {
			return err
		}
		if err := logger.Close(); err != nil {
			return err
		}
		return nil
	}

	// create kratos app
	app := kratos.New(
		kratos.Context(ctx),
		kratos.Logger(kratosLogger),
		kratos.Server(grpcServer),
		kratos.AfterStop(cleanup),
	)

	return app, nil
}
