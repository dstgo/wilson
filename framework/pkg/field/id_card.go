package field

import (
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/dstgo/wilson/framework/pkg/valx"
)

type idCardType struct {
}

func (it idCardType) Name() string {
	return "id card"
}

func (it idCardType) Validate(in *structpb.Value) bool {
	return valx.IsIDCard(in.GetStringValue())
}

func (it idCardType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (it idCardType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
