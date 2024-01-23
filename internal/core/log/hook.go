package log

import (
	"github.com/dstgo/filebox"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
)

type HookCloser interface {
	logrus.Hook
	io.Closer
}

type levelFileHook struct {
	path   string
	mu     sync.Mutex
	levels []logrus.Level
	writer io.WriteCloser
}

func newLevelFileHook(path string, levels ...logrus.Level) (*levelFileHook, error) {
	var writer io.WriteCloser
	file, err := filebox.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	writer = file

	return &levelFileHook{levels: levels, writer: writer}, nil
}

func (l *levelFileHook) Close() error {
	return l.writer.Close()
}

func (l *levelFileHook) Levels() []logrus.Level {
	return l.levels
}

func (l *levelFileHook) Fire(entry *logrus.Entry) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	bytes, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = l.writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
