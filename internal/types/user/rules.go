package user

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
)

var (
	RuleUsername = helper.Rules(is.Alphanumeric, vax.RangeLength(6, 20, false))

	RulePassword = helper.Rules(vax.RangeLength(10, 30, false))
)
