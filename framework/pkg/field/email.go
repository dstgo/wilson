package field

import (
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/dstgo/wilson/framework/pkg/valx"
)

type emailType struct {
}

func (et emailType) Name() string {
	return "email"
}

func (et emailType) Validate(in *structpb.Value) bool {
	return valx.IsEmail(in.GetStringValue())
}

func (et emailType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (et emailType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
