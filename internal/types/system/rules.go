package system

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/pkg/vax"
)

var (
	RulePing = helper.Rules(vax.RangeLenRune(1, 10), vax.In("wilson", "wendy"))
)
