package logw

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

type LevelLoggerH struct {
	closer    io.Closer
	out       *bufio.Writer
	mu        sync.Mutex
	formatter logrus.Formatter
	levels    []logrus.Level
}

func NewLevelLoggerH(writer io.WriteCloser, formatter logrus.Formatter, buf int, levels ...logrus.Level) *LevelLoggerH {
	return &LevelLoggerH{
		out:       bufio.NewWriterSize(writer, buf),
		formatter: formatter,
		closer:    writer,
		mu:        sync.Mutex{},
		levels:    levels,
	}
}

func (i *LevelLoggerH) Levels() []logrus.Level {
	return i.levels
}

func (i *LevelLoggerH) Fire(entry *logrus.Entry) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	bytes, err := i.formatter.Format(entry)
	if err != nil {
		return err
	}

	if _, err := i.out.Write(bytes); err != nil {
		return err
	}
	i.out.Flush()
	return nil
}

func (i *LevelLoggerH) Close() error {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.out.Flush()
	return i.closer.Close()
}
