package request

import (
	"github.com/dstgo/wilson/pkg/vax"
)

type IdString struct {
	Id string `json:"id" uri:"id" form:"id" label:"field.id"`
}

func (i IdString) Validate(lang string) error {
	return vax.Struct(&i, lang,
		vax.Field(i.Id, vax.Required),
	)
}

type IdInt struct {
	Id int `json:"id" uri:"id" form:"id" label:"field.id"`
}

func (i IdInt) Validate(lang string) error {
	return vax.Struct(&i, lang,
		vax.Field(i.Id, vax.Required),
	)
}
