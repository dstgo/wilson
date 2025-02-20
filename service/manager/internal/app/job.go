package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/framework/kratosx"
	ktypes "github.com/dstgo/wilson/framework/kratosx/types"
	"github.com/dstgo/wilson/framework/pkg/valx"

	"github.com/dstgo/wilson/api/gen/errors"
	pb "github.com/dstgo/wilson/api/gen/manager/job/v1"
	"github.com/dstgo/wilson/service/manager/internal/conf"
	"github.com/dstgo/wilson/service/manager/internal/domain/entity"
	"github.com/dstgo/wilson/service/manager/internal/domain/service"
	"github.com/dstgo/wilson/service/manager/internal/infra/dbs"
	"github.com/dstgo/wilson/service/manager/internal/types"
)

type Job struct {
	pb.UnimplementedJobServer
	srv *service.Job
}

func NewJob(conf *conf.Config) *Job {
	return &Job{
		srv: service.NewJob(conf, dbs.NewJob()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewJob(c)
		pb.RegisterJobHTTPServer(hs, srv)
		pb.RegisterJobServer(gs, srv)
	})
}

// ListJob 获取职位信息列表
func (s *Job) ListJob(c context.Context, req *pb.ListJobRequest) (*pb.ListJobReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListJob(ctx, &types.ListJobRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListJobReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateJob 创建职位信息
func (s *Job) CreateJob(c context.Context, req *pb.CreateJobRequest) (*pb.CreateJobReply, error) {
	id, err := s.srv.CreateJob(kratosx.MustContext(c), &entity.Job{
		Keyword:     req.Keyword,
		Name:        req.Name,
		Weight:      req.Weight,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateJobReply{Id: id}, nil
}

// UpdateJob 更新职位信息
func (s *Job) UpdateJob(c context.Context, req *pb.UpdateJobRequest) (*pb.UpdateJobReply, error) {
	if err := s.srv.UpdateJob(kratosx.MustContext(c), &entity.Job{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Name:        req.Name,
		Weight:      req.Weight,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateJobReply{}, nil
}

// DeleteJob 删除职位信息
func (s *Job) DeleteJob(c context.Context, req *pb.DeleteJobRequest) (*pb.DeleteJobReply, error) {
	return &pb.DeleteJobReply{}, s.srv.DeleteJob(kratosx.MustContext(c), req.Id)
}
