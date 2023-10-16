package auth

import (
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
)

type LoginOption struct {
	// length of the username must be in range [6,20], and username must be Alphanumeric
	Username string `json:"username" label:"field.username" example:"123456@example.com/username"`
	// length of the password must be in range [10,30]
	Password string `json:"password" label:"field.password" example:"123456789"`
}

func (l LoginOption) Validate(lang string) error {
	return vax.Struct(&l, lang,
		vax.Field(&l.Username, vax.Required, is.Alphanumeric, vax.RangeLength(6, 20, false)),
		vax.Field(&l.Password, vax.Required, vax.RangeLength(10, 30, false)),
	)
}

type RegisterOption struct {
	LoginOption
	// 8-digit auth code get from user binding email, and it must be Alphanumeric
	Code string `json:"code" label:"field.code" example:"F294484D"`
}

func (r RegisterOption) Validate(lang string) error {
	return vax.Struct(&r, lang,
		vax.Field(&r.LoginOption),
		vax.Field(&r.Code, vax.Required, is.Alphanumeric),
	)
}

type ForgotPasswordOption struct {
	// length of the password must be in range [10,30]
	Password string `json:"password" label:"field.password" example:"123456789"`
	// auth code get from user binding email
	Code string `json:"code" label:"field.code" example:"F294484D"`
}

func (f ForgotPasswordOption) Validate(lang string) error {
	return vax.Struct(&f, lang,
		vax.Field(&f.Password, vax.Required, is.Alphanumeric, vax.RangeLength(10, 30, false)),
		vax.Field(&f.Code, vax.Required, is.Alphanumeric),
	)
}
