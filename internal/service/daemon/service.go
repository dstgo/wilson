package daemon

import v1 "github.com/dstgo/wilson/internal/proto/api/v1"

type DaemonService struct {
	v1.UnimplementedDaemonServiceServer
}
