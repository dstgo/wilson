package repository

import (
	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Configure interface {
	// GetConfigure 获取指定的配置
	GetConfigure(ctx kratosx.Context, id uint32) (*entity.Configure, error)

	// GetConfigureByEnvAndSrv 通过envId和srvId获取指定的配置
	GetConfigureByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*entity.Configure, error)

	// ListConfigure 获取配置列表
	ListConfigure(ctx kratosx.Context, req *types.ListConfigureRequest) ([]*entity.Configure, uint32, error)

	// CreateConfigure 创建配置
	CreateConfigure(ctx kratosx.Context, c *entity.Configure) (uint32, error)

	// UpdateConfigure 更新配置
	UpdateConfigure(ctx kratosx.Context, c *entity.Configure) error

	// BroadcastConfigure 广播配置变更
	BroadcastConfigure(ctx kratosx.Context, envId uint32, srvId uint32) error

	// SubscribeConfigure 监听配置变更广播
	SubscribeConfigure(f func(ctx kratosx.Context, envId uint32, srvId uint32) error)
}
