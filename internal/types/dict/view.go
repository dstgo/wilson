package dict

const (
	StringType uint8 = iota
	Int64Type
	Float64Type
	BoolType
)

type DictDataInfo struct {
	Label string `json:"label"`
	Key   string `json:"key"`
	Value any    `json:"value"`
	Type  uint8  `json:"type"`
}

type DictDetail struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CreatedAt uint64 `json:"createdAt"`
	UpdatedAt uint64 `json:"updatedAt"`
}

type DictDataDetail struct {
	Id     uint   `json:"id"`
	Label  string `json:"label"`
	Key    string `json:"key"`
	Value  any    `json:"value"`
	Type   uint8  `json:"type"`
	Order  int    `json:"order"`
	Enable bool   `json:"enable"`

	CreatedAt uint64 `json:"createdAt"`
	UpdatedAt uint64 `json:"updatedAt"`
}
