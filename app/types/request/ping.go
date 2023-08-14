package request

import "github.com/dstgo/wilson/app/core/vax"

type PingRequest struct {
	Name string `json:"name" uri:"name" form:"name" label:"field.name"`
}

func (p PingRequest) Validate(lang string) error {
	return vax.Struct(&p, lang,
		vax.Field(&p.Name, vax.Required, vax.RangeLength(1, 10, false), vax.In("wilson", "wendy", "uzi", "faker")),
	)
}
