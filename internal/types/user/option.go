package user

import (
	"github.com/dstgo/wilson/internal/types/email"
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/pkg/vax"
)

type PageOption struct {
	helper.PageOption
	// specified field
	Order string `json:"order" uri:"order" form:"order" example:"email"`
	// search text, should be of one username or email
	Search string `json:"search" uri:"search" form:"search" example:"jacklove"`
}

func (p PageOption) Validate(lang string) error {
	return vax.Struct(&p, lang,
		vax.Field(&p.PageOption),
	)
}

type CreateUserOption struct {
	// new username
	Username string `json:"username" example:"jack"`
	// new email
	Email string `json:"email" example:"jack@google.com"`
	// new password
	Password string `json:"password" example:"123456"`
	// new roles
	Roles []string `json:"roles"`
}

func (c CreateUserOption) Validate(lang string) error {
	return vax.Struct(&c, lang,
		vax.Field(&c.Username, helper.RequiredRules(RuleUsername)...),
		vax.Field(&c.Email, helper.RequiredRules(email.RuleEmail)...),
		vax.Field(&c.Password, helper.RequiredRules(RulePassword)...),
		vax.Field(&c.Roles, vax.Required),
	)
}

type UpdateInfoOption struct {
	UUID string `json:"-" swaggerignore:"true" example:"55BBA4ED-18D3-790F-EABF-A5330E527586"`
	// new username
	Username string `json:"username" example:"jack"`
	// new email
	Email string `json:"email" example:"jack@google.com"`
	// new password
	Password string `json:"password" example:"123456"`
}

func (u UpdateInfoOption) Validate(lang string) error {
	return vax.Struct(&u, lang,
		vax.Field(&u.Username, RuleUsername...),
		vax.Field(&u.Email, email.RuleEmail...),
		vax.Field(&u.Password, vax.When(u.Password != "", RulePassword...)),
	)
}

type SaveUserDetailOption struct {
	UUID string `json:"uuid" example:"55BBA4ED-18D3-790F-EABF-A5330E527586"`
	// new username
	Username string `json:"username" example:"jack"`
	// new email
	Email string `json:"email" example:"jack@google.com"`
	// new password
	Password string `json:"password" example:"123456"`
	// new roles
	Roles []string `json:"roles"`
}

func (u SaveUserDetailOption) Validate(lang string) error {
	return vax.Struct(&u, lang,
		vax.Field(&u.UUID, vax.Required),
		vax.Field(&u.Username, helper.RequiredRules(RuleUsername)...),
		vax.Field(&u.Email, helper.RequiredRules(email.RuleEmail)...),
		vax.Field(&u.Password, helper.RequiredRules(RulePassword)...),
		vax.Field(&u.Roles, vax.Required),
	)
}
