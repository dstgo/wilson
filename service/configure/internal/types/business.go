package types

type ListBusinessRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	ServerId uint32  `json:"serverId"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Keyword  *string `json:"keyword"`
}

type ListBusinessValueRequest struct {
	ServerId   *uint32 `json:"serverId"`
	BusinessId *uint32 `json:"businessId"`
	EnvId      *uint32 `json:"envId"`
}
