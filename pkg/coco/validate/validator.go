package validate

import (
	"github.com/gin-gonic/gin"
)

type Validator interface {
	Validate() error
}

type BindPair struct {
	B Binding
	V Validator
}

func ShouldBindAll(ctx *gin.Context, pairs ...BindPair) error {
	for _, pair := range pairs {
		err := ShouldBind(ctx, pair)
		if err != nil {
			return nil
		}
	}
	return nil
}

func ShouldBind(ctx *gin.Context, pair BindPair) error {
	err := pair.B.Bind(ctx, pair.V)
	if err != nil {
		return nil
	}
	return nil
}

func Pair(b Binding, val Validator) BindPair {
	return BindPair{B: b, V: val}
}

func Json(val Validator) BindPair {
	return Pair(JSON, val)
}

func Query(val Validator) BindPair {
	return Pair(QUERY, val)
}

func Xml(val Validator) BindPair {
	return Pair(XML, val)
}

func Yaml(val Validator) BindPair {
	return Pair(YAML, val)
}

func Toml(val Validator) BindPair {
	return Pair(TOML, val)
}

func Header(val Validator) BindPair {
	return Pair(HEADER, val)
}

func Uri(val Validator) BindPair {
	return Pair(URI, val)
}
