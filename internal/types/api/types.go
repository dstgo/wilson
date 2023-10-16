package api

import (
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/spf13/cast"
)

// Id
// represent query or path param ID
type Id struct {
	Id string `json:"id" uri:"id" form:"id" label:"field.id"`
}

func (i Id) Int() int {
	return cast.ToInt(i.Id)
}

func (i Id) Uint() uint {
	return cast.ToUint(i.Id)
}

func (i Id) String() string {
	return cast.ToString(i.Id)
}

func (i Id) Validate(lang string) error {
	return vax.Struct(&i, lang,
		vax.Field(i.Id, vax.Required),
	)
}
