package types

type ListEnvRequest struct {
	Keyword *string  `json:"keyword"`
	Name    *string  `json:"name"`
	Status  *bool    `json:"status"`
	Ids     []uint32 `json:"ids"`
}
