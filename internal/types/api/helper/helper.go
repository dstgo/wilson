package helper

import (
	"github.com/dstgo/wilson/pkg/vax"
	"time"
)

type CreatedAt struct {
	CreatedAt time.Time `json:"createdAt" example:"2022-01-01 00:00:00"`
}

type UpdatedAt struct {
	UpdatedAt time.Time `json:"updatedAt" example:"2022-01-01 00:00:00"`
}

type PageOption struct {
	// page
	Page int `json:"page" label:"field.page" example:"1"`
	// size of page
	Size int `json:"size" label:"field.pagesize" example:"10"`
	// if is reverse order
	Desc bool `json:"desc" example:"true"`
}

func (p PageOption) Validate(lang string) error {
	return vax.Struct(&p, lang,
		vax.Field(p.Page, vax.Required, vax.Gt(1)),
		vax.Field(p.Size, vax.Required, vax.Gte(1)),
	)
}
