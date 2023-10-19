package role

import (
	"errors"
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/api/role"
	"github.com/dstgo/wilson/internal/types/errs"
	"gorm.io/gorm"
)

func NewPermService(perm PermData, role RoleData, ds *data.DataSource) PermService {
	return PermService{permData: perm, roleData: role, ds: ds}
}

type PermService struct {
	permData PermData
	roleData RoleData
	ds       *data.DataSource
}

func (p PermService) CreatePerm(createOpt role.CreatePermOption) error {

	record, err := p.permData.FindPerm(p.ds.ORM(), createOpt.Object, createOpt.Action, createOpt.Group, createOpt.Tag)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.DataBaseErr(err)
	} else if record.ID > 0 {
		return errs.NewI18nError("perm.conflict")
	}

	err = p.permData.CreatePerm(p.ds.ORM(), entity.Permission{
		Name:   createOpt.Name,
		Object: createOpt.Object,
		Action: createOpt.Action,
		Group:  createOpt.Group,
		Tag:    createOpt.Tag,
	})

	if err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (p PermService) UpdatePerm(updateOpt role.UpdatePermOption) error {
	err := p.permData.UpdateRole(p.ds.ORM(), entity.Permission{
		Model: gorm.Model{ID: updateOpt.Id},
		Name:  updateOpt.Name,
	})

	if err != nil {
		return errs.DataBaseErr(err)
	}

	return nil
}

func (p PermService) RemovePerm(permId uint) error {
	err := p.permData.RemoveRole(p.ds.ORM(), permId)
	if err != nil {
		return errs.DataBaseErr(err)
	}
	return nil
}

func (p PermService) GetPermList(pageOpt role.PageOption) ([]role.PermInfo, error) {
	var permList []role.PermInfo
	perms, err := p.permData.GetPagePermListBy(p.ds.ORM(), pageOpt)
	if err != nil {
		return []role.PermInfo{}, errs.DataBaseErr(err)
	}
	if err := cp.Copy(&perms, &permList); err != nil {
		return []role.PermInfo{}, errs.ProgramErr(err)
	}
	return permList, nil
}

func (p PermService) GetPermById(permId uint) (role.PermInfo, error) {
	permInfo, err := p.permData.GetById(p.ds.ORM(), permId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return role.PermInfo{}, errs.NewI18nError("perm.notfound")
	} else if err != nil {
		return role.PermInfo{}, errs.DataBaseErr(err)
	}
	return role.PermInfo{Name: permInfo.Name, Object: permInfo.Object, Action: permInfo.Object}, nil
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
