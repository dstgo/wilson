package vax

import (
	"errors"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/gin-gonic/gin"
	"io"
)

var (
	BindingJson   = JsonBinding{}
	BindingXml    = XmlBinding{}
	BingQuery     = QueryBinding{}
	BindingYaml   = YamlBinding{}
	BindingToml   = TomlBinding{}
	BindingHeader = HeaderBinding{}
	BindingUri    = UriBinding{}
)

type Binding interface {
	Bind(ctx *gin.Context, val Validatable) error
}

type JsonBinding struct {
}

func (j JsonBinding) Bind(ctx *gin.Context, val Validatable) error {
	if err := ctx.ShouldBindJSON(val); !errors.Is(err, io.EOF) && err != nil {
		return err
	}
	return val.Validate(httpx.GetFirstAcceptLanguage(ctx))
}

type XmlBinding struct {
}

func (x XmlBinding) Bind(ctx *gin.Context, val Validatable) error {
	if err := ctx.ShouldBindXML(val); !errors.Is(err, io.EOF) && err != nil {
		return err
	}
	return val.Validate(httpx.GetFirstAcceptLanguage(ctx))
}

type QueryBinding struct {
}

func (q QueryBinding) Bind(ctx *gin.Context, val Validatable) error {
	if err := ctx.ShouldBindQuery(val); !errors.Is(err, io.EOF) && err != nil {
		return err
	}
	return val.Validate(httpx.GetFirstAcceptLanguage(ctx))
}

type YamlBinding struct {
}

func (y YamlBinding) Bind(ctx *gin.Context, val Validatable) error {
	if err := ctx.ShouldBindYAML(val); !errors.Is(err, io.EOF) && err != nil {
		return err
	}
	return val.Validate(httpx.GetFirstAcceptLanguage(ctx))
}

type TomlBinding struct {
}

func (y TomlBinding) Bind(ctx *gin.Context, val Validatable) error {
	if err := ctx.ShouldBindTOML(val); !errors.Is(err, io.EOF) && err != nil {
		return err
	}
	return val.Validate(httpx.GetFirstAcceptLanguage(ctx))
}

type HeaderBinding struct {
}

func (y HeaderBinding) Bind(ctx *gin.Context, val Validatable) error {
	if err := ctx.ShouldBindHeader(val); !errors.Is(err, io.EOF) && err != nil {
		return err
	}
	return val.Validate(httpx.GetFirstAcceptLanguage(ctx))
}

type UriBinding struct {
}

func (y UriBinding) Bind(ctx *gin.Context, val Validatable) error {
	if err := ctx.ShouldBindUri(val); !errors.Is(err, io.EOF) && err != nil {
		return err
	}
	return val.Validate(httpx.GetFirstAcceptLanguage(ctx))
}
