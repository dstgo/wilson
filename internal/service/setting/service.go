package setting

import (
	"context"
	v1 "github.com/dstgo/wilson/internal/proto/api/v1"
	"github.com/google/wire"
)

var SettingProvider = wire.NewSet(
	NewService,
)

func NewService() *Service {
	return &Service{}
}

type Service struct {
	v1.UnimplementedSettingServiceServer
}

func (s *Service) GetRoomSetting(ctx context.Context, id *v1.InstanceId) (*v1.RoomSetting, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SaveRoomSetting(ctx context.Context, setting *v1.RoomSetting) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetWorldSetting(ctx context.Context, id *v1.InstanceId) (*v1.WorldSetting, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SaveWorldSetting(ctx context.Context, setting *v1.WorldSetting) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetRawWorldSetting(ctx context.Context, id *v1.InstanceId) (*v1.RawWorldSetting, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SaveRawWorldSetting(ctx context.Context, setting *v1.RawWorldSetting) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}
