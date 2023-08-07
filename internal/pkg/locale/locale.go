package locale

import (
	"fmt"
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

func (l *Locale) GetWithLocale(locale string, key string, args ...any) (string, error) {
	v, e := l.mv[locale]
	if !e {
		return "", errors.Wrap(LocaleNotFoundErr, locale)
	}
	format := v.GetString(key)
	return fmt.Sprintf(format, args...), nil
}

func (l *Locale) Get(key string, args ...any) string {
	locale, err := l.GetWithLocale(l.locale, key, args)
	if err != nil {
		return err.Error()
	}
	return locale
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
