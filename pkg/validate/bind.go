package validate

import (
	"github.com/dstgo/wilson/app/pkg/errorx"
	"github.com/gin-gonic/gin"
)

var (
	JSON   = JsonBinding{}
	XML    = XmlBinding{}
	QUERY  = QueryBinding{}
	YAML   = YamlBinding{}
	TOML   = TomlBinding{}
	HEADER = HeaderBinding{}
	URI    = UriBinding{}
)

type Binding interface {
	Bind(ctx *gin.Context, val Validator) error
}

type JsonBinding struct {
}

func (j JsonBinding) Bind(ctx *gin.Context, val Validator) error {
	return errorx.Join(ctx.ShouldBindJSON(val), val.Validate()).Err()
}

type XmlBinding struct {
}

func (x XmlBinding) Bind(ctx *gin.Context, val Validator) error {
	return errorx.Join(ctx.ShouldBindXML(val), val.Validate()).Err()
}

type QueryBinding struct {
}

func (q QueryBinding) Bind(ctx *gin.Context, val Validator) error {
	return errorx.Join(ctx.ShouldBindQuery(val), val.Validate()).Err()
}

type YamlBinding struct {
}

func (y YamlBinding) Bind(ctx *gin.Context, val Validator) error {
	return errorx.Join(ctx.ShouldBindYAML(val), val.Validate()).Err()
}

type TomlBinding struct {
}

func (y TomlBinding) Bind(ctx *gin.Context, val Validator) error {
	return errorx.Join(ctx.ShouldBindTOML(val), val.Validate()).Err()
}

type HeaderBinding struct {
}

func (y HeaderBinding) Bind(ctx *gin.Context, val Validator) error {
	return errorx.Join(ctx.ShouldBindHeader(val), val.Validate()).Err()
}

type UriBinding struct {
}

func (y UriBinding) Bind(ctx *gin.Context, val Validator) error {
	return errorx.Join(ctx.ShouldBindUri(val), val.Validate()).Err()
}
