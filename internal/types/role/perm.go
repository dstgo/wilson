package role

import (
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/helper/rules"
	"github.com/dstgo/wilson/pkg/vax"
	"gorm.io/gorm"
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

type PermGroup struct {
	// group name
	Group string     `json:"group" example:"user group"`
	Perms []PermInfo `json:"perms"`
}

type PermInfo struct {
	// permission id
	ID uint `json:"id" example:"1"`
	// permission name
	Name string `json:"name" example:"updateUser"`
	// define the object will be accessed
	Object string `json:"object" example:"/user/update"`
	// permission group
	Group string `json:"group" example:"admin"`
	// how to access the object
	Action string `json:"action" example:"POST"`
	// tag of permissions
	Tag string `json:"tag" example:"AppAPI"`
}

func MakePermGroup(perms []entity.Permission) []PermGroup {
	pg := make(map[string][]PermInfo, len(perms)/10)

	for _, perm := range perms {
		permInfo := PermInfo{Name: perm.Name, Object: perm.Object, Action: perm.Action}
		if _, e := pg[perm.Group]; !e {
			pg[perm.Group] = []PermInfo{permInfo}
		} else {
			pg[perm.Group] = append(pg[perm.Group], permInfo)
		}
	}

	var groups []PermGroup

	for groupName, perms := range pg {
		groups = append(groups, PermGroup{Group: groupName, Perms: perms})
	}

	return groups
}

func MakePermInfo(perm entity.Permission) PermInfo {
	return PermInfo{
		ID:     perm.ID,
		Name:   perm.Name,
		Object: perm.Object,
		Group:  perm.Group,
		Action: perm.Action,
		Tag:    perm.Tag,
	}
}

func MakePermInfoList(perms []entity.Permission) (infos []PermInfo) {
	for _, perm := range perms {
		infos = append(infos, MakePermInfo(perm))
	}
	return
}

func MakePermRecord(perm PermInfo) entity.Permission {
	return entity.Permission{
		Model:  gorm.Model{ID: perm.ID},
		Name:   perm.Name,
		Object: perm.Object,
		Action: perm.Action,
		Group:  perm.Group,
		Tag:    perm.Tag,
	}
}

func MakePermRecordList(perms []PermInfo) (ens []entity.Permission) {
	for _, perm := range perms {
		ens = append(ens, MakePermRecord(perm))
	}
	return
}
