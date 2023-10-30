package system

import (
	"github.com/dstgo/wilson/internal/types/code"
	"github.com/dstgo/wilson/internal/types/errs"
	"net/http"
)

var (
	ErrBadParams        = errs.NewI18nError("err.badparams").FallBack("invalid params").Status(http.StatusBadRequest).Code(code.BadRequest)
	ErrForbidden        = errs.NewI18nError("err.forbidden").FallBack("forbidden").Status(http.StatusForbidden).Code(code.Forbidden)
	ErrUnAthorzied      = errs.NewI18nError("err.unauthorized").FallBack("unauthorized").Status(http.StatusUnauthorized).Code(code.UnAuthorized)
	ErrResourceNotFound = errs.NewI18nError("err.notfound").FallBack("resource not found").Status(http.StatusNotFound).Code(code.ResourceNotFound)
)

var (

	// ErrDatabase
	// This error is used to describe database operations.
	// These operations may not have any exceptions themselves, but they have not caused any impact, so they are invalid operations.
	ErrDatabase = errs.NewI18nError("err.database").FallBack("database error").Status(http.StatusInternalServerError).Code(code.DatabaseError)

	ErrFileSystem = errs.NewI18nError("err.filesystem").FallBack("file system error").Status(http.StatusInternalServerError).Code(code.FilesystemError)

	ErrNetwork = errs.NewI18nError("err.network").FallBack("network error").Status(http.StatusInternalServerError).Code(code.NetworkError)

	ErrProgram = errs.NewI18nError("err.program").FallBack("program error").Status(http.StatusInternalServerError).Code(code.ProgramError)

	ErrUnknown = errs.NewI18nError("err.unknown").FallBack("unknown error").Status(http.StatusInternalServerError).Code(code.UnknownError)
)
