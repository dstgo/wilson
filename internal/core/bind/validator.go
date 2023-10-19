package bind

import (
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/gin-gonic/gin"
)

type BindPair struct {
	B Binding
	V vax.Validatable
}

func BindAll(ctx *gin.Context, pairs ...BindPair) error {
	for _, pair := range pairs {
		err := Bind(ctx, pair)
		if err != nil {
			return err
		}
	}
	return nil
}

func BindAndResp(ctx *gin.Context, pairs ...BindPair) error {
	err := BindAll(ctx, pairs...)
	if err != nil {
		switch err.(type) {
		// validate internal occur error
		case vax.InternalError:
			resp.InternalFailed(ctx).MsgI18n("err.program").Error(err).Send()
		default:
			resp.Fail(ctx).MsgI18n("err.badparams").Error(err).Send()
		}
	}
	return err
}

func Bind(ctx *gin.Context, pair BindPair) error {
	err := pair.B.Bind(ctx, pair.V)
	if err != nil {
		return err
	}
	return nil
}

func Pair(b Binding, val vax.Validatable) BindPair {
	return BindPair{B: b, V: val}
}

func Json(val vax.Validatable) BindPair {
	return Pair(BindingJson, val)
}

func Query(val vax.Validatable) BindPair {
	return Pair(BingQuery, val)
}

func Xml(val vax.Validatable) BindPair {
	return Pair(BindingXml, val)
}

func Yaml(val vax.Validatable) BindPair {
	return Pair(BindingYaml, val)
}

func Toml(val vax.Validatable) BindPair {
	return Pair(BindingToml, val)
}

func Header(val vax.Validatable) BindPair {
	return Pair(BindingHeader, val)
}

func Uri(val vax.Validatable) BindPair {
	return Pair(BindingUri, val)
}
