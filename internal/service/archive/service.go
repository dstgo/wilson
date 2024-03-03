package archive

import (
	"context"
	v1 "github.com/dstgo/wilson/internal/proto/api/v1"
	"github.com/google/wire"
)

var ArchiveProvider = wire.NewSet(
	NewService,
)

func NewService() *Service {
	return &Service{}
}

type Service struct {
	v1.UnimplementedArchiveServiceServer
}

func (s *Service) Info(ctx context.Context, id *v1.InstanceId) (*v1.ArchiveInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ListBackups(ctx context.Context, id *v1.InstanceId) (*v1.BackUpList, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CreateBackup(ctx context.Context, opt *v1.BackupOpt) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteBackUp(ctx context.Context, opt *v1.BackupOpt) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RestoreBackUp(ctx context.Context, opt *v1.BackupOpt) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UploadBackup(ctx context.Context, file *v1.BackupFile) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DownloadBackup(ctx context.Context, opt *v1.BackupOpt) (*v1.BackupFile, error) {
	//TODO implement me
	panic("implement me")
}
