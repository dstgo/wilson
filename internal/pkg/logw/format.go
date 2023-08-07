package logw

import (
	"github.com/sirupsen/logrus"
)

func NewJsonFormatter() *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999 -07:00",
	}
}

func NewTextFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999 -07:00",
	}
}
