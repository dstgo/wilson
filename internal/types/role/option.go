package role

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/dstgo/wilson/pkg/vax"
)

type PageOption struct {
	helper.PageOption
	Search string `json:"search" uri:"search" form:"search" example:"admin"`
	Tag    string `json:"tag" uri:"tag" form:"tag" example:"appapi"`
}

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
		vax.Field(&c.Name, helper.RequiredRules(RuleRoleName)...),
		vax.Field(&c.Object, helper.RequiredRules(RulePermObj)...),
		vax.Field(&c.Action, helper.RequiredRules(RulePermAct)...),
		vax.Field(&c.Group, helper.RequiredRules(RulePermGroup)...),
		vax.Field(&c.Tag, helper.RequiredRules(RulePermTag)...),
	)
}

type UpdatePermOption struct {
	Id   uint   `json:"id" example:"1"`
	Name string `json:"name" label:"field.perm.name" example:"updateUser"`
}

func (u UpdatePermOption) Validate(lang string) error {
	return vax.Struct(&u, lang,
		vax.Field(&u.Id, vax.Required),
		vax.Field(&u.Name, helper.RequiredRules(RuleRoleName)...))
}

type CreateRoleOption struct {
	// role name
	Name string `json:"name" label:"field.role.name" example:"admin"`
	// role code, must be alpha numeric
	Code string `json:"code" label:"field.role.code" example:"ADMIN"`
}

func (c CreateRoleOption) Validate(lang string) error {
	return vax.Struct(&c, lang,
		vax.Field(&c.Name, helper.RequiredRules(RuleRoleName)...),
		vax.Field(&c.Code, helper.RequiredRules(RuleRoleCode)...),
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
		vax.Field(&u.Name, helper.RequiredRules(RuleRoleName)...),
	)
}

type GrantOption struct {
	RoleId  uint   `json:"roleId" label:"field.role.id" example:"1"`
	Tag     string `json:"tag" label:"field.role.tag" example:"AppApi"`
	PermIds []uint `json:"permId" label:"field.perm.list"`
}

func (g GrantOption) Validate(lang string) error {
	return vax.Struct(&g, lang,
		vax.Field(&g.RoleId, vax.Required),
		vax.Field(&g.Tag, helper.RequiredRules(RulePermTag)...),
		vax.Field(&g.PermIds, vax.Required),
	)
}

type QueryRolePermsOption struct {
	system.Id
	Tag string `json:"tag" uri:"tag" form:"tag" label:"field.perm.tag" example:"AppAPI"`
}

func (q QueryRolePermsOption) Validate(lang string) error {
	return vax.Struct(&q, lang,
		vax.Field(&q.Id),
		vax.Field(q.Tag, helper.RequiredRules(RulePermTag)...),
	)
}
