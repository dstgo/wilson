package dict

import (
	"errors"
	"github.com/dstgo/wilson/internal/types/errs"
)

var (
	ErrInvalidDicType = errors.New("invalid dict data type")
)

var (
	ErrDictNotFound        = errs.NewI18nError("dict.notfound").FallBack("dict not found")
	ErrDictCodeConflict    = errs.NewI18nError("dict.codeConflict").FallBack("dict code conflict")
	ErrDictDataNotFound    = errs.NewI18nError("dict.data.notfound").FallBack("dict data not found")
	ErrDictDataKeyConflict = errs.NewI18nError("dict.data.keyConflict").FallBack("dict data key conflict")

	ErrInvalidDictDatType = errs.NewI18nError("dict.data.invalidType").FallBack("invalid dict data type")
)
