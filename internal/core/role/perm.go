package role

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/types/role"
	"github.com/dstgo/wilson/internal/types/system"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

func (g GormResolver) GetPerm(permId uint) (role.PermInfo, error) {
	perm, err := getPermById(g.db, permId)
	if err != nil {
		return role.PermInfo{}, err
	}
	return role.MakePermInfo(perm), nil
}

func (g GormResolver) GetPermInBatch(ids []uint, tag string) ([]role.PermInfo, error) {
	var perms []entity.Permission
	db := g.db
	if tag != "" {
		db = g.db.Where("tag = ?", tag)
	}
	err := db.Find(&perms, "id IN ?", ids).Error
	if err != nil {
		return []role.PermInfo{}, err
	}
	return role.MakePermInfoList(perms), nil
}

func (g GormResolver) MatchPerm(name, obj, act, group, tag string) (role.PermInfo, error) {
	var permInfo entity.Permission
	db := g.db
	if name != "" {
		db = db.Where("name = ?", name)
	}
	if group != "" {
		db = db.Where("group = ?", group)
	}
	if tag != "" {
		db = db.Where("tag = ?", tag)
	}
	err := db.Where("object = ?", obj).Where("action = ?", act).Find(&permInfo).Error
	return role.MakePermInfo(permInfo), err
}

func (g GormResolver) CreatePerm(permInfo role.PermInfo) error {
	_, err := createPerm(g.db, role.MakePermRecord(permInfo))
	if err != nil {
		return err
	}
	return nil
}

func (g GormResolver) CreatePermInBatch(perms []role.PermInfo) error {
	records := role.MakePermRecordList(perms)
	_, err := createPermInBatch(g.db, records)
	return err
}

func (g GormResolver) ListPerms(option role.PageOption) ([]role.PermInfo, error) {
	list, err := getPagePermList(g.db, option)
	if err != nil {
		return []role.PermInfo{}, err
	}
	return role.MakePermInfoList(list), nil
}

func (g GormResolver) ListAllPerms(tag string) ([]role.PermInfo, error) {
	perms, err := listPermsByTag(g.db, tag)
	if err != nil {
		return []role.PermInfo{}, err
	}
	return role.MakePermInfoList(perms), nil
}

func (g GormResolver) UpdatePerm(permInfo role.PermInfo) error {
	err := updatePerm(g.db, role.MakePermRecord(permInfo))
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

func getPermById(db *gorm.DB, id uint) (entity.Permission, error) {
	var perm entity.Permission
	found, err := data.HasRecordFound(db.Find(&perm, "id = ?", id))
	if err != nil {
		return perm, system.ErrDatabase.Wrap(err)
	} else if !found {
		return perm, nil
	}
	return perm, err
}

func getPermByName(db *gorm.DB, name string) (entity.Permission, error) {
	var perm entity.Permission
	result := db.Find(&perm, "name = ?", name)
	found, err := data.HasRecordFound(result)
	if err != nil {
		return perm, system.ErrDatabase.Wrap(err)
	} else if !found {
		return perm, nil
	}
	return perm, nil
}

func listPermsByTag(db *gorm.DB, tag string) ([]entity.Permission, error) {
	var perms []entity.Permission
	err := db.Model(entity.Permission{}).Find(&perms, "tag = ?", tag).Error
	if err != nil {
		return perms, system.ErrDatabase.Wrap(err)
	}
	return perms, nil
}

func listAllPerms(db *gorm.DB) ([]entity.Permission, error) {
	var perms []entity.Permission
	err := db.Find(&perms).Error
	if err != nil {
		return perms, system.ErrDatabase.Wrap(err)
	}
	return perms, nil
}

func listAllPermsByPerms(db *gorm.DB, perms []entity.Permission) ([]entity.Permission, error) {
	var (
		objs   []string
		acts   []string
		groups []string
		tags   []string
		ens    = make([]entity.Permission, 0, len(perms))
	)

	for _, perm := range perms {
		if !slices.Contains(objs, perm.Object) {
			objs = append(objs, perm.Object)
		}

		if !slices.Contains(acts, perm.Object) {
			acts = append(acts, perm.Action)
		}

		if !slices.Contains(groups, perm.Group) {
			groups = append(groups, perm.Group)
		}

		if !slices.Contains(tags, perm.Tag) {
			tags = append(tags, perm.Tag)
		}
	}

	db = db.Model(entity.Permission{})
	if len(objs) > 0 {
		db = db.Where("`object` IN ?", objs)
	}

	if len(acts) > 0 {
		db = db.Where("`action` IN ?", acts)
	}

	if len(groups) > 0 {
		db = db.Where("`group` IN ?", groups)
	}

	if len(tags) > 0 {
		db = db.Where("`tag` IN ?", tags)
	}

	err := db.Find(&ens).Error
	if err != nil {
		return ens, system.ErrDatabase.Wrap(err)
	}
	return ens, nil
}

func getPagePermList(db *gorm.DB, pageOpt role.PageOption) ([]entity.Permission, error) {
	pageDB := db
	pageDB.Scopes(data.Pages(pageOpt.Page, pageOpt.Size))
	if len(pageOpt.Search) > 0 {
		pageDB = pageDB.Where("name LIKE ?", data.Like(pageOpt.Search))
	}
	var perms []entity.Permission
	err := pageDB.Find(&perms).Error
	if err != nil {
		return perms, system.ErrDatabase.Wrap(err)
	}
	return perms, nil
}

func findPerm(db *gorm.DB, obj, act, g, t string) (entity.Permission, error) {
	var perm entity.Permission
	err := db.Model(entity.Permission{}).
		Where("object = ? AND action = ? AND `group` = ? AND tag = ?", obj, act, g, t).
		Find(&perm).Error
	if err != nil {
		return perm, system.ErrDatabase.Wrap(err)
	}
	return perm, nil
}

func createPerm(db *gorm.DB, roleInfo entity.Permission) (entity.Permission, error) {
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "object"}, {Name: "action"}, {Name: "group"}, {Name: "tag"}},
		DoNothing: true,
	}).Create(&roleInfo).Error
	if err != nil {
		return entity.Permission{}, system.ErrDatabase.Wrap(err)
	}
	return roleInfo, err
}

func createPermInBatch(db *gorm.DB, perms []entity.Permission) ([]entity.Permission, error) {
	// create permission batch if conflicting do nothing
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "object"}, {Name: "action"}, {Name: "group"}, {Name: "tag"}},
		DoNothing: true,
	}).Create(&perms).Error
	return perms, err
}

func removePerm(db *gorm.DB, permId uint) error {
	return db.Unscoped().Model(entity.Permission{}).Delete("id = ?", permId).Error
}

func updatePerm(db *gorm.DB, perm entity.Permission) error {
	return db.Model(perm).Where("id = ?", perm.Id).Updates(&perm).Error
}
