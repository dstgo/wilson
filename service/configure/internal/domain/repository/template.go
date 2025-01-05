package repository

import (
	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Template interface {
	// GetTemplateByVersion 获取指定版本的模板
	GetTemplateByVersion(ctx kratosx.Context, version string) (*entity.Template, error)

	// GetTemplate 获取指定ID的模板
	GetTemplate(ctx kratosx.Context, id uint32) (*entity.Template, error)

	// CurrentTemplate 获取当前正在使用的模板
	CurrentTemplate(ctx kratosx.Context, srvId uint32) (*entity.Template, error)

	// ListTemplate 获取模板列表
	ListTemplate(ctx kratosx.Context, options *types.ListTemplateRequest) ([]*entity.Template, uint32, error)

	// CreateTemplate 创建模板
	CreateTemplate(ctx kratosx.Context, c *entity.Template) (uint32, error)

	// UpdateTemplate 更新模板
	UpdateTemplate(ctx kratosx.Context, c *entity.Template) error

	// UseTemplate 使用指定模板
	UseTemplate(ctx kratosx.Context, srvId, tpId uint32) error

	// DeleteTemplate 删除指定模板
	DeleteTemplate(ctx kratosx.Context, id uint32) error
}
