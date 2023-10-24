package bind

import (
	"github.com/dstgo/wilson/pkg/vax"
	"github.com/gin-gonic/gin"
)

// Handler will be called when bind error occurs
type Handler func(ctx *gin.Context, bindErr error)

var HandlerChain []Handler

type BindPair struct {
	B Binding
	V vax.Validatable
}

func Binds(ctx *gin.Context, pairs ...BindPair) error {
	var bindErr error
	for _, pair := range pairs {
		if err := Bind(ctx, pair); err != nil {
			bindErr = err
			break
		}
	}

	if len(HandlerChain) > 0 && bindErr != nil {
		for _, handle := range HandlerChain {
			handle(ctx, bindErr)
		}
	}
	return bindErr
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
