package rules

import (
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
)

var (
	Username = Rules(is.Alphanumeric, vax.RangeLength(6, 20, false))

	Password = Rules(vax.RangeLength(10, 30, false))
)
