package role

import "github.com/dstgo/wilson/internal/data/entity"

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

type RoleInfo struct {
	// role id
	Id uint `json:"id" example:"1"`
	// role name
	Name string `json:"name" example:"admin"`
	// role code, must be alpha numeric
	Code string `json:"code" example:"ADMIN"`
}

type PermGroup struct {
	// group name
	Group string     `json:"group" example:"user group"`
	Perms []PermInfo `json:"perms"`
}

type PermInfo struct {
	// permission id
	Id uint `json:"id" example:"1"`
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
		Id:     perm.Id,
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
		Id:     perm.Id,
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

func MakeRolePerms(roleId uint, permIds []uint) []entity.RolePermission {
	rolePermList := make([]entity.RolePermission, 0, len(permIds))
	for _, permId := range permIds {
		rolePermList = append(rolePermList, entity.RolePermission{RoleId: roleId, PermissionId: permId})
	}
	return rolePermList
}
