package log

import (
	"errors"
	"github.com/dstgo/wilson/app/conf"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	logger *logrus.Logger
)

func L() *logrus.Logger {
	if logger == nil {
		panic("logger is not initialized")
	}
	return logger
}

func Setup(l *logrus.Logger) {
	logger = l
}

type Logger struct {
	conf *conf.LogConf
	l    *logrus.Logger
	hs   []HookCloser
}

func (l *Logger) L() *logrus.Logger {
	return l.l
}

func (l *Logger) Close() error {
	for _, h := range l.hs {
		if h == nil {
			continue
		}

		if err := h.Close(); err != nil {
			return err
		}
	}
	return nil
}

func NewLogger(logConf *conf.LogConf) (*Logger, error) {
	if logConf == nil {
		return nil, errors.New("log logConf is empty")
	}

	parseLevel, err := logrus.ParseLevel(logConf.Level)
	if err != nil {
		return nil, err
	}

	var (
		logger    = logrus.New()
		formatter logrus.Formatter
		infoHook  HookCloser
		errorHook HookCloser
	)

	formatter, err = newFormatter(logConf)
	if err != nil {
		return nil, err
	}

	if parseLevel >= logrus.DebugLevel {
		logger.SetReportCaller(true)
	}

	logger.SetLevel(parseLevel)
	logger.SetFormatter(formatter)
	logger.SetOutput(os.Stdout)

	if len(logConf.InfoLog) > 0 {
		infoHook, err = newLevelFileHook(logConf.InfoLog, logrus.WarnLevel, logrus.InfoLevel, logrus.DebugLevel, logrus.TraceLevel)
		if err != nil {
			return nil, err
		}
		logger.AddHook(infoHook)
	}

	if len(logConf.ErrorLog) > 0 {
		errorHook, err = newLevelFileHook(logConf.ErrorLog, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel)
		if err != nil {
			return nil, err
		}
		logger.AddHook(errorHook)
	}

	return &Logger{
		conf: logConf,
		l:    logger,
		hs:   []HookCloser{infoHook, errorHook},
	}, nil
}
