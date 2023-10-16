package errs

import (
	"github.com/dstgo/wilson/internal/sys/locale"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func WrapI18nCtx(ctx *gin.Context, err error, key string, args ...any) error {
	return errors.Wrap(err, locale.GetWithCtx(ctx, key, args...))
}

func WrapI18n(err error, key string, args ...any) error {
	return errors.Wrap(err, locale.Get(key, args...))
}
