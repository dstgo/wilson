package request

import (
	vax2 "github.com/dstgo/wilson/app/pkg/vax"
	"github.com/dstgo/wilson/app/pkg/vax/is"
)

type Email struct {
	Email string `json:"email" uri:"email" form:"email" label:"field.email"`
}

func (e Email) Validate(lang string) error {
	return vax2.Struct(&e, lang,
		vax2.Field(&e.Email, vax2.Required, is.Email),
	)
}
