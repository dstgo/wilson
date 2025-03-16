package manager

import (
	"github.com/dstgo/wilson/api/gen/errors"
	mauth "github.com/dstgo/wilson/api/gen/manager/auth/v1"
	mresourcev1 "github.com/dstgo/wilson/api/gen/manager/resource/v1"
	"github.com/dstgo/wilson/framework/kratosx"
)

func mAuthClient(ctx kratosx.Context) (mauth.AuthClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Manager)
	if err != nil {
		return nil, errors.ManagerServiceErrorWrap(err)
	}
	return mauth.NewAuthClient(conn), nil
}

func mResourceClient(ctx kratosx.Context) (mresourcev1.ResourceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Manager)
	if err != nil {
		return nil, errors.ManagerServiceErrorWrap(err)
	}
	return mresourcev1.NewResourceClient(conn), nil
}
