package app

import (
	"context"

	ktypes "github.com/dstgo/wilson/framework/kratosx/types"

	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/types"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/framework/kratosx"

	pb "github.com/dstgo/wilson/api/gen/configure/business/v1"
	"github.com/dstgo/wilson/api/rpc/manager"
	"github.com/dstgo/wilson/service/configure/internal/conf"
	"github.com/dstgo/wilson/service/configure/internal/domain/service"
	"github.com/dstgo/wilson/service/configure/internal/infra/dbs"
)

type Business struct {
	pb.UnimplementedBusinessServer
	srv *service.Business
}

func NewBusiness(conf *conf.Config) *Business {
	return &Business{
		srv: service.NewBusiness(
			conf,
			dbs.NewBusiness(),
			manager.NewPermission(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewBusiness(c)
		pb.RegisterBusinessHTTPServer(hs, srv)
		pb.RegisterBusinessServer(gs, srv)
	})
}

// ListBusiness 获取业务配置信息列表
func (s *Business) ListBusiness(c context.Context, req *pb.ListBusinessRequest) (*pb.ListBusinessReply, error) {
	list, total, err := s.srv.ListBusiness(kratosx.MustContext(c), &types.ListBusinessRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		ServerId: req.ServerId,
		Keyword:  req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListBusinessReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListBusinessReply_Business{
			Id:          item.Id,
			ServerId:    item.ServerId,
			Keyword:     item.Keyword,
			Type:        item.Type,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateBusiness 创建业务配置信息
func (s *Business) CreateBusiness(c context.Context, req *pb.CreateBusinessRequest) (*pb.CreateBusinessReply, error) {
	id, err := s.srv.CreateBusiness(kratosx.MustContext(c), &entity.Business{
		ServerId:    req.ServerId,
		Keyword:     req.Keyword,
		Type:        req.Type,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateBusinessReply{Id: id}, nil
}

// UpdateBusiness 更新业务配置信息
func (s *Business) UpdateBusiness(c context.Context, req *pb.UpdateBusinessRequest) (*pb.UpdateBusinessReply, error) {
	if err := s.srv.UpdateBusiness(kratosx.MustContext(c), &entity.Business{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Type:        req.Type,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateBusinessReply{}, nil
}

// DeleteBusiness 删除业务配置信息
func (s *Business) DeleteBusiness(c context.Context, req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessReply, error) {
	if err := s.srv.DeleteBusiness(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteBusinessReply{}, nil
}

// ListBusinessValue 获取业务配置值信息列表
func (s *Business) ListBusinessValue(c context.Context, req *pb.ListBusinessValueRequest) (*pb.ListBusinessValueReply, error) {
	list, err := s.srv.ListBusinessValue(kratosx.MustContext(c), req.BusinessId)
	if err != nil {
		return nil, err
	}
	reply := pb.ListBusinessValueReply{
		Total: uint32(len(list)),
	}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListBusinessValueReply_BusinessValue{
			Id:         item.Id,
			EnvId:      item.EnvId,
			BusinessId: item.BusinessId,
			Value:      item.Value,
			CreatedAt:  uint32(item.CreatedAt),
			UpdatedAt:  uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// UpdateBusinessValue 更新业务配置值信息
func (s *Business) UpdateBusinessValue(c context.Context, req *pb.UpdateBusinessValueRequest) (*pb.UpdateBusinessValueReply, error) {
	var list []*entity.BusinessValue
	for _, item := range req.List {
		list = append(list, &entity.BusinessValue{
			BusinessId: req.BusinessId,
			EnvId:      item.EnvId,
			Value:      item.Value,
		})
	}

	if err := s.srv.UpdateBusinessValue(kratosx.MustContext(c), list); err != nil {
		return nil, err
	}
	return &pb.UpdateBusinessValueReply{}, nil
}
