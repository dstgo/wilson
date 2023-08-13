package logw

import (
	"errors"
	"github.com/dstgo/filebox"
	"github.com/dstgo/wilson/app/conf"
	"github.com/sirupsen/logrus"
	"os"
)

type LoggerW struct {
	conf *conf.LogConf
	l    *logrus.Logger
	hs   []*LevelLoggerH
}

func (l *LoggerW) L() *logrus.Logger {
	return l.l
}

func (l *LoggerW) Close() error {
	for _, h := range l.hs {
		if err := h.Close(); err != nil {
			return err
		}
	}
	return nil
}

func NewLogger(conf *conf.LogConf) (*LoggerW, error) {
	if conf == nil {
		return nil, errors.New("log conf is empty")
	}

	textFormatter := NewTextFormatter(conf.TimeFormat, conf.Order)

	logger := logrus.New()

	logger.SetReportCaller(true)
	logger.SetFormatter(textFormatter)

	level, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		return nil, err
	}

	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)

	flag := os.O_WRONLY | os.O_APPEND | os.O_CREATE

	infoLog, err := filebox.OpenFile(conf.InfoLog, flag, 0666)
	if err != nil {
		return nil, err
	}

	errorLog, err := filebox.OpenFile(conf.ErrorLog, flag, 0666)
	if err != nil {
		return nil, err
	}

	// file logger

	infoHook := NewLevelLoggerH(infoLog, textFormatter, 40960, logrus.DebugLevel, logrus.TraceLevel, logrus.InfoLevel)

	errorHook := NewLevelLoggerH(errorLog, textFormatter, 40960, logrus.WarnLevel, logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel)

	logger.AddHook(infoHook)
	logger.AddHook(errorHook)

	return &LoggerW{
		conf: conf,
		l:    logger,
		hs:   []*LevelLoggerH{infoHook, errorHook},
	}, nil
}
