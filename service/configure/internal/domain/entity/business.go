package entity

import (
	"strconv"

	json "github.com/json-iterator/go"

	"github.com/dstgo/wilson/framework/kratosx/types"

	"github.com/dstgo/wilson/api/gen/errors"
)

type Business struct {
	ServerId    uint32  `json:"serverId" gorm:"column:server_id"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Type        string  `json:"type" gorm:"column:type"`
	Description *string `json:"description" gorm:"column:description"`
	types.BaseModel
}

type BusinessValue struct {
	EnvId      uint32    `json:"envId" gorm:"column:env_id"`
	BusinessId uint32    `json:"businessId" gorm:"column:business_id"`
	Value      string    `json:"value" gorm:"column:value"`
	Business   *Business `json:"business"`
	types.BaseModel
}

func (bv *BusinessValue) MarshalValue(tp string) (string, error) {
	var (
		isAllow  bool
		value    = bv.Value
		newValue = value
	)
	switch tp {
	case "int":
		_, err := strconv.Atoi(value)
		isAllow = err == nil
	case "float":
		_, err := strconv.ParseFloat(value, 64)
		isAllow = err == nil
	case "string":
		isAllow = true
	case "bool":
		isAllow = value == "true" || value == "false"
	case "object":
		var m any
		isAllow = json.Unmarshal([]byte(value), &m) == nil
		content, _ := json.Marshal(m)
		newValue = string(content)
	default:
		isAllow = false
	}
	if !isAllow {
		return "", errors.BusinessValueTypeError()
	}
	return newValue, nil
}
