package request

import (
	vax2 "github.com/dstgo/wilson/app/pkg/vax"
)

type IdString struct {
	Id string `json:"id" uri:"id" form:"id" label:"field.id"`
}

func (i IdString) Validate(lang string) error {
	return vax2.Struct(&i, lang,
		vax2.Field(i.Id, vax2.Required),
	)
}

type IdInt struct {
	Id int `json:"id" uri:"id" form:"id" label:"field.id"`
}

func (i IdInt) Validate(lang string) error {
	return vax2.Struct(&i, lang,
		vax2.Field(i.Id, vax2.Required),
	)
}
