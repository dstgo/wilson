package types

type ListTemplateRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"pageSize"`
	ServerId uint32 `json:"serverId"`
}

type CompareTemplateRequest struct {
	Id      uint32 `json:"id"`
	Format  string `json:"format"`
	Content string `json:"content"`
}

type CompareTemplateReply struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	Old  string `json:"old"`
	Cur  string `json:"cur"`
}

type PreviewTemplateRequest struct {
	EnvId    uint32 `json:"envId"`
	ServerId uint32 `json:"serverId"`
	Format   string `json:"format"`
	Content  string `json:"content"`
}

type PreviewTemplateReply struct {
	Format  string `json:"format"`
	Content string `json:"content"`
}

type PreviewCurrentTemplateRequest struct {
	EnvId    uint32 `json:"envId"`
	ServerId uint32 `json:"serverId"`
}

type ParseByContentRequest struct {
	EnvId    uint32 `json:"env_id"`
	ServerId uint32 `json:"server_id"`
	Format   string `json:"format"`
	Content  string `json:"content"`
}

type TemplateValue struct {
	Value   string
	Exclude bool
}
