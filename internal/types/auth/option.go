package auth

import (
	"github.com/dstgo/wilson/internal/types/helper/rules"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/dstgo/wilson/pkg/vax"
)

type LoginOption struct {
	// length of the username must be in range [6,20], and username must be Alphanumeric
	Username string `json:"username" label:"field.username" example:"123456@example.com/username"`
	// length of the password must be in range [10,30]
	Password string `json:"password" label:"field.password" example:"123456789"`

	Persistent bool `json:"persistent" example:"true"`
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

type RefreshTokenOption struct {
	Refresh string `json:"refresh" form:"refresh" uri:"refresh" label:"field.token.access"`
	Access  string `json:"access" form:"access" uri:"access" label:"field.token.refresh"`
}

func (r RefreshTokenOption) Validate(lang string) error {
	return vax.Struct(&r, lang,
		vax.Field(&r.Refresh, vax.Required),
		vax.Field(&r.Access, vax.Required),
	)
}

type KeyCreateOption struct {
	Uid       string `json:"-" swaggerignore:"true"`
	Name      string `json:"name"`
	Perms     []uint `json:"perms"`
	ExpiredAt int64  `json:"expiredAt"`
}

func (c KeyCreateOption) Validate(lang string) error {
	return vax.Struct(&c, lang,
		vax.Field(&c.Name, vax.Required, vax.RangeLength(1, 50, false)),
		vax.Field(&c.Perms, vax.Required),
		vax.Field(&c.ExpiredAt, vax.Required, vax.Gt(0)),
	)
}

type KeyRemoveOption struct {
	system.Uid `json:"-" swaggerignore:"true"`
	Key        string `json:"key" uri:"key" form:"key"`
}

func (c KeyRemoveOption) Validate(lang string) error {
	return vax.Struct(&c, lang,
		vax.Field(&c.Uid),
		vax.Field(&c.Key, vax.Required),
	)
}
