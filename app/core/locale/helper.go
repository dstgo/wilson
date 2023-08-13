package locale

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func (l *Locale) NewError(key string, args ...any) error {
	return errors.New(l.Get(key, args...))
}

func (l *Locale) NewErrorWithCtx(ctx *gin.Context, key string, args ...any) error {
	return errors.New(l.GetWithCtx(ctx, key, args...))
}

func (l *Locale) NewErrorWithLocale(locale string, key string, args ...any) error {
	return errors.New(l.GetWithLocale(locale, key, args...))
}
