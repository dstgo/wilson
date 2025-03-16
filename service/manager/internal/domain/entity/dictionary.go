package entity

import "github.com/dstgo/wilson/framework/kratosx/types"

type Dictionary struct {
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Description *string `json:"description" gorm:"column:description"`
	types.BaseModel
}

type DictionaryValue struct {
	DictionaryId uint32             `json:"dictionaryId" gorm:"column:dictionary_id"`
	Label        string             `json:"label" gorm:"column:label"`
	Value        string             `json:"value" gorm:"column:value"`
	Status       *bool              `json:"status" gorm:"column:status"`
	Weight       *int32             `json:"weight" gorm:"column:weight"`
	Type         *string            `json:"type" gorm:"column:type"`
	Extra        *string            `json:"extra" gorm:"column:extra"`
	Description  *string            `json:"description" gorm:"column:description"`
	Dictionary   *Dictionary        `json:"dictionary"`
	Children     []*DictionaryValue `json:"children" gorm:"-"`
	types.BaseModel
}

// ID 获取ID
func (m *DictionaryValue) ID() uint32 {
	return m.Id
}

// AppendChildren 添加子节点
func (m *DictionaryValue) AppendChildren(child *DictionaryValue) {
	m.Children = append(m.Children, child)
}

// ChildrenNode 获取子节点
func (m *DictionaryValue) ChildrenNode() []*DictionaryValue {
	return append([]*DictionaryValue{}, m.Children...)
}
