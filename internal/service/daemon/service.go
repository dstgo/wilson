package daemon

import (
	"context"
	v1 "github.com/dstgo/wilson/internal/proto/api/v1"
	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/emptypb"
)

var DaemonServiceProvider = wire.NewSet()

type Service struct {
	v1.UnimplementedDaemonServiceServer
}

func (s *Service) HostInfo(ctx context.Context, empty *emptypb.Empty) (*v1.SystemInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*v1.HealthInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) List(ctx context.Context, req *v1.ListContainerReq) (*v1.ListContainerResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Create(ctx context.Context, req *v1.CreateContainerReq) (*v1.CreateContainerResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateQuota(ctx context.Context, resource *v1.Resource) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Log(ctx context.Context, id *v1.InstanceId) (*v1.ContainerLog, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Stats(ctx context.Context, id *v1.InstanceId) (*v1.HealthInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Start(ctx context.Context, id *v1.InstanceId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Stop(ctx context.Context, id *v1.InstanceId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Restart(ctx context.Context, id *v1.InstanceId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, id *v1.InstanceId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}
func (s *Service) ForceDelete(ctx context.Context, id *v1.InstanceId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}
