package locale

import (
	"fmt"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	LocaleNotFoundErr     = errors.New("locale not found")
	LocaleFileLeastOneErr = errors.New("expected least one toml language file")
)

type Group = map[string]*viper.Viper

type Locale struct {
	locale string
	mv     Group
}

func (l *Locale) Default() string {
	return l.locale
}

func (l *Locale) GetWithLocale(locale string, key string, args ...any) string {
	v, e := l.mv[locale]
	if !e {
		return LocaleNotFoundErr.Error()
	}
	format := v.GetString(key)
	return fmt.Sprintf(format, args...)
}

func (l *Locale) GetWithCtx(ctx *gin.Context, key string, args ...any) string {
	lang := l.locale
	language := httpx.GetAcceptLanguage(ctx)
	if len(language) != 0 {
		lang = language[0]
	}
	return l.GetWithLocale(lang, key, args...)
}

func (l *Locale) Get(key string, args ...any) string {
	return l.GetWithLocale(l.locale, key, args...)
}

func NewLocale(locale string, mv Group) *Locale {
	return &Locale{
		locale: locale,
		mv:     mv,
	}
}

func NewLocaleWithConf(conf *Conf) (*Locale, error) {
	group, err := conf.Load()
	if err != nil {
		return nil, err
	}
	return &Locale{conf.Locale, group}, nil
}
