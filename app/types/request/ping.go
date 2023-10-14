package request

import (
	vax2 "github.com/dstgo/wilson/app/pkg/vax"
)

type PingRequest struct {
	Name string `json:"name" uri:"name" form:"name" label:"field.name"`
}

func (p PingRequest) Validate(lang string) error {
	return vax2.Struct(&p, lang,
		vax2.Field(&p.Name, vax2.Required, vax2.RangeLength(1, 10, false), vax2.In("wilson", "wendy")),
	)
}
