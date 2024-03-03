package server

import (
	"context"
	v1 "github.com/dstgo/wilson/internal/proto/api/v1"
	"github.com/google/wire"
)

var ServerProvider = wire.NewSet(
	NewService,
)

func NewService() *Service {
	return &Service{}
}

type Service struct {
	v1.UnimplementedDstServerServiceServer
}

func (s *Service) Boot(ctx context.Context, request *v1.ControlRequest) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ReBoot(ctx context.Context, request *v1.ControlRequest) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Stop(ctx context.Context, request *v1.ControlRequest) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) State(ctx context.Context, request *v1.ControlRequest) (*v1.StateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Logs(ctx context.Context, request *v1.LogsRequest) (*v1.LogsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ExecuteCommand(ctx context.Context, request *v1.CommandRequest) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Version(ctx context.Context, id *v1.InstanceId) (*v1.VersionResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Update(ctx context.Context, id *v1.InstanceId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RollBack(ctx context.Context, id *v1.InstanceId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Reset(ctx context.Context, id *v1.InstanceId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Clear(ctx context.Context, id *v1.InstanceId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}
