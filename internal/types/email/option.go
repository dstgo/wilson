package email

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/pkg/vax"
)

type SendCodeEmailOption struct {
	// valid email format
	Email string `json:"email" uri:"email" form:"email" label:"field.email" example:"abc@example.com"`
}

func (e SendCodeEmailOption) Validate(lang string) error {
	return vax.Struct(&e, lang,
		vax.Field(&e.Email, helper.RequiredRules(RuleEmail)...),
	)
}
