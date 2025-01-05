package field

import (
	"strconv"

	"google.golang.org/protobuf/types/known/structpb"

	"github.com/dstgo/wilson/framework/pkg/valx"
)

type numberType struct {
}

func (nt numberType) Name() string {
	return "number"
}

func (nt numberType) Validate(in *structpb.Value) bool {
	return valx.IsNumber(in.GetStringValue())
}

func (nt numberType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (nt numberType) ToValue(in string) *structpb.Value {
	value, _ := strconv.ParseFloat(in, 64)
	return structpb.NewNumberValue(value)
}
