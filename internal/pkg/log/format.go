package log

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

func newFormatter(logConf *conf.LogConf) (logrus.Formatter, error) {
	if logConf.Format == "text" {
		return newTextFormatter(logConf.TimeFormat, logConf.Order), nil
	} else if logConf.Format == "json" {
		return newJsonFormatter(logConf.TimeFormat), nil
	} else {
		return nil, fmt.Errorf("unsupported log format: %s", logConf.Format)
	}
}

func newJsonFormatter(layout string) logrus.Formatter {
	return &logrus.JSONFormatter{
		TimestampFormat: layout,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return path.Base(frame.Func.Name()), fmt.Sprintf("%s:%d", frame.File, frame.Line)
		},
	}
}

func newTextFormatter(layout string, order []string) logrus.Formatter {
	return &nested.Formatter{
		NoColors:        true,
		ShowFullLevel:   true,
		FieldsOrder:     order,
		TimestampFormat: layout,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			return fmt.Sprintf("\t(%s:%d\t%s)", frame.File, frame.Line, path.Base(frame.Func.Name()))
		},
	}
}
