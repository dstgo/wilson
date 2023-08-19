package request

import (
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/core/vax/is"
)

type Email struct {
	Email string `json:"email" uri:"email" form:"email" label:"field.email"`
}

func (e Email) Validate(lang string) error {
	return vax.Struct(&e, lang,
		vax.Field(&e.Email, vax.Required, is.Email),
	)
}
