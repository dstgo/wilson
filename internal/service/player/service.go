package player

import (
	"context"
	v1 "github.com/dstgo/wilson/internal/proto/api/v1"
	"github.com/google/wire"
)

var PlayerProvider = wire.NewSet(
	NewService,
)

func NewService() *Service {
	return &Service{}
}

type Service struct {
	v1.UnimplementedPlayerServiceServer
}

func (s *Service) GetPlayerStats(ctx context.Context, id *v1.InstanceId) (*v1.PlayerStatisticInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetPlayerChatLog(ctx context.Context, id *v1.InstanceId) (*v1.PlayerChatLog, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ExecutePlayer(ctx context.Context, req *v1.ExecutePlayerReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetWhiteList(ctx context.Context, id *v1.InstanceId) (*v1.PlayerListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetBlackList(ctx context.Context, id *v1.InstanceId) (*v1.PlayerListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetAdminList(ctx context.Context, id *v1.InstanceId) (*v1.PlayerListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AddWhiteList(ctx context.Context, req *v1.PlayerListReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AddBlackList(ctx context.Context, req *v1.PlayerListReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AddAdminList(ctx context.Context, req *v1.PlayerListReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RemoveWhiteList(ctx context.Context, req *v1.PlayerListReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RemoveBlackList(ctx context.Context, req *v1.PlayerListReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RemoveAdminList(ctx context.Context, req *v1.PlayerListReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}
