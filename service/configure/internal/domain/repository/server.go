package repository

import (
	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Server interface {
	// GetServerByKeyword 获取指定的服务信息
	GetServerByKeyword(ctx kratosx.Context, keyword string) (*entity.Server, error)

	// GetServer 获取指定的服务信息
	GetServer(ctx kratosx.Context, id uint32) (*entity.Server, error)

	// ListServer 获取服务信息列表
	ListServer(ctx kratosx.Context, req *types.ListServerRequest) ([]*entity.Server, uint32, error)

	// CreateServer 创建服务信息
	CreateServer(ctx kratosx.Context, server *entity.Server) (uint32, error)

	// UpdateServer 更新服务信息
	UpdateServer(ctx kratosx.Context, server *entity.Server) error

	// DeleteServer 删除服务信息
	DeleteServer(ctx kratosx.Context, id uint32) error
}
