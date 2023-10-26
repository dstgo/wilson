package user

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/internal/types/helper/rules"
	"github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/dstgo/wilson/pkg/vax/is"
)

type Info struct {
	UUID      string          `json:"uuid" example:"55BBA4ED-18D3-790F-EABF-A5330E527586"`
	Username  string          `json:"username" example:"jack"`
	Email     string          `json:"email" example:"jacklove@lol.com"`
	CreatedAt uint64          `json:"createdAt" example:"947416200"`
	Roles     []role.RoleInfo `json:"roles"`
}

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
		vax.Field(&u.Username, rules.Username...),
		vax.Field(&u.Email, rules.Email...),
		vax.Field(&u.Password, vax.When(len(u.Password) > 0, rules.Password...)),
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
		vax.Field(&u.UUID, vax.Required, is.UUID),
		vax.Field(&u.Username, rules.Required(rules.Username)...),
		vax.Field(&u.Email, rules.Required(rules.Email)...),
		vax.Field(&u.Password, rules.Required(rules.Password)...),
		vax.Field(&u.Roles, vax.Required),
	)
}

var InitialUser = CreateUserOption{
	Username: "dstadmin",
	Email:    "",
	Password: "0123456789",
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
		vax.Field(&c.Username, rules.Required(rules.Username)...),
		vax.Field(&c.Email, rules.Required(rules.Email)...),
		vax.Field(&c.Password, rules.Required(rules.Password)...),
		vax.Field(&c.Roles, vax.Required),
	)
}
