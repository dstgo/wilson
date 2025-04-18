// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: configure/env/v1/configure_env_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Env_ListEnv_FullMethodName       = "/wilson.api.configure.env.v1.Env/ListEnv"
	Env_CreateEnv_FullMethodName     = "/wilson.api.configure.env.v1.Env/CreateEnv"
	Env_UpdateEnv_FullMethodName     = "/wilson.api.configure.env.v1.Env/UpdateEnv"
	Env_DeleteEnv_FullMethodName     = "/wilson.api.configure.env.v1.Env/DeleteEnv"
	Env_GetEnvToken_FullMethodName   = "/wilson.api.configure.env.v1.Env/GetEnvToken"
	Env_ResetEnvToken_FullMethodName = "/wilson.api.configure.env.v1.Env/ResetEnvToken"
)

// EnvClient is the client API for Env service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvClient interface {
	// ListEnv
	ListEnv(ctx context.Context, in *ListEnvRequest, opts ...grpc.CallOption) (*ListEnvReply, error)
	// CreateEnv
	CreateEnv(ctx context.Context, in *CreateEnvRequest, opts ...grpc.CallOption) (*CreateEnvReply, error)
	// UpdateEnv
	UpdateEnv(ctx context.Context, in *UpdateEnvRequest, opts ...grpc.CallOption) (*UpdateEnvReply, error)
	// DeleteEnv
	DeleteEnv(ctx context.Context, in *DeleteEnvRequest, opts ...grpc.CallOption) (*DeleteEnvReply, error)
	// GetEnvToken
	GetEnvToken(ctx context.Context, in *GetEnvTokenRequest, opts ...grpc.CallOption) (*GetEnvTokenReply, error)
	// ResetEnvToken
	ResetEnvToken(ctx context.Context, in *ResetEnvTokenRequest, opts ...grpc.CallOption) (*ResetEnvTokenReply, error)
}

type envClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvClient(cc grpc.ClientConnInterface) EnvClient {
	return &envClient{cc}
}

func (c *envClient) ListEnv(ctx context.Context, in *ListEnvRequest, opts ...grpc.CallOption) (*ListEnvReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEnvReply)
	err := c.cc.Invoke(ctx, Env_ListEnv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *envClient) CreateEnv(ctx context.Context, in *CreateEnvRequest, opts ...grpc.CallOption) (*CreateEnvReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateEnvReply)
	err := c.cc.Invoke(ctx, Env_CreateEnv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *envClient) UpdateEnv(ctx context.Context, in *UpdateEnvRequest, opts ...grpc.CallOption) (*UpdateEnvReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateEnvReply)
	err := c.cc.Invoke(ctx, Env_UpdateEnv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *envClient) DeleteEnv(ctx context.Context, in *DeleteEnvRequest, opts ...grpc.CallOption) (*DeleteEnvReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteEnvReply)
	err := c.cc.Invoke(ctx, Env_DeleteEnv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *envClient) GetEnvToken(ctx context.Context, in *GetEnvTokenRequest, opts ...grpc.CallOption) (*GetEnvTokenReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEnvTokenReply)
	err := c.cc.Invoke(ctx, Env_GetEnvToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *envClient) ResetEnvToken(ctx context.Context, in *ResetEnvTokenRequest, opts ...grpc.CallOption) (*ResetEnvTokenReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetEnvTokenReply)
	err := c.cc.Invoke(ctx, Env_ResetEnvToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvServer is the server API for Env service.
// All implementations must embed UnimplementedEnvServer
// for forward compatibility.
type EnvServer interface {
	// ListEnv
	ListEnv(context.Context, *ListEnvRequest) (*ListEnvReply, error)
	// CreateEnv
	CreateEnv(context.Context, *CreateEnvRequest) (*CreateEnvReply, error)
	// UpdateEnv
	UpdateEnv(context.Context, *UpdateEnvRequest) (*UpdateEnvReply, error)
	// DeleteEnv
	DeleteEnv(context.Context, *DeleteEnvRequest) (*DeleteEnvReply, error)
	// GetEnvToken
	GetEnvToken(context.Context, *GetEnvTokenRequest) (*GetEnvTokenReply, error)
	// ResetEnvToken
	ResetEnvToken(context.Context, *ResetEnvTokenRequest) (*ResetEnvTokenReply, error)
	mustEmbedUnimplementedEnvServer()
}

// UnimplementedEnvServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEnvServer struct{}

func (UnimplementedEnvServer) ListEnv(context.Context, *ListEnvRequest) (*ListEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnv not implemented")
}
func (UnimplementedEnvServer) CreateEnv(context.Context, *CreateEnvRequest) (*CreateEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEnv not implemented")
}
func (UnimplementedEnvServer) UpdateEnv(context.Context, *UpdateEnvRequest) (*UpdateEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEnv not implemented")
}
func (UnimplementedEnvServer) DeleteEnv(context.Context, *DeleteEnvRequest) (*DeleteEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEnv not implemented")
}
func (UnimplementedEnvServer) GetEnvToken(context.Context, *GetEnvTokenRequest) (*GetEnvTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnvToken not implemented")
}
func (UnimplementedEnvServer) ResetEnvToken(context.Context, *ResetEnvTokenRequest) (*ResetEnvTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetEnvToken not implemented")
}
func (UnimplementedEnvServer) mustEmbedUnimplementedEnvServer() {}
func (UnimplementedEnvServer) testEmbeddedByValue()             {}

// UnsafeEnvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvServer will
// result in compilation errors.
type UnsafeEnvServer interface {
	mustEmbedUnimplementedEnvServer()
}

func RegisterEnvServer(s grpc.ServiceRegistrar, srv EnvServer) {
	// If the following call pancis, it indicates UnimplementedEnvServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Env_ServiceDesc, srv)
}

func _Env_ListEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvServer).ListEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Env_ListEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvServer).ListEnv(ctx, req.(*ListEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Env_CreateEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvServer).CreateEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Env_CreateEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvServer).CreateEnv(ctx, req.(*CreateEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Env_UpdateEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvServer).UpdateEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Env_UpdateEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvServer).UpdateEnv(ctx, req.(*UpdateEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Env_DeleteEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvServer).DeleteEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Env_DeleteEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvServer).DeleteEnv(ctx, req.(*DeleteEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Env_GetEnvToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnvTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvServer).GetEnvToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Env_GetEnvToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvServer).GetEnvToken(ctx, req.(*GetEnvTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Env_ResetEnvToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetEnvTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvServer).ResetEnvToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Env_ResetEnvToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvServer).ResetEnvToken(ctx, req.(*ResetEnvTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Env_ServiceDesc is the grpc.ServiceDesc for Env service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Env_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wilson.api.configure.env.v1.Env",
	HandlerType: (*EnvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEnv",
			Handler:    _Env_ListEnv_Handler,
		},
		{
			MethodName: "CreateEnv",
			Handler:    _Env_CreateEnv_Handler,
		},
		{
			MethodName: "UpdateEnv",
			Handler:    _Env_UpdateEnv_Handler,
		},
		{
			MethodName: "DeleteEnv",
			Handler:    _Env_DeleteEnv_Handler,
		},
		{
			MethodName: "GetEnvToken",
			Handler:    _Env_GetEnvToken_Handler,
		},
		{
			MethodName: "ResetEnvToken",
			Handler:    _Env_ResetEnvToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configure/env/v1/configure_env_service.proto",
}
