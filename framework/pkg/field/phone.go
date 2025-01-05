package field

import (
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/dstgo/wilson/framework/pkg/valx"
)

type phoneType struct {
}

func (pt phoneType) Name() string {
	return "phone"
}

func (pt phoneType) Validate(in *structpb.Value) bool {
	return valx.IsPhone(in.GetStringValue())
}

func (pt phoneType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (pt phoneType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
