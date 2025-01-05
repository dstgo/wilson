package types

type ListResourceRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Keyword  *string `json:"keyword"`
	Tag      *string `json:"tag"`
	Private  *bool   `json:"private"`
	ServerId *uint32 `json:"serverId"`
}

type ListResourceValueRequest struct {
	ServerId   *uint32 `json:"serverId"`
	ResourceId *uint32 `json:"resourceId"`
	EnvId      *uint32 `json:"envId"`
}
