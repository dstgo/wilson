package request

import (
	vax2 "github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
)

type LoginRequest struct {
	Username string `json:"username" label:"field.username"`
	Password string `json:"password" label:"field.password"`
}

func (l LoginRequest) Validate(lang string) error {
	return vax2.Struct(&l, lang,
		vax2.Field(&l.Username, vax2.Required, is.Alphanumeric, vax2.RangeLength(6, 20, false)),
		vax2.Field(&l.Password, vax2.Required, is.Alphanumeric, vax2.RangeLength(10, 30, false)),
	)
}

type RegisterRequest struct {
	LoginRequest
	Code string `json:"code" label:"field.code"`
}

func (r RegisterRequest) Validate(lang string) error {
	return vax2.Struct(&r, lang,
		vax2.Field(&r.LoginRequest),
		vax2.Field(&r.Code, vax2.Required, is.Alphanumeric),
	)
}

type ForgotPasswordRequest struct {
	Password string `json:"password" label:"field.password"`
	Code     string `json:"code" label:"field.code"`
}

func (f ForgotPasswordRequest) Validate(lang string) error {
	return vax2.Struct(&f, lang,
		vax2.Field(&f.Password, vax2.Required, is.Alphanumeric, vax2.RangeLength(10, 30, false)),
		vax2.Field(&f.Code, vax2.Required, is.Alphanumeric),
	)
}
