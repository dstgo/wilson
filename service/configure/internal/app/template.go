package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/client/rpc/manager"
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/pkg/valx"

	pb "github.com/dstgo/wilson/api/gen/configure/template/v1"
	"github.com/dstgo/wilson/api/gen/errors"
	"github.com/dstgo/wilson/service/configure/internal/conf"
	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/domain/service"
	"github.com/dstgo/wilson/service/configure/internal/infra/dbs"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Template struct {
	pb.UnimplementedTemplateServer
	srv *service.Template
}

func NewTemplate(conf *conf.Config) *Template {
	return &Template{
		srv: service.NewTemplate(
			conf,
			dbs.NewTemplate(),
			dbs.NewBusiness(),
			dbs.NewResource(),
			manager.NewPermission(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewTemplate(c)
		pb.RegisterTemplateHTTPServer(hs, srv)
		pb.RegisterTemplateServer(gs, srv)
	})
}

func (s *Template) CurrentTemplate(ctx context.Context, req *pb.CurrentTemplateRequest) (*pb.CurrentTemplateReply, error) {
	res, err := s.srv.CurrentTemplate(kratosx.MustContext(ctx), req.ServerId)
	if err != nil {
		return nil, err
	}
	reply := pb.CurrentTemplateReply{}
	if err := valx.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Template) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateReply, error) {
	res, err := s.srv.GetTemplate(kratosx.MustContext(ctx), req.Id)
	if err != nil {
		return nil, err
	}
	reply := pb.GetTemplateReply{}
	if err := valx.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Template) ListTemplate(ctx context.Context, req *pb.ListTemplateRequest) (*pb.ListTemplateReply, error) {
	res, total, err := s.srv.ListTemplate(kratosx.MustContext(ctx), &types.ListTemplateRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		ServerId: req.ServerId,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListTemplateReply{Total: total}
	if err := valx.Transform(res, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Template) CreateTemplate(c context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateReply, error) {
	id, err := s.srv.CreateTemplate(kratosx.MustContext(c), &entity.Template{
		ServerId:    req.ServerId,
		Content:     req.Content,
		Description: req.Description,
		Format:      req.Format,
	})
	return &pb.CreateTemplateReply{Id: id}, err
}

func (s *Template) SwitchTemplate(ctx context.Context, req *pb.SwitchTemplateRequest) (*pb.SwitchTemplateReply, error) {
	return &pb.SwitchTemplateReply{}, s.srv.SwitchTemplate(kratosx.MustContext(ctx), req.ServerId, req.Id)
}

func (s *Template) CompareTemplate(ctx context.Context, req *pb.CompareTemplateRequest) (*pb.CompareTemplateReply, error) {
	list, err := s.srv.CompareTemplate(kratosx.MustContext(ctx), &types.CompareTemplateRequest{
		Id:      req.Id,
		Format:  req.Format,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.CompareTemplateReply{}
	if err := valx.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Template) PreviewTemplate(ctx context.Context, req *pb.PreviewTemplateRequest) (*pb.PreviewTemplateReply, error) {
	res, err := s.srv.PreviewTemplate(kratosx.MustContext(ctx), &types.PreviewTemplateRequest{
		EnvId:    req.EnvId,
		ServerId: req.ServerId,
		Format:   req.Format,
		Content:  req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PreviewTemplateReply{Content: res.Content}, nil
}
