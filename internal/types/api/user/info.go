package user

import (
	"github.com/dstgo/wilson/internal/types/api/helper"
	"github.com/dstgo/wilson/internal/types/api/helper/rules"
	"github.com/dstgo/wilson/pkg/vax"
)

type Info struct {
	UUID     string `json:"uuid" example:"55BBA4ED-18D3-790F-EABF-A5330E527586"`
	Username string `json:"username" example:"jack"`
	Email    string `json:"email" example:"jacklove@lol.com"`
	// used for copy from gorm.model
	helper.CreatedAt `copier:"Model"`
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
	UUID string `json:"-" swaggerignore:"true"`
	// new username
	Username string `json:"username" example:"jack"`
	// new email
	Email string `json:"email" example:"jack@google.com"`
	// new password
	Password string `json:"password" example:"123456"`
}

func (u UpdateInfoOption) Validate(lang string) error {
	return vax.Struct(&u, lang,
		vax.Field(&u.Username, rules.Required(rules.Username)...),
		vax.Field(&u.Email, rules.Required(rules.Email)...),
		vax.Field(&u.Password, rules.Required(rules.Password)...),
	)
}
