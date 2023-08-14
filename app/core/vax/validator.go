package vax

import (
	"github.com/gin-gonic/gin"
)

type BindPair struct {
	B Binding
	V Validatable
}

func ShouldBindAll(ctx *gin.Context, pairs ...BindPair) error {
	for _, pair := range pairs {
		err := ShouldBind(ctx, pair)
		if err != nil {
			return err
		}
	}
	return nil
}

func ShouldBind(ctx *gin.Context, pair BindPair) error {
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
