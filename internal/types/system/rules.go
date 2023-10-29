package system

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/pkg/vax"
)

var (
	RulePing = helper.Rules(vax.RangeLength(1, 10, false), vax.In("wilson", "wendy"))
)
