package auth

import (
	rules2 "github.com/dstgo/wilson/internal/types/api/helper/rules"
	"github.com/dstgo/wilson/pkg/vax"
)

type LoginOption struct {
	// length of the username must be in range [6,20], and username must be Alphanumeric
	Username string `json:"username" label:"field.username" example:"123456@example.com/username"`
	// length of the password must be in range [10,30]
	Password string `json:"password" label:"field.password" example:"123456789"`
}

func (l LoginOption) Validate(lang string) error {
	return vax.Struct(&l, lang,
		vax.Field(&l.Username, rules2.Required(rules2.Username)...),
		vax.Field(&l.Password, rules2.Required(rules2.Password)...),
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
		vax.Field(&r.Code, rules2.Required(rules2.EmailCode)...),
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
		vax.Field(&f.Password, rules2.Required(rules2.Username)...),
		vax.Field(&f.Code, rules2.Required(rules2.EmailCode)...),
	)
}
