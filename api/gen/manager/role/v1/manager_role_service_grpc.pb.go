// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: manager/role/v1/manager_role_service.proto

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
	Role_GetRoleMenuIds_FullMethodName   = "/wilson.api.manager.role.v1.Role/GetRoleMenuIds"
	Role_ListRole_FullMethodName         = "/wilson.api.manager.role.v1.Role/ListRole"
	Role_CreateRole_FullMethodName       = "/wilson.api.manager.role.v1.Role/CreateRole"
	Role_UpdateRole_FullMethodName       = "/wilson.api.manager.role.v1.Role/UpdateRole"
	Role_UpdateRoleMenu_FullMethodName   = "/wilson.api.manager.role.v1.Role/UpdateRoleMenu"
	Role_UpdateRoleStatus_FullMethodName = "/wilson.api.manager.role.v1.Role/UpdateRoleStatus"
	Role_DeleteRole_FullMethodName       = "/wilson.api.manager.role.v1.Role/DeleteRole"
	Role_GetRole_FullMethodName          = "/wilson.api.manager.role.v1.Role/GetRole"
)

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleClient interface {
	// GetRoleMenuIds 获取指定角色的菜单id列表
	GetRoleMenuIds(ctx context.Context, in *GetRoleMenuIdsRequest, opts ...grpc.CallOption) (*GetRoleMenuIdsReply, error)
	// ListRole 获取角色信息列表
	ListRole(ctx context.Context, in *ListRoleRequest, opts ...grpc.CallOption) (*ListRoleReply, error)
	// CreateRole 创建角色信息
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleReply, error)
	// UpdateRole 更新角色信息
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleReply, error)
	// UpdateRole 更新角色信息
	UpdateRoleMenu(ctx context.Context, in *UpdateRoleMenuRequest, opts ...grpc.CallOption) (*UpdateRoleMenuReply, error)
	// UpdateRoleStatus 更新角色信息状态
	UpdateRoleStatus(ctx context.Context, in *UpdateRoleStatusRequest, opts ...grpc.CallOption) (*UpdateRoleStatusReply, error)
	// DeleteRole 删除角色信息
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleReply, error)
	// GetRole 获取指定的角色信息
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleReply, error)
}

type roleClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleClient(cc grpc.ClientConnInterface) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) GetRoleMenuIds(ctx context.Context, in *GetRoleMenuIdsRequest, opts ...grpc.CallOption) (*GetRoleMenuIdsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoleMenuIdsReply)
	err := c.cc.Invoke(ctx, Role_GetRoleMenuIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) ListRole(ctx context.Context, in *ListRoleRequest, opts ...grpc.CallOption) (*ListRoleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRoleReply)
	err := c.cc.Invoke(ctx, Role_ListRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRoleReply)
	err := c.cc.Invoke(ctx, Role_CreateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRoleReply)
	err := c.cc.Invoke(ctx, Role_UpdateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) UpdateRoleMenu(ctx context.Context, in *UpdateRoleMenuRequest, opts ...grpc.CallOption) (*UpdateRoleMenuReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRoleMenuReply)
	err := c.cc.Invoke(ctx, Role_UpdateRoleMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) UpdateRoleStatus(ctx context.Context, in *UpdateRoleStatusRequest, opts ...grpc.CallOption) (*UpdateRoleStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRoleStatusReply)
	err := c.cc.Invoke(ctx, Role_UpdateRoleStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRoleReply)
	err := c.cc.Invoke(ctx, Role_DeleteRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRoleReply)
	err := c.cc.Invoke(ctx, Role_GetRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
// All implementations must embed UnimplementedRoleServer
// for forward compatibility.
type RoleServer interface {
	// GetRoleMenuIds 获取指定角色的菜单id列表
	GetRoleMenuIds(context.Context, *GetRoleMenuIdsRequest) (*GetRoleMenuIdsReply, error)
	// ListRole 获取角色信息列表
	ListRole(context.Context, *ListRoleRequest) (*ListRoleReply, error)
	// CreateRole 创建角色信息
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleReply, error)
	// UpdateRole 更新角色信息
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleReply, error)
	// UpdateRole 更新角色信息
	UpdateRoleMenu(context.Context, *UpdateRoleMenuRequest) (*UpdateRoleMenuReply, error)
	// UpdateRoleStatus 更新角色信息状态
	UpdateRoleStatus(context.Context, *UpdateRoleStatusRequest) (*UpdateRoleStatusReply, error)
	// DeleteRole 删除角色信息
	DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleReply, error)
	// GetRole 获取指定的角色信息
	GetRole(context.Context, *GetRoleRequest) (*GetRoleReply, error)
	mustEmbedUnimplementedRoleServer()
}

// UnimplementedRoleServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoleServer struct{}

func (UnimplementedRoleServer) GetRoleMenuIds(context.Context, *GetRoleMenuIdsRequest) (*GetRoleMenuIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleMenuIds not implemented")
}
func (UnimplementedRoleServer) ListRole(context.Context, *ListRoleRequest) (*ListRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedRoleServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRoleServer) UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleServer) UpdateRoleMenu(context.Context, *UpdateRoleMenuRequest) (*UpdateRoleMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoleMenu not implemented")
}
func (UnimplementedRoleServer) UpdateRoleStatus(context.Context, *UpdateRoleStatusRequest) (*UpdateRoleStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoleStatus not implemented")
}
func (UnimplementedRoleServer) DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRoleServer) GetRole(context.Context, *GetRoleRequest) (*GetRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedRoleServer) mustEmbedUnimplementedRoleServer() {}
func (UnimplementedRoleServer) testEmbeddedByValue()              {}

// UnsafeRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServer will
// result in compilation errors.
type UnsafeRoleServer interface {
	mustEmbedUnimplementedRoleServer()
}

func RegisterRoleServer(s grpc.ServiceRegistrar, srv RoleServer) {
	// If the following call pancis, it indicates UnimplementedRoleServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Role_ServiceDesc, srv)
}

func _Role_GetRoleMenuIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleMenuIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).GetRoleMenuIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_GetRoleMenuIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).GetRoleMenuIds(ctx, req.(*GetRoleMenuIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_ListRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).ListRole(ctx, req.(*ListRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_UpdateRoleMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).UpdateRoleMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_UpdateRoleMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).UpdateRoleMenu(ctx, req.(*UpdateRoleMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_UpdateRoleStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).UpdateRoleStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_UpdateRoleStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).UpdateRoleStatus(ctx, req.(*UpdateRoleStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Role_ServiceDesc is the grpc.ServiceDesc for Role service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Role_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wilson.api.manager.role.v1.Role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoleMenuIds",
			Handler:    _Role_GetRoleMenuIds_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _Role_ListRole_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _Role_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _Role_UpdateRole_Handler,
		},
		{
			MethodName: "UpdateRoleMenu",
			Handler:    _Role_UpdateRoleMenu_Handler,
		},
		{
			MethodName: "UpdateRoleStatus",
			Handler:    _Role_UpdateRoleStatus_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _Role_DeleteRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _Role_GetRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager/role/v1/manager_role_service.proto",
}
