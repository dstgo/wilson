package email

import (
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
)

type Email struct {
	// valid email format
	Email string `json:"email" uri:"email" form:"email" label:"field.email" example:"abc@example.com"`
}

func (e Email) Validate(lang string) error {
	return vax.Struct(&e, lang,
		vax.Field(&e.Email, vax.Required, is.Email),
	)
}
