package errorx

import (
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func WrapI18n(ctx *gin.Context, err error, key string, args ...any) error {
	return errors.Wrap(err, locale.L().GetWithCtx(ctx, key, args...))
}
