package request

import (
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/core/vax/is"
)

type LoginRequest struct {
	Username string `json:"username" label:"field.username"`
	Password string `json:"password" label:"field.password"`
}

func (l LoginRequest) Validate(lang string) error {
	return vax.Struct(&l, lang,
		vax.Field(&l.Username, vax.Required, is.Alphanumeric, vax.RangeLength(6, 20, false)),
		vax.Field(&l.Password, vax.Required, is.Alphanumeric, vax.RangeLength(10, 30, false)),
	)
}

type RegisterRequest struct {
	LoginRequest
	Code string `json:"code" label:"field.code"`
}

func (r RegisterRequest) Validate(lang string) error {
	return vax.Struct(&r, lang,
		vax.Field(&r.LoginRequest),
		vax.Field(&r.Code, vax.Required, is.Alphanumeric),
	)
}
