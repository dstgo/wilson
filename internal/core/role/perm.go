package role

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/role"
	"gorm.io/gorm"
)

func (g GormResolver) GetPerm(permId uint) (role.PermInfo, error) {
	perm, err := getPermById(g.db, permId)
	if err != nil {
		return role.PermInfo{}, err
	}
	return makePermInfo(perm), nil
}

func (g GormResolver) CreatePerm(permInfo role.PermInfo) error {
	err := createPerm(g.db, makePermEntity(permInfo))
	if err != nil {
		return err
	}
	return nil
}

func (g GormResolver) ListPerms(option role.PageOption) ([]role.PermInfo, error) {
	list, err := getPagePermList(g.db, option)
	if err != nil {
		return []role.PermInfo{}, err
	}
	return makePermInfoList(list), nil
}

func (g GormResolver) ListAllPerms(tag string) ([]role.PermInfo, error) {
	perms, err := listPermsByTag(g.db, tag)
	if err != nil {
		return []role.PermInfo{}, err
	}
	return makePermInfoList(perms), nil
}

func (g GormResolver) UpdatePerm(permInfo role.PermInfo) error {
	err := updatePerm(g.db, makePermEntity(permInfo))
	if err != nil {
		return err
	}
	return nil
}

func (g GormResolver) RemovePerm(permId uint) error {
	err := removePerm(g.db, permId)
	if err != nil {
		return err
	}
	return nil
}

func makePermInfo(perm entity.Permission) role.PermInfo {
	return role.PermInfo{
		ID:     perm.ID,
		Name:   perm.Name,
		Object: perm.Object,
		Group:  perm.Group,
		Action: perm.Action,
		Tag:    perm.Tag,
	}
}

func makePermInfoList(perms []entity.Permission) (infos []role.PermInfo) {
	for _, perm := range perms {
		infos = append(infos, makePermInfo(perm))
	}
	return
}

func makePermEntity(perm role.PermInfo) entity.Permission {
	return entity.Permission{
		Model:  gorm.Model{ID: perm.ID},
		Name:   perm.Name,
		Object: perm.Object,
		Action: perm.Action,
		Group:  perm.Group,
		Tag:    perm.Tag,
	}
}

func makePermEntityList(perms []role.PermInfo) (ens []entity.Permission) {
	for _, perm := range perms {
		ens = append(ens, makePermEntity(perm))
	}
	return
}

func makeRolePerms(roleId uint, permIds []uint) []entity.RolePermission {
	rolePermList := make([]entity.RolePermission, 0, len(permIds))
	for _, permId := range permIds {
		rolePermList = append(rolePermList, entity.RolePermission{RoleId: roleId, PermId: permId})
	}
	return rolePermList
}

func getPermById(db *gorm.DB, id uint) (entity.Permission, error) {
	var perm entity.Permission
	err := db.Find(&perm, "id = ?", id).Error
	return perm, err
}

func getPermByName(db *gorm.DB, name string) (entity.Permission, error) {
	var perm entity.Permission
	err := db.Find(&perm, "name = ?", name).Error
	return perm, err
}

func listPermsByTag(db *gorm.DB, tag string) ([]entity.Permission, error) {
	var perms []entity.Permission
	err := db.Model(entity.Permission{}).Find(&perms, "tag = ?", tag).Error
	return perms, err
}

func listAllPerms(db *gorm.DB) ([]entity.Permission, error) {
	var perms []entity.Permission
	err := db.Find(&perms).Error
	return perms, err
}

func getPagePermList(db *gorm.DB, pageOpt role.PageOption) ([]entity.Permission, error) {
	pageDB := db
	pageDB.Scopes(data.Pages(pageOpt.Page, pageOpt.Size))
	if len(pageOpt.Search) > 0 {
		pageDB = pageDB.Where("name LIKE ?", data.Like(pageOpt.Search))
	}
	var perms []entity.Permission
	err := pageDB.Find(&perms).Error
	return perms, err
}

func findPerm(db *gorm.DB, obj, act, g, t string) (entity.Permission, error) {
	var perm entity.Permission
	err := db.Model(entity.Permission{}).
		Where("object = ? AND action = ? AND `group` = ? AND tag = ?", obj, act, g, t).
		Find(&perm).Error
	return perm, err
}

func createPerm(db *gorm.DB, role entity.Permission) error {
	return db.Create(&role).Error
}

func removePerm(db *gorm.DB, permId uint) error {
	return db.Unscoped().Model(entity.Permission{}).Delete("id = ?", permId).Error
}

func updatePerm(db *gorm.DB, perm entity.Permission) error {
	return db.Model(perm).Where("id = ?", perm.ID).Updates(&perm).Error
}

func makePermGroup(perms []entity.Permission) []role.PermGroup {
	pg := make(map[string][]role.PermInfo, len(perms)/10)

	for _, perm := range perms {
		permInfo := role.PermInfo{Name: perm.Name, Object: perm.Object, Action: perm.Action}
		if _, e := pg[perm.Group]; !e {
			pg[perm.Group] = []role.PermInfo{permInfo}
		} else {
			pg[perm.Group] = append(pg[perm.Group], permInfo)
		}
	}

	var groups []role.PermGroup

	for groupName, perms := range pg {
		groups = append(groups, role.PermGroup{Group: groupName, Perms: perms})
	}

	return groups
}
