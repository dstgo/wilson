package vax

import (
	"github.com/dstgo/wilson/app/pkg/errorx"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/gin-gonic/gin"
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
	return errorx.Join(ctx.ShouldBindJSON(val), val.Validate(httpx.GetFirstAcceptLanguage(ctx))).Err()
}

type XmlBinding struct {
}

func (x XmlBinding) Bind(ctx *gin.Context, val Validatable) error {
	return errorx.Join(ctx.ShouldBindXML(val), val.Validate(httpx.GetFirstAcceptLanguage(ctx))).Err()
}

type QueryBinding struct {
}

func (q QueryBinding) Bind(ctx *gin.Context, val Validatable) error {
	return errorx.Join(ctx.ShouldBindQuery(val), val.Validate(httpx.GetFirstAcceptLanguage(ctx))).Err()
}

type YamlBinding struct {
}

func (y YamlBinding) Bind(ctx *gin.Context, val Validatable) error {
	return errorx.Join(ctx.ShouldBindYAML(val), val.Validate(httpx.GetFirstAcceptLanguage(ctx))).Err()
}

type TomlBinding struct {
}

func (y TomlBinding) Bind(ctx *gin.Context, val Validatable) error {
	return errorx.Join(ctx.ShouldBindTOML(val), val.Validate(httpx.GetFirstAcceptLanguage(ctx))).Err()
}

type HeaderBinding struct {
}

func (y HeaderBinding) Bind(ctx *gin.Context, val Validatable) error {
	return errorx.Join(ctx.ShouldBindHeader(val), val.Validate(httpx.GetFirstAcceptLanguage(ctx))).Err()
}

type UriBinding struct {
}

func (y UriBinding) Bind(ctx *gin.Context, val Validatable) error {
	return errorx.Join(ctx.ShouldBindUri(val), val.Validate(httpx.GetFirstAcceptLanguage(ctx))).Err()
}
