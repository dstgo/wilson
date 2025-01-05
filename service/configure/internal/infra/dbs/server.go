package dbs

import (
	"fmt"
	"sync"

	"google.golang.org/protobuf/proto"

	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Server struct {
}

var (
	serverIns  *Server
	serverOnce sync.Once
)

func NewServer() *Server {
	serverOnce.Do(func() {
		serverIns = &Server{}
	})
	return serverIns
}

// GetServerByKeyword 获取指定数据
func (r Server) GetServerByKeyword(ctx kratosx.Context, keyword string) (*entity.Server, error) {
	var (
		server entity.Server
		fs     = []string{"*"}
	)
	return &server, ctx.DB().Select(fs).Where("keyword = ?", keyword).First(&server).Error
}

// GetServer 获取指定的数据
func (r Server) GetServer(ctx kratosx.Context, id uint32) (*entity.Server, error) {
	var (
		server entity.Server
		fs     = []string{"*"}
	)
	return &server, ctx.DB().Select(fs).First(&server, id).Error
}

// ListServer 获取列表
func (r Server) ListServer(ctx kratosx.Context, req *types.ListServerRequest) ([]*entity.Server, uint32, error) {
	var (
		list  []*entity.Server
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.Server{}).Select(fs)
	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Ids != nil {
		db = db.Where("id in ?", req.Ids)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateServer 创建数据
func (r Server) CreateServer(ctx kratosx.Context, server *entity.Server) (uint32, error) {
	return server.Id, ctx.DB().Create(server).Error
}

// UpdateServer 更新数据
func (r Server) UpdateServer(ctx kratosx.Context, server *entity.Server) error {
	return ctx.DB().Where("id = ?", server.Id).Updates(server).Error
}

// DeleteServer 删除数据
func (r Server) DeleteServer(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id=?", id).Delete(&entity.Server{}).Error
}
