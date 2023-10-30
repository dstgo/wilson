package entity

import (
	"github.com/dstgo/wilson/internal/types/dict"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Dict struct {
	Id uint `gorm:"primaryKey;"`

	Name string `gorm:"type:varchar(100);"`
	Code string `gorm:"type:varchar(100);"`

	Data []DictData

	CreatedAt uint64 `gorm:"autoCreateTime:nano;"`
	UpdatedAt uint64 `gorm:"autoUpdateTime:nano;"`
}

func (d Dict) TableName() string {
	return "dict"
}

func (d Dict) TableComment() string {
	return "dict info table"
}

type DictData struct {
	Id uint `gorm:"primaryKey;"`

	Name  string `gorm:"type:varchar(100);comment:data name;"`
	Key   string `gorm:"type:varchar(100);uniqueIndex;comment:data key;"`
	Value string `gorm:"type:varchar(255);comment:data value;"`
	Type  uint8  `gorm:"comment:data type"`
	Order int    `gorm:"comment:order flag;"`

	CreatedAt uint64         `gorm:"autoCreateTime:nano;"`
	UpdatedAt uint64         `gorm:"autoUpdateTime:nano;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:disable status;"`

	DictId uint `gorm:"comment:dict type reference key;"`
}

func (d DictData) TableName() string {
	return "dict_data"
}

func (d DictData) TableComment() string {
	return "dict data info table"
}

func (d DictData) DataValue() (any, error) {
	switch d.Type {
	case dict.StringType:
		return d.Value, nil
	case dict.Int64Type:
		return cast.ToInt64E(d.Value)
	case dict.Float64Type:
		return cast.ToFloat64E(d.Value)
	case dict.BoolType:
		return cast.ToBoolE(d.Value)
	default:
		return nil, dict.ErrInvalidDicType
	}
}
