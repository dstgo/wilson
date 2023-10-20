package rules

import (
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
)

var (
	Email = Rules(is.Email)

	EmailCode = Rules(is.Alphanumeric, vax.EqLength(8, false))
)
