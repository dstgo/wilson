package vax

import (
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/types/code"
	"github.com/gin-gonic/gin"
)

type BindPair struct {
	B Binding
	V Validatable
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
		resp.Fail(ctx).Code(code.BadRequest).MsgI18n("error.badparams").Error(err).Send()
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

func Pair(b Binding, val Validatable) BindPair {
	return BindPair{B: b, V: val}
}

func Json(val Validatable) BindPair {
	return Pair(BindingJson, val)
}

func Query(val Validatable) BindPair {
	return Pair(BingQuery, val)
}

func Xml(val Validatable) BindPair {
	return Pair(BindingXml, val)
}

func Yaml(val Validatable) BindPair {
	return Pair(BindingYaml, val)
}

func Toml(val Validatable) BindPair {
	return Pair(BindingToml, val)
}

func Header(val Validatable) BindPair {
	return Pair(BindingHeader, val)
}

func Uri(val Validatable) BindPair {
	return Pair(BindingUri, val)
}
