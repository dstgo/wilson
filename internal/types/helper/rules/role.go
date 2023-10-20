package rules

import (
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
	"regexp"
)

var (
	PermObjectRule = RegexRule{regexp.MustCompilePOSIX("^[a-zA-Z0-9/_-]*$"), "rule.perm.object"}
	PermObject     = Rules(vax.Match(PermObjectRule.Regx).Code(PermObjectRule.I18n), vax.RangeLength(1, 200, false))
	PermAction     = Rules(vax.RangeLength(1, 60, false), is.Alpha)
	PermTag        = Rules(vax.RangeLength(1, 20, false), is.Alphanumeric)
	RoleName       = Rules(vax.RangeLength(1, 60, false))
	RoleCode       = Rules(vax.RangeLength(1, 30, false), is.Alphanumeric)
)
