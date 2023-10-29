package helper

import (
	"github.com/dstgo/wilson/pkg/vax"
	"regexp"
)

func Rules(rule ...vax.Rule) []vax.Rule {
	return rule
}

func RequiredRules(rule []vax.Rule) []vax.Rule {
	return append([]vax.Rule{vax.Required}, rule...)
}

type RegexRule struct {
	Regx *regexp.Regexp
	I18n string
}
