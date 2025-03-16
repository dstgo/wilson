package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/client/rpc/manager"
	"github.com/dstgo/wilson/framework/kratosx"
	"github.com/dstgo/wilson/framework/pkg/valx"

	pb "github.com/dstgo/wilson/api/gen/configure/configure/v1"
	"github.com/dstgo/wilson/api/gen/errors"
	"github.com/dstgo/wilson/service/configure/internal/conf"
	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/domain/service"
	"github.com/dstgo/wilson/service/configure/internal/infra/dbs"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Configure struct {
	pb.UnimplementedConfigureServer
	srv *service.Configure
}

func NewConfigure(conf *conf.Config) *Configure {
	return &Configure{
		srv: service.NewConfigure(
			conf,
			dbs.NewConfigure(),
			dbs.NewServer(),
			dbs.NewEnv(),
			dbs.NewBusiness(),
			dbs.NewResource(),
			dbs.NewTemplate(),
			manager.NewPermission(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewConfigure(c)
		pb.RegisterConfigureHTTPServer(hs, srv)
		pb.RegisterConfigureServer(gs, srv)
	})
}

func (s *Configure) CompareConfigure(c context.Context, in *pb.CompareConfigureRequest) (*pb.CompareConfigureReply, error) {
	list, err := s.srv.CompareConfigure(kratosx.MustContext(c), &types.CompareConfigureRequest{
		EnvId:    in.EnvId,
		ServerId: in.ServerId,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.CompareConfigureReply{}
	if err := valx.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Configure) GetConfigure(ctx context.Context, in *pb.GetConfigureRequest) (*pb.GetConfigureReply, error) {
	res, err := s.srv.GetConfigureByEnvAndSrv(kratosx.MustContext(ctx), in.EnvId, in.ServerId)
	if err != nil {
		return nil, err
	}
	reply := pb.GetConfigureReply{}
	if err := valx.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Configure) UpdateConfigure(ctx context.Context, in *pb.UpdateConfigureRequest) (*pb.UpdateConfigureReply, error) {
	return &pb.UpdateConfigureReply{}, s.srv.UpdateConfigure(kratosx.MustContext(ctx), &entity.Configure{
		ServerId:    in.ServerId,
		EnvId:       in.EnvId,
		Description: &in.Description,
	})
}

func (s *Configure) WatchConfigure(in *pb.WatchConfigureRequest, reply pb.Configure_WatchConfigureServer) error {
	return s.srv.Watch(kratosx.MustContext(reply.Context()), &types.WatcherConfigRequest{
		Server: in.Server,
		Token:  in.Token,
	}, func(data *types.WatcherConfigureReply) error {
		return reply.Send(&pb.WatchConfigureReply{
			Format:  data.Format,
			Content: data.Content,
		})
	})
}
