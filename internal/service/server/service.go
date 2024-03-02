package server

import v1 "github.com/dstgo/wilson/internal/proto/api/v1"

type Service struct {
	v1.UnimplementedDstServerServiceServer
}
