package field

import (
	"google.golang.org/protobuf/types/known/structpb"
)

type stringType struct {
}

func (t stringType) Name() string {
	return "text"
}

func (t stringType) Validate(_ *structpb.Value) bool {
	return true
}

func (t stringType) ToString(in *structpb.Value) string {
	return in.GetStringValue()
}

func (t stringType) ToValue(in string) *structpb.Value {
	return structpb.NewStringValue(in)
}
