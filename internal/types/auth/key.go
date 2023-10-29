package auth

import (
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/pkg/vax"
)

type CreateKeyOption struct {
	Uid       string `json:"-" swaggerignore:"true"`
	Name      string `json:"name"`
	Perms     []uint `json:"perms"`
	ExpiredAt int64  `json:"expiredAt"`
}

func (c CreateKeyOption) Validate(lang string) error {
	return vax.Struct(&c, lang,
		vax.Field(&c.Name, vax.Required, vax.RangeLength(1, 50, false)),
		vax.Field(&c.Perms, vax.Required),
		vax.Field(&c.ExpiredAt, vax.Required, vax.Gt(0)),
	)
}

type RemoveKeyOption struct {
	types.Uid `json:"-" swaggerignore:"true"`
	Key       string `json:"key" uri:"key" form:"key"`
}

func (c RemoveKeyOption) Validate(lang string) error {
	return vax.Struct(&c, lang,
		vax.Field(&c.Uid),
		vax.Field(&c.Key, vax.Required),
	)
}

type APIKey struct {
	Name      string `json:"name"`
	Key       string `json:"key"`
	ExpiredAt uint64 `json:"expiredAt"`
}
