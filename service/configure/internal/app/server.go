package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/client/rpc/manager"
	"github.com/dstgo/wilson/framework/kratosx"
	ktypes "github.com/dstgo/wilson/framework/kratosx/types"

	pb "github.com/dstgo/wilson/api/gen/configure/server/v1"
	"github.com/dstgo/wilson/service/configure/internal/conf"
	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/domain/service"
	"github.com/dstgo/wilson/service/configure/internal/infra/dbs"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Server struct {
	pb.UnimplementedServerServer
	srv *service.Server
}

func NewServer(conf *conf.Config) *Server {
	return &Server{
		srv: service.NewServer(
			conf,
			dbs.NewServer(),
			manager.NewPermission(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewServer(c)
		pb.RegisterServerHTTPServer(hs, srv)
		pb.RegisterServerServer(gs, srv)
	})
}

// ListServer 获取服务信息列表
func (s *Server) ListServer(c context.Context, req *pb.ListServerRequest) (*pb.ListServerReply, error) {
	list, total, err := s.srv.ListServer(kratosx.MustContext(c), &types.ListServerRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Keyword:  req.Keyword,
		Name:     req.Name,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListServerReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListServerReply_Server{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Name:        item.Name,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateServer 创建服务信息
func (s *Server) CreateServer(c context.Context, req *pb.CreateServerRequest) (*pb.CreateServerReply, error) {
	id, err := s.srv.CreateServer(kratosx.MustContext(c), &entity.Server{
		Keyword:     req.Keyword,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateServerReply{Id: id}, nil
}

// UpdateServer 更新服务信息
func (s *Server) UpdateServer(c context.Context, req *pb.UpdateServerRequest) (*pb.UpdateServerReply, error) {
	if err := s.srv.UpdateServer(kratosx.MustContext(c), &entity.Server{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateServerReply{}, nil
}

// DeleteServer 删除服务信息
func (s *Server) DeleteServer(c context.Context, req *pb.DeleteServerRequest) (*pb.DeleteServerReply, error) {
	if err := s.srv.DeleteServer(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteServerReply{}, nil
}
