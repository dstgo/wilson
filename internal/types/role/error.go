package role

import (
	"github.com/dstgo/wilson/internal/types/code"
	"github.com/dstgo/wilson/internal/types/errs"
)

var (
	ErrRoleNotFound = errs.NewI18nError("role.notfound").FallBack("role not found").Code(code.ResourceNotFound)
	ErrRoleConflict = errs.NewI18nError("role.conflict").FallBack("role already exists")
	ErrInvalidRoles = errs.NewI18nError("role.invalidList").FallBack("invalid roles")
)

var (
	ErrPermNotFound = errs.NewI18nError("perm.notfound").FallBack("permission not found").Code(code.ResourceNotFound)
	ErrPermConflict = errs.NewI18nError("perm.conflict").FallBack("permission already exists")
	ErrPermNoAccess = errs.NewI18nError("perm.noPerm").FallBack("no enough permission to access")
	ErrInvalidPerm  = errs.NewI18nError("perm.invalidPerm").FallBack("invalid permissions")
)
