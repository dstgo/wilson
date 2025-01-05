package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/dstgo/wilson/framework/kratosx"

	pb "github.com/dstgo/wilson/api/gen/configure/user/v1"
	"github.com/dstgo/wilson/service/configure/internal/conf"
	"github.com/dstgo/wilson/service/configure/internal/domain/service"
)

type User struct {
	pb.UnimplementedUserServer
	srv *service.User
}

func NewUser(conf *conf.Config) *User {
	return &User{
		srv: service.NewUser(conf),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewUser(c)
		pb.RegisterUserHTTPServer(hs, srv)
		pb.RegisterUserServer(gs, srv)
	})
}

func (s *User) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := s.srv.Login(kratosx.MustContext(ctx), in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: token}, nil
}

func (s *User) RefreshToken(ctx context.Context, _ *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	token, err := s.srv.RefreshToken(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	return &pb.RefreshTokenReply{Token: token}, nil
}
