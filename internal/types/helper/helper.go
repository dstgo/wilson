package helper

import (
	"github.com/dstgo/wilson/pkg/vax"
)

type PageOption struct {
	// page
	Page int `json:"page" uri:"page" form:"page"  label:"field.page" example:"1"`
	// size of page
	Size int `json:"size" uri:"size" form:"size" label:"field.pagesize" example:"10"`
	// if is reverse order
	Desc bool `json:"desc" uri:"desc" form:"desc" example:"true"`
}

func (p PageOption) Validate(lang string) error {
	return vax.Struct(&p, lang,
		vax.Field(&p.Page, vax.Required, vax.Gt(0)),
		vax.Field(&p.Size, vax.Required, vax.Gte(1)),
	)
}
