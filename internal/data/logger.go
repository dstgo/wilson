package data

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func ormOption(l logger.Interface) gorm.Option {
	return GormOptions{l: l}
}

type GormOptions struct {
	l logger.Interface
}

func (l GormOptions) Apply(config *gorm.Config) error {
	config.Logger = l.l
	return nil
}

func (l GormOptions) AfterInitialize(db *gorm.DB) error {
	return nil
}

func ormLogger(logger *logrus.Logger) logger.Interface {
	return &GormLogger{
		l: logger,
	}
}

// GormLogger
// an adaptor between gorm and logrus
type GormLogger struct {
	level logger.LogLevel
	l     *logrus.Logger
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	// do nothing
	return g
}

func (g *GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	g.l.WithContext(ctx).Infof(s, i...)
}

func (g *GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	g.l.WithContext(ctx).Warnf(s, i...)
}

func (g *GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	g.l.WithContext(ctx).Errorf(s, i...)
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, affected := fc()
	g.l.WithContext(ctx).
		WithField("begin", begin).
		WithField("sql", sql).
		WithField("rowAffected", affected).
		WithError(err).
		Traceln("gorm trace")
}
