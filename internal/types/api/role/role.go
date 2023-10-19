package role

import (
	"github.com/dstgo/wilson/internal/types/api"
	"github.com/dstgo/wilson/internal/types/api/helper"
	"github.com/dstgo/wilson/internal/types/api/helper/rules"
	"github.com/dstgo/wilson/pkg/vax"
)

type CreatePermOption struct {
	Name string `json:"name" label:"field.perm.name" example:"updateUser"`
	// define the object will be accessed
	Object string `json:"object" label:"field.perm.object" example:"/user/update"`
	// how to access the object
	Action string `json:"action" label:"field.perm.action" example:"POST"`
	// permission group
	Group string `json:"group" label:"field.perm.group" example:"UserGroup"`
	// tag of permission
	Tag string `json:"tag" label:"field.perm.tag" example:"AppAPI"`
}

func (c CreatePermOption) Validate(lang string) error {
	return vax.Struct(&c, lang,
		vax.Field(&c.Name, rules.Required(rules.RoleName)...),
		vax.Field(&c.Object, rules.Required(rules.PermObject)...),
		vax.Field(&c.Action, rules.Required(rules.PermAction)...),
		vax.Field(&c.Group, vax.Required, vax.RangeLength(1, 60, false)),
		vax.Field(&c.Tag, rules.Required(rules.PermTag)...),
	)
}

type UpdatePermOption struct {
	Id   uint   `json:"id" example:"1"`
	Name string `json:"name" label:"field.perm.name" example:"updateUser"`
}

func (u UpdatePermOption) Validate(lang string) error {
	return vax.Struct(&u, lang,
		vax.Field(&u.Id, vax.Required),
		vax.Field(&u.Name, rules.Required(rules.RoleName)...))
}

type CreateRoleOption struct {
	// role name
	Name string `json:"name" label:"field.role.name" example:"admin"`
	// role code, must be alpha numeric
	Code string `json:"code" label:"field.role.code" example:"ADMIN"`
}

func (c CreateRoleOption) Validate(lang string) error {
	return vax.Struct(&c, lang,
		vax.Field(&c.Name, rules.Required(rules.RoleName)...),
		vax.Field(&c.Code, rules.Required(rules.RoleCode)...),
	)
}

type UpdateRoleOption struct {
	Id uint `json:"id" example:"1"`
	// role name
	Name string `json:"name" label:"field.role.name" example:"admin"`
}

func (u UpdateRoleOption) Validate(lang string) error {
	return vax.Struct(&u, lang,
		vax.Field(&u.Id, vax.Required),
		vax.Field(&u.Name, rules.Required(rules.RoleName)...),
	)
}

type GrantOption struct {
	RoleId uint   `json:"roleId" label:"field.role.id" example:"1"`
	Tag    string `json:"tag" label:"field.role.tag" example:"AppApi"`
	PermId []uint `json:"permId" label:"field.perm.id"`
}

func (g GrantOption) Validate(lang string) error {
	return vax.Struct(&g, lang,
		vax.Field(&g.RoleId, vax.Required),
		vax.Field(&g.Tag, rules.Required(rules.PermTag)...),
		vax.Field(&g.PermId, vax.Required),
	)
}

type PageOption struct {
	helper.PageOption
	Search string `json:"search" uri:"search" form:"search" example:"admin"`
}

type QueryRolePermsOption struct {
	api.Id
	Tag string `json:"tag" uri:"tag" form:"tag" label:"field.perm.tag" example:"AppAPI"`
}

func (q QueryRolePermsOption) Validate(lang string) error {
	return vax.Struct(&q, lang,
		vax.Field(&q.Id),
		vax.Field(q.Tag, rules.Required(rules.PermTag)...),
	)
}
