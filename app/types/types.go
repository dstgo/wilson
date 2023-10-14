package types

var (
	NopObj = NopType{}
)

type (
	H       = map[string]any
	S       = []any
	NopType = struct{}
	Strings = []string
)
