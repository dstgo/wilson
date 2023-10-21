package role

var (
	// AdminRole app static admin role,
	AdminRole = RoleInfo{
		Name: "AppAdmin",
		Code: "1024",
	}
)

type RoleInfo struct {
	// role id
	ID uint `json:"id" example:"1"`
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
