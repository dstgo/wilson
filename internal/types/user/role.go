package user

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/vax"
)

type SaveRoleOption struct {
	types.Uid
	RoleIds []uint `json:"role_ids" label:"field.role.list"`
}

func (s SaveRoleOption) Validate(lang string) error {
	return vax.Struct(&s, lang,
		vax.Field(&s.Uid),
		vax.Field(&s.RoleIds, vax.Required, vax.MinLength(1, false)),
	)
}
