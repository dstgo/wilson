package role

import "github.com/gin-gonic/gin"

type RoleHandler struct {
}

// AddRole
//
//	@Summary		NewRole
//	@Description	create a new role
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Router			/role/create [POST]
func (r RoleHandler) AddRole(ctx *gin.Context) {

}

// UpdateRole
//
//	@Summary		UpdateRole
//	@Description	update the specified role info
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Router			/role/update [POST]
func (r RoleHandler) UpdateRole(ctx *gin.Context) {

}

// RemoveRole
//
//	@Summary		RemoveRole
//	@Description	remove a role
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Router			/role/remove [DELETE]
func (r RoleHandler) RemoveRole(ctx *gin.Context) {

}

// AddPermission
//
//	@Summary		NewPermission
//	@Description	create a new permission
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Router			/role/perm/create [POST]
func (r RoleHandler) AddPermission(ctx *gin.Context) {

}

// UpdatePermission
//
//	@Summary		UpdatePermission
//	@Description	update the specified permission info
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Router			/role/perm/update [POST]
func (r RoleHandler) UpdatePermission(ctx *gin.Context) {

}

// RemovePermission
//
//	@Summary		RemovePermission
//	@Description	remove the specified permission
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Router			/role/perm/remove [DELETE]
func (r RoleHandler) RemovePermission(ctx *gin.Context) {

}

// GrantPermForRole
//
//	@Summary		GrantPermForRole
//	@Description	grant permissions for the specified role
//	@Tags			role
//	@Accept			json
//	@Produce		json
//	@Router			/role/grant [POST]
func (r RoleHandler) GrantPermForRole(ctx *gin.Context) {

}
