package auth

import (
	"github.com/dstgo/wilson/internal/types/code"
	"github.com/dstgo/wilson/internal/types/errs"
	"net/http"
)

var (
	ErrTokenInvalid        = errs.NewI18nError("jwt.invalid").FallBack("invalid token").Code(code.UnAuthorized)
	ErrTokenParsedFailed   = errs.NewI18nError("jwt.parsed.failed").FallBack("token parsed failed").Code(code.Forbidden)
	ErrTokenExpired        = errs.NewI18nError("jwt.expired").FallBack("token expired").Code(code.UnAuthorized)
	ErrTokenNeedRefresh    = errs.NewI18nError("jwt.refresh").FallBack("token need to be refresh").Code(code.TokenNeedFresh)
	ErrWrongPassword       = errs.NewI18nError("user.wrongPassword").FallBack("wrong password")
	ErrTokenIssueFailed    = errs.NewI18nError("jwt.issue.failed").FallBack("token issue failed").Status(http.StatusInternalServerError)
	ErrRedundantExpiration = errs.NewI18nError("jwt.redundant").FallBack("redundant expiration").Code(code.RedundantRefresh)
)

var (
	ErrInvalidKey          = errs.NewI18nError("key.invalid").FallBack("invalid key")
	ErrKeyNoPerm           = errs.NewI18nError("key.noperm").FallBack("key has no permission")
	ErrInvalidKeyExpration = errs.NewI18nError("key.nopexp").FallBack("invalid expiration time")
)
