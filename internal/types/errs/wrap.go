package errs

import (
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/pkg/errors"
)

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func WrapI18n(err error, key string, args ...any) error {
	return errors.Wrap(err, locale.Get(key, args...))
}
