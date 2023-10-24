package auth

import "github.com/dstgo/wilson/internal/types/errs"

var (
	ErrJwtParsedFailed = errs.NewI18nError("jwt.parsed.failed").FallBack("jwt parsing failed")
	ErrJwtExpired      = errs.NewI18nError("jwt.expired").FallBack("jwt expired")
	ErrWrongPassword   = errs.NewI18nError("user.wrongPassword").FallBack("wrong password")
)
