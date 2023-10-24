package user

import (
	"github.com/dstgo/wilson/internal/types/code"
	"github.com/dstgo/wilson/internal/types/errs"
)

var (
	ErrUserNotFound        = errs.NewI18nError("user.notfound").FallBack("user not found").Code(code.ResourceNotFound)
	ErrUsernameAlreadyUsed = errs.NewI18nError("user.nameUsed").FallBack("username already used")
	ErrEmailAlreadyUsed    = errs.NewI18nError("user.emailUsed").FallBack("email already used")
	ErrUserAlreadyExists   = errs.NewI18nError("user.alreadyExist").FallBack("user already exists")
)
