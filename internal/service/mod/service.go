package mod

import (
	"context"
	v1 "github.com/dstgo/wilson/internal/proto/api/v1"
	"github.com/google/wire"
)

var ModProvider = wire.NewSet(
	NewService,
)

func NewService() *Service {
	return &Service{}
}

type Service struct {
	v1.UnimplementedModServiceServer
}

func (s *Service) GetWorkShopModList(ctx context.Context, req *v1.ModListReq) (*v1.ModListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Subscribe(ctx context.Context, id *v1.ModId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Unsubscribe(ctx context.Context, id *v1.ModId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateMod(ctx context.Context, id *v1.ModId) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CheckUpdate(ctx context.Context, id *v1.InstanceId) (*v1.CheckUpdateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetModSettings(ctx context.Context, id *v1.InstanceId) (*v1.ModSettings, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SaveModSettings(ctx context.Context, req *v1.SaveModSettingsReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetRawModSettings(ctx context.Context, id *v1.InstanceId) (*v1.RawModSettings, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SaveRawModSettings(ctx context.Context, req *v1.SaveRawModSettingsReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}
