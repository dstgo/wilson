package rules

import "github.com/dstgo/wilson/pkg/vax"

func Rules(rule ...vax.Rule) []vax.Rule {
	return rule
}

func Required(rule []vax.Rule) []vax.Rule {
	return append([]vax.Rule{vax.Required}, rule...)
}
