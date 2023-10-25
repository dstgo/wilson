package role

import (
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types"
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/internal/types/helper/rules"
	"github.com/dstgo/wilson/pkg/vax"
)

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
	RoleId  uint   `json:"roleId" label:"field.role.id" example:"1"`
	Tag     string `json:"tag" label:"field.role.tag" example:"AppApi"`
	PermIds []uint `json:"permId" label:"field.perm.list"`
}

func (g GrantOption) Validate(lang string) error {
	return vax.Struct(&g, lang,
		vax.Field(&g.RoleId, vax.Required),
		vax.Field(&g.Tag, rules.Required(rules.PermTag)...),
		vax.Field(&g.PermIds, vax.Required),
	)
}

type PageOption struct {
	helper.PageOption
	Search string `json:"search" uri:"search" form:"search" example:"admin"`
}

type QueryRolePermsOption struct {
	types.Id
	Tag string `json:"tag" uri:"tag" form:"tag" label:"field.perm.tag" example:"AppAPI"`
}

func (q QueryRolePermsOption) Validate(lang string) error {
	return vax.Struct(&q, lang,
		vax.Field(&q.Id),
		vax.Field(q.Tag, rules.Required(rules.PermTag)...),
	)
}

type RoleInfo struct {
	// role id
	Id uint `json:"id" example:"1"`
	// role name
	Name string `json:"name" example:"admin"`
	// role code, must be alpha numeric
	Code string `json:"code" example:"ADMIN"`
}

var (
	// AdminRole app static admin role,
	AdminRole = RoleInfo{
		Name: "Admin",
		Code: "1024",
	}

	UserRole = RoleInfo{
		Name: "User",
		Code: "0512",
	}
	AnonymousRole = RoleInfo{
		Name: "Guest",
		Code: "0000",
	}
)

func MakeRoleInfo(record entity.Role) RoleInfo {
	return RoleInfo{
		Id:   record.Id,
		Name: record.Name,
		Code: record.Code,
	}
}

func MakeRoleInfoList(records []entity.Role) (infos []RoleInfo) {
	for _, record := range records {
		infos = append(infos, MakeRoleInfo(record))
	}
	return
}

func MakeRoleRecord(info RoleInfo) entity.Role {
	return entity.Role{
		Id:   info.Id,
		Name: info.Name,
		Code: info.Code,
	}
}

func MakeRoleRecordList(infos []RoleInfo) (records []entity.Role) {
	for _, info := range infos {
		records = append(records, MakeRoleRecord(info))
	}
	return
}

func MakeUserRoleRecordList(userId uint, roleIds []uint) []entity.UserRole {
	var records []entity.UserRole
	for _, id := range roleIds {
		records = append(records, entity.UserRole{
			UserId: userId,
			RoleId: id,
		})
	}
	return records
}
