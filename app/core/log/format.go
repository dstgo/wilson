package log

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

func NewJsonFormatter(layout string) logrus.Formatter {
	return &logrus.JSONFormatter{
		TimestampFormat: layout,
	}
}

func NewTextFormatter(layout string, order []string) logrus.Formatter {
	return &nested.Formatter{
		NoColors:      true,
		ShowFullLevel: true,
		FieldsOrder:   order,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			return fmt.Sprintf("\t(%s %s:%d)", frame.Func.Name(), path.Base(frame.File), frame.Line)
		},
		TimestampFormat: layout,
	}
}
