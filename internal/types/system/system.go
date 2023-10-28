package system

import (
	"github.com/dstgo/wilson/pkg/vax"
)

const (
	AppAPI = "appapi"

	OpenAPI = "openapi"
)

type PingRequest struct {
	// name must be one of [wilson, wendy]
	Name string `json:"name" uri:"name" form:"name" label:"field.name" example:"wilson"`
}

func (p PingRequest) Validate(lang string) error {
	return vax.Struct(&p, lang,
		vax.Field(&p.Name, vax.Required, vax.RangeLength(1, 10, false), vax.In("wilson", "wendy")),
	)
}
