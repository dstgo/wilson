package role

import (
	"errors"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/alg/collection"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/api/role"
	"github.com/dstgo/wilson/internal/types/errs"
	"gorm.io/gorm"
)

// the predefine roles
const (
	AdminRole = "admin"

	UserRole = "user"

	AnonymousRole = "anonymous"
)

func NewRoleService(ds *data.DataSource, perm PermData, role RoleData) RoleService {
	return RoleService{ds: ds, permData: perm, roleData: role}
}

type RoleService struct {
	permData PermData
	roleData RoleData
	ds       *data.DataSource
}

func (r RoleService) CreateRole(createOpt role.CreateRoleOption) error {
	// check if already exists
	record, err := r.roleData.GetRoleByCode(r.ds.ORM(), createOpt.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.DataBaseErr(err)
	} else if record.ID > 0 {
		return errs.NewI18nError("role.codeConflict")
	}

	roleEntity := entity.Role{Name: createOpt.Name, Code: createOpt.Code}
	if err := r.roleData.CreateRole(r.ds.ORM(), roleEntity); err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (r RoleService) GetRoleById(id uint) (role.RoleInfo, error) {
	record, err := r.roleData.GetRoleById(r.ds.ORM(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return role.RoleInfo{}, errs.NewI18nError("role.notfound")
	} else if err != nil {
		return role.RoleInfo{}, errs.DataBaseErr(err)
	}
	return role.RoleInfo{Name: record.Name, Code: record.Code}, nil
}

func (r RoleService) ListRoleByPage(option role.PageOption) ([]role.RoleInfo, error) {
	list, err := r.roleData.GetPageRoleList(r.ds.ORM(), option)
	if err != nil {
		return []role.RoleInfo{}, errs.DataBaseErr(err)
	}

	var roleList []role.RoleInfo
	if err := cp.Copy(&list, &roleList); err != nil {
		return []role.RoleInfo{}, errs.ProgramErr(err)
	}

	return roleList, nil
}

func (r RoleService) UpdateRole(updateOpt role.UpdateRoleOption) error {
	err := r.roleData.UpdateRole(r.ds.ORM(), entity.Role{
		Model: gorm.Model{ID: updateOpt.Id},
		Name:  updateOpt.Name,
	})
	if err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (r RoleService) RemoveRole(roleId uint) error {
	if err := r.roleData.RemoveRole(r.ds.ORM(), roleId); err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (r RoleService) GetRolePerms(roleId uint) ([]role.PermGroup, error) {
	perms, err := r.roleData.GetRolePerms(r.ds.ORM(), roleId)
	if err != nil {
		return []role.PermGroup{}, errs.DataBaseErr(err)
	}
	return makePermGroup(perms), nil
}

func (r RoleService) GrantRolePerms(roleId uint, tag string, newPermIds []uint) error {

	orm := r.ds.ORM()

	allPermIds, err := r.permData.GetIds(orm, tag)
	if err != nil {
		return errs.DataBaseErr(err)
	}

	// there has unexpected permId in params if is not subset
	if !collection.IsSubset(newPermIds, allPermIds) {
		return errs.NewI18nError("perm.invalidList")
	}

	// find the complement set
	deletedSet := collection.ComplementSet(newPermIds, allPermIds)

	// start transaction
	tx := orm.Begin()

	// insert new permIds
	if err = r.roleData.InsertRolePerms(tx, roleId, newPermIds); err != nil {
		tx.Rollback()
		return errs.DataBaseErr(err)
	}

	// remove the complement set
	if err = r.roleData.RemoveRolePerms(tx, roleId, deletedSet); err != nil {
		tx.Rollback()
		return errs.DataBaseErr(err)
	}

	// commit
	if tx.Commit().Error != nil {
		return errs.DataBaseErr(tx.Commit().Error)
	}

	return nil
}
