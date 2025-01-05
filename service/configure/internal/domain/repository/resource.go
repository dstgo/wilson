package repository

import (
	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Resource interface {
	// GetResource 获取指定的资源配置信息
	GetResource(ctx kratosx.Context, id uint32) (*entity.Resource, error)

	// ListResource 获取资源配置信息列表
	ListResource(ctx kratosx.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error)

	// CreateResource 创建资源配置信息
	CreateResource(ctx kratosx.Context, resource *entity.Resource) (uint32, error)

	// UpdateResource 更新资源配置信息
	UpdateResource(ctx kratosx.Context, resource *entity.Resource) error

	// DeleteResource 删除资源配置信息
	DeleteResource(ctx kratosx.Context, id uint32) error

	// GetResourceByKeyword 获取指定的资源配置信息
	GetResourceByKeyword(ctx kratosx.Context, keyword string) (*entity.Resource, error)

	// ListResourceValue 获取业务配置值信息列表
	ListResourceValue(ctx kratosx.Context, req *types.ListResourceValueRequest) ([]*entity.ResourceValue, error)

	// UpdateResourceValues 更新业务配置值信息
	UpdateResourceValues(ctx kratosx.Context, rvs []*entity.ResourceValue) error

	// AllResourceField 获取全部可用的字段
	AllResourceField(ctx kratosx.Context, sid uint32) ([]string, error)
}
