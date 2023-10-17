package locale

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	LocaleUnSupportedErr  = errors.New("locale is unsupported")
	LocaleFileLeastOneErr = errors.New("expected least one language file")
)

var (
	locale *Locale
)

func L() *Locale {
	if locale == nil {
		panic("locale is not initialized")
	}
	return locale
}

func Setup(l *Locale) {
	locale = l
}

func Get(key string, args ...any) string {
	if locale != nil {
		return locale.GetDefault(key, args...)
	}
	return ""
}

func GetWithLang(lang string, key string, args ...any) string {
	if locale != nil {
		return locale.Get(lang, key, args...)
	}
	return ""
}

type Group = map[string]*viper.Viper

type Locale struct {
	locale string
	mv     Group
}

func (l *Locale) Default() string {
	return l.locale
}

func (l *Locale) Get(lang string, key string, args ...any) string {
	if len(lang) == 0 {
		lang = l.locale
	}
	v, e := l.mv[lang]
	if !e {
		return errors.Wrap(LocaleUnSupportedErr, lang).Error()
	}
	return fmt.Sprintf(v.GetString(key), args...)
}

func (l *Locale) GetDefault(key string, args ...any) string {
	return l.Get(l.locale, key, args...)
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
