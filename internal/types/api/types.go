package api

import (
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
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

type UUID struct {
	UUID string `json:"uuid" uri:"uuid" form:"uuid" label:"field.uuid"`
}

func (u UUID) Validate(lang string) error {
	return vax.Struct(&u, lang,
		vax.Field(u.UUID, vax.Required, is.UUID),
	)
}

// Response
// just used to generate swagger api doc, you should use resp.Response instead
type Response struct {
	Code int    `json:"code" example:"2000"`
	Msg  string `json:"msg" example:"operation success"`
	Err  string `json:"err"`
	Data any    `json:"data"`
}
