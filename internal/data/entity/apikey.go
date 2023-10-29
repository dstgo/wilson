package entity

type APIKey struct {
	Key    string `json:"key"`
	UserId string `json:"userId"`

	Name  string `json:"name"`
	Perms []uint `json:"perms"`

	CreatedAt uint64 `json:"createdAt"`
	ExpiredAt uint64 `json:"expiredAt"`
}
