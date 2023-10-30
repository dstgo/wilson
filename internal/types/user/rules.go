package user

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
)

var (
	RuleUsername = helper.Rules(is.Alphanumeric, vax.RangeLenRune(6, 20))

	RulePassword = helper.Rules(vax.RangeLenRune(10, 30))
)
