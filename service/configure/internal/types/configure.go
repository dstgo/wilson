package types

type ListConfigureRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"pageSize"`
	ServerId uint32 `json:"serverId"`
	EnvId    uint32 `json:"envId"`
}

type CompareConfigureRequest struct {
	ServerId uint32 `json:"serverId"`
	EnvId    uint32 `json:"envId"`
}

type CompareConfigureReply struct {
	Type string `json:"type"`
	Key  string `json:"key"`
	Old  string `json:"old"`
	Cur  string `json:"cur"`
}

type WatcherConfigRequest struct {
	Token  string
	Server string
}

type WatcherConfigureReply struct {
	Format  string `json:"format"`
	Content string `json:"content"`
}

type WatcherConfigReplyFunc func(*WatcherConfigureReply) error
