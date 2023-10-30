package dict

import (
	"github.com/dstgo/wilson/internal/types/helper"
	"github.com/dstgo/wilson/pkg/vax"
)

type DictPageOption struct {
	helper.PageOption
	Search string `json:"search" uri:"search" form:"search"`
}

func (p DictPageOption) Validate(lang string) error {
	return vax.Struct(&p, lang,
		vax.Field(&p.PageOption),
	)
}

type DictDataPageOption struct {
	helper.PageOption
	Code   string `json:"code" uri:"code" form:"code" label:"dict.code"`
	Search string `json:"search" uri:"search" form:"search"`
}

func (d DictDataPageOption) Validate(lang string) error {
	return vax.Struct(&d, lang,
		vax.Field(&d.PageOption),
		vax.Field(&d.Code, vax.Required),
	)
}

type CodeOption struct {
	Code string `json:"code" uri:"code" form:"code" label:"dict.code"`
}

func (d CodeOption) Validate(lang string) error {
	return vax.Struct(&d, lang,
		vax.Field(&d.Code, helper.RequiredRules(RuleDictCode)...),
	)
}

type DictSaveOption struct {
	Name string `json:"name" label:"dict.name"`
	Code string `json:"code" label:"dict.code"`
}

func (d DictSaveOption) Validate(lang string) error {
	return vax.Struct(&d, lang,
		vax.Field(&d.Name, helper.RequiredRules(RuleDictName)...),
		vax.Field(&d.Code, helper.RequiredRules(RuleDictCode)...),
	)
}

type DictUpdateOption struct {
	Id uint `json:"id"`
	DictSaveOption
}

func (d DictUpdateOption) Validate(lang string) error {
	return vax.Struct(&d, lang,
		vax.Field(&d.Id, vax.Required),
		vax.Field(&d.DictSaveOption),
	)
}

type CodeKeyOption struct {
	CodeOption
	Key string `json:"key" uri:"key" form:"key" label:"dict.key"`
}

func (d CodeKeyOption) Validate(lang string) error {
	return vax.Struct(&d, lang,
		vax.Field(&d.CodeOption),
		vax.Field(&d.Key, helper.RequiredRules(RuleDictKey)...),
	)
}

type DictDataSaveOption struct {
	DictId uint   `json:"dictId"`
	Name   string `json:"name" label:"dict.name"`
	Key    string `json:"key" label:"dict.key"`
	Value  string `json:"value" label:"dict.value"`
	Type   uint8  `json:"type" label:"dict.type"`
	Order  int    `json:"order" label:"dict.order"`
}

func (d DictDataSaveOption) Validate(lang string) error {
	return vax.Struct(&d, lang,
		vax.Field(&d.DictId, vax.Required),
		vax.Field(&d.Name, helper.RequiredRules(RuleDictName)...),
		vax.Field(&d.Key, helper.RequiredRules(RuleDictKey)...),
		vax.Field(&d.Value, helper.RequiredRules(RuleDictValue)...),
		vax.Field(&d.Type, RuleDictDataType...),
	)
}

type DictDataUpdateOption struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	Type   uint8  `json:"type"`
	Order  int    `json:"order"`
	Enable bool
}

func (d DictDataUpdateOption) Validate(lang string) error {
	return vax.Struct(&d, lang,
		vax.Field(&d.Id, vax.Required),
		vax.Field(&d.Name, helper.RequiredRules(RuleDictName)...),
		vax.Field(&d.Key, helper.RequiredRules(RuleDictKey)...),
		vax.Field(&d.Value, helper.RequiredRules(RuleDictValue)...),
		vax.Field(&d.Type, RuleDictDataType...),
	)
}
