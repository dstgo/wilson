package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"

	"github.com/dstgo/wilson/client/rpc/manager"
	"github.com/dstgo/wilson/framework/kratosx"
	ktypes "github.com/dstgo/wilson/framework/kratosx/types"
	"github.com/dstgo/wilson/framework/pkg/cryptox"

	pb "github.com/dstgo/wilson/api/gen/configure/env/v1"
	"github.com/dstgo/wilson/service/configure/internal/conf"
	"github.com/dstgo/wilson/service/configure/internal/domain/entity"
	"github.com/dstgo/wilson/service/configure/internal/domain/service"
	"github.com/dstgo/wilson/service/configure/internal/infra/dbs"
	"github.com/dstgo/wilson/service/configure/internal/types"
)

type Env struct {
	pb.UnimplementedEnvServer
	srv *service.Env
}

func NewEnv(conf *conf.Config) *Env {
	return &Env{
		srv: service.NewEnv(conf, dbs.NewEnv(), manager.NewPermission()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewEnv(c)
		pb.RegisterEnvHTTPServer(hs, srv)
		pb.RegisterEnvServer(gs, srv)
	})
}

// ListEnv 获取环境信息列表
func (app *Env) ListEnv(c context.Context, req *pb.ListEnvRequest) (*pb.ListEnvReply, error) {
	list, total, err := app.srv.ListEnv(kratosx.MustContext(c), &types.ListEnvRequest{
		Keyword: req.Keyword,
		Name:    req.Name,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListEnvReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListEnvReply_Env{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Name:        item.Name,
			Status:      item.Status,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateEnv 创建环境信息
func (app *Env) CreateEnv(c context.Context, req *pb.CreateEnvRequest) (*pb.CreateEnvReply, error) {
	id, err := app.srv.CreateEnv(kratosx.MustContext(c), &entity.Env{
		Keyword:     req.Keyword,
		Name:        req.Name,
		Status:      req.Status,
		Description: req.Description,
		Token:       cryptox.Sha256HexUpper([]byte(uuid.NewString())),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateEnvReply{Id: id}, nil
}

// UpdateEnv 更新环境信息
func (app *Env) UpdateEnv(c context.Context, req *pb.UpdateEnvRequest) (*pb.UpdateEnvReply, error) {
	if err := app.srv.UpdateEnv(kratosx.MustContext(c), &entity.Env{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Name:        req.Name,
		Status:      req.Status,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateEnvReply{}, nil
}

// DeleteEnv 删除环境信息
func (app *Env) DeleteEnv(c context.Context, req *pb.DeleteEnvRequest) (*pb.DeleteEnvReply, error) {
	if err := app.srv.DeleteEnv(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteEnvReply{}, nil
}

// GetEnvToken 获取环境token
func (app *Env) GetEnvToken(c context.Context, req *pb.GetEnvTokenRequest) (*pb.GetEnvTokenReply, error) {
	token, err := app.srv.GetEnvToken(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetEnvTokenReply{Token: token}, nil
}

// ResetEnvToken 重置token密码
func (app *Env) ResetEnvToken(c context.Context, req *pb.ResetEnvTokenRequest) (*pb.ResetEnvTokenReply, error) {
	in := entity.Env{
		BaseModel: ktypes.BaseModel{Id: req.Id},
		Token:     cryptox.Sha256HexUpper([]byte(uuid.NewString())),
	}
	return &pb.ResetEnvTokenReply{Token: in.Token}, app.srv.UpdateEnv(kratosx.MustContext(c), &in)
}
