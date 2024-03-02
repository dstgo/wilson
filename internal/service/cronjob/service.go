package cronjob

import (
	"context"
	v1 "github.com/dstgo/wilson/internal/proto/api/v1"
)

type Service struct {
	v1.UnimplementedCronJobServiceServer
}

func (s *Service) Create(ctx context.Context, req *v1.CreateJobReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, req *v1.CreateJobReq) (*v1.NotifyResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) List(ctx context.Context, req *v1.CreateJobReq) (*v1.JobList, error) {
	//TODO implement me
	panic("implement me")
}
