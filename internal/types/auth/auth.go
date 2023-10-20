package auth

import (
	"github.com/dstgo/wilson/internal/types/helper/rules"
	"github.com/dstgo/wilson/pkg/vax"
)

type PingReply struct {
	Reply string `json:"reply" example:"hello wendy! Now is 2023-10-17 11:07:21.696 +08:00."`
}

type LoginOption struct {
	// length of the username must be in range [6,20], and username must be Alphanumeric
	Username string `json:"username" label:"field.username" example:"123456@example.com/username"`
	// length of the password must be in range [10,30]
	Password string `json:"password" label:"field.password" example:"123456789"`
}

func (l LoginOption) Validate(lang string) error {
	return vax.Struct(&l, lang,
		vax.Field(&l.Username, rules.Required(rules.Username)...),
		vax.Field(&l.Password, rules.Required(rules.Password)...),
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
		vax.Field(&r.Code, rules.Required(rules.EmailCode)...),
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
		vax.Field(&f.Password, rules.Required(rules.Username)...),
		vax.Field(&f.Code, rules.Required(rules.EmailCode)...),
	)
}
