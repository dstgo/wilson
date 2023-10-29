package role

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
	"regexp"
)

var (
	PermObjectRule = helper.RegexRule{Regx: regexp.MustCompilePOSIX("^[a-zA-Z0-9/_-]*$"), I18n: "rule.perm.object"}
)

var (
	RulePermObj   = helper.Rules(vax.Match(PermObjectRule.Regx).Code(PermObjectRule.I18n), vax.RangeLength(1, 200, false))
	RulePermAct   = helper.Rules(vax.RangeLength(1, 60, false), is.Alpha)
	RulePermTag   = helper.Rules(vax.RangeLength(1, 20, false), is.Alphanumeric)
	RulePermGroup = helper.Rules(vax.RangeLength(1, 60, false))
	RuleRoleName  = helper.Rules(vax.RangeLength(1, 60, false))
	RuleRoleCode  = helper.Rules(vax.RangeLength(1, 30, false), is.Alphanumeric)
)
