package email

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
)

var (
	RuleEmail = helper.Rules(is.Email)

	RuleEmailCode = helper.Rules(is.Alphanumeric, vax.EqLength(8, false))
)
