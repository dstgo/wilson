// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: manager/user/v1/manager_user_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	User_GetUser_FullMethodName                   = "/wilson.api.manager.user.v1.User/GetUser"
	User_ListUser_FullMethodName                  = "/wilson.api.manager.user.v1.User/ListUser"
	User_CreateUser_FullMethodName                = "/wilson.api.manager.user.v1.User/CreateUser"
	User_UpdateUser_FullMethodName                = "/wilson.api.manager.user.v1.User/UpdateUser"
	User_UpdateUserStatus_FullMethodName          = "/wilson.api.manager.user.v1.User/UpdateUserStatus"
	User_DeleteUser_FullMethodName                = "/wilson.api.manager.user.v1.User/DeleteUser"
	User_GetCurrentUser_FullMethodName            = "/wilson.api.manager.user.v1.User/GetCurrentUser"
	User_ResetUserPassword_FullMethodName         = "/wilson.api.manager.user.v1.User/ResetUserPassword"
	User_UpdateCurrentUser_FullMethodName         = "/wilson.api.manager.user.v1.User/UpdateCurrentUser"
	User_UpdateCurrentUserRole_FullMethodName     = "/wilson.api.manager.user.v1.User/UpdateCurrentUserRole"
	User_UpdateCurrentUserPassword_FullMethodName = "/wilson.api.manager.user.v1.User/UpdateCurrentUserPassword"
	User_UpdateCurrentUserSetting_FullMethodName  = "/wilson.api.manager.user.v1.User/UpdateCurrentUserSetting"
	User_SendCurrentUserCaptcha_FullMethodName    = "/wilson.api.manager.user.v1.User/SendCurrentUserCaptcha"
	User_GetUserLoginCaptcha_FullMethodName       = "/wilson.api.manager.user.v1.User/GetUserLoginCaptcha"
	User_UserLogin_FullMethodName                 = "/wilson.api.manager.user.v1.User/UserLogin"
	User_UserLogout_FullMethodName                = "/wilson.api.manager.user.v1.User/UserLogout"
	User_UserRefreshToken_FullMethodName          = "/wilson.api.manager.user.v1.User/UserRefreshToken"
	User_ListLoginLog_FullMethodName              = "/wilson.api.manager.user.v1.User/ListLoginLog"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// GetUser
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	// ListUser
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error)
	// CreateUser
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error)
	// UpdateUser
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error)
	// UpdateUserStatus
	UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...grpc.CallOption) (*UpdateUserStatusReply, error)
	// DeleteUser
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error)
	// GetCurrentUser
	GetCurrentUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUserReply, error)
	// ResetUserPassword
	ResetUserPassword(ctx context.Context, in *ResetUserPasswordRequest, opts ...grpc.CallOption) (*ResetUserPasswordReply, error)
	// UpdateCurrentUser
	UpdateCurrentUser(ctx context.Context, in *UpdateCurrentUserRequest, opts ...grpc.CallOption) (*UpdateCurrentUserReply, error)
	// UpdateCurrentUserRole
	UpdateCurrentUserRole(ctx context.Context, in *UpdateCurrentUserRoleRequest, opts ...grpc.CallOption) (*UpdateCurrentUserRoleReply, error)
	// UpdateCurrentUserPassword
	UpdateCurrentUserPassword(ctx context.Context, in *UpdateCurrentUserPasswordRequest, opts ...grpc.CallOption) (*UpdateCurrentUserPasswordReply, error)
	// UpdateCurrentUserSetting
	UpdateCurrentUserSetting(ctx context.Context, in *UpdateCurrentUserSettingRequest, opts ...grpc.CallOption) (*UpdateCurrentUserSettingReply, error)
	// SendCurrentUserCaptcha
	SendCurrentUserCaptcha(ctx context.Context, in *SendCurrentUserCaptchaRequest, opts ...grpc.CallOption) (*SendCurrentUserCaptchaReply, error)
	// GetUserLoginCaptcha
	GetUserLoginCaptcha(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUserLoginCaptchaReply, error)
	// UserLogin
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginReply, error)
	// UserLogout
	UserLogout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UserRefreshToken
	UserRefreshToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserRefreshTokenReply, error)
	// ListLoginLog
	ListLoginLog(ctx context.Context, in *ListLoginLogRequest, opts ...grpc.CallOption) (*ListLoginLogReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, User_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, User_ListUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, User_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserReply)
	err := c.cc.Invoke(ctx, User_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...grpc.CallOption) (*UpdateUserStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserStatusReply)
	err := c.cc.Invoke(ctx, User_UpdateUserStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserReply)
	err := c.cc.Invoke(ctx, User_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetCurrentUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, User_GetCurrentUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ResetUserPassword(ctx context.Context, in *ResetUserPasswordRequest, opts ...grpc.CallOption) (*ResetUserPasswordReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetUserPasswordReply)
	err := c.cc.Invoke(ctx, User_ResetUserPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateCurrentUser(ctx context.Context, in *UpdateCurrentUserRequest, opts ...grpc.CallOption) (*UpdateCurrentUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentUserReply)
	err := c.cc.Invoke(ctx, User_UpdateCurrentUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateCurrentUserRole(ctx context.Context, in *UpdateCurrentUserRoleRequest, opts ...grpc.CallOption) (*UpdateCurrentUserRoleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentUserRoleReply)
	err := c.cc.Invoke(ctx, User_UpdateCurrentUserRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateCurrentUserPassword(ctx context.Context, in *UpdateCurrentUserPasswordRequest, opts ...grpc.CallOption) (*UpdateCurrentUserPasswordReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentUserPasswordReply)
	err := c.cc.Invoke(ctx, User_UpdateCurrentUserPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateCurrentUserSetting(ctx context.Context, in *UpdateCurrentUserSettingRequest, opts ...grpc.CallOption) (*UpdateCurrentUserSettingReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentUserSettingReply)
	err := c.cc.Invoke(ctx, User_UpdateCurrentUserSetting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendCurrentUserCaptcha(ctx context.Context, in *SendCurrentUserCaptchaRequest, opts ...grpc.CallOption) (*SendCurrentUserCaptchaReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendCurrentUserCaptchaReply)
	err := c.cc.Invoke(ctx, User_SendCurrentUserCaptcha_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserLoginCaptcha(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetUserLoginCaptchaReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserLoginCaptchaReply)
	err := c.cc.Invoke(ctx, User_GetUserLoginCaptcha_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginReply)
	err := c.cc.Invoke(ctx, User_UserLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLogout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_UserLogout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserRefreshToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserRefreshTokenReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRefreshTokenReply)
	err := c.cc.Invoke(ctx, User_UserRefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListLoginLog(ctx context.Context, in *ListLoginLogRequest, opts ...grpc.CallOption) (*ListLoginLogReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLoginLogReply)
	err := c.cc.Invoke(ctx, User_ListLoginLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility.
type UserServer interface {
	// GetUser
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	// ListUser
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	// CreateUser
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	// UpdateUser
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
	// UpdateUserStatus
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusReply, error)
	// DeleteUser
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	// GetCurrentUser
	GetCurrentUser(context.Context, *emptypb.Empty) (*GetUserReply, error)
	// ResetUserPassword
	ResetUserPassword(context.Context, *ResetUserPasswordRequest) (*ResetUserPasswordReply, error)
	// UpdateCurrentUser
	UpdateCurrentUser(context.Context, *UpdateCurrentUserRequest) (*UpdateCurrentUserReply, error)
	// UpdateCurrentUserRole
	UpdateCurrentUserRole(context.Context, *UpdateCurrentUserRoleRequest) (*UpdateCurrentUserRoleReply, error)
	// UpdateCurrentUserPassword
	UpdateCurrentUserPassword(context.Context, *UpdateCurrentUserPasswordRequest) (*UpdateCurrentUserPasswordReply, error)
	// UpdateCurrentUserSetting
	UpdateCurrentUserSetting(context.Context, *UpdateCurrentUserSettingRequest) (*UpdateCurrentUserSettingReply, error)
	// SendCurrentUserCaptcha
	SendCurrentUserCaptcha(context.Context, *SendCurrentUserCaptchaRequest) (*SendCurrentUserCaptchaReply, error)
	// GetUserLoginCaptcha
	GetUserLoginCaptcha(context.Context, *emptypb.Empty) (*GetUserLoginCaptchaReply, error)
	// UserLogin
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginReply, error)
	// UserLogout
	UserLogout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// UserRefreshToken
	UserRefreshToken(context.Context, *emptypb.Empty) (*UserRefreshTokenReply, error)
	// ListLoginLog
	ListLoginLog(context.Context, *ListLoginLogRequest) (*ListLoginLogReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServer struct{}

func (UnimplementedUserServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) ListUser(context.Context, *ListUserRequest) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedUserServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserStatus not implemented")
}
func (UnimplementedUserServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServer) GetCurrentUser(context.Context, *emptypb.Empty) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedUserServer) ResetUserPassword(context.Context, *ResetUserPasswordRequest) (*ResetUserPasswordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetUserPassword not implemented")
}
func (UnimplementedUserServer) UpdateCurrentUser(context.Context, *UpdateCurrentUserRequest) (*UpdateCurrentUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentUser not implemented")
}
func (UnimplementedUserServer) UpdateCurrentUserRole(context.Context, *UpdateCurrentUserRoleRequest) (*UpdateCurrentUserRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentUserRole not implemented")
}
func (UnimplementedUserServer) UpdateCurrentUserPassword(context.Context, *UpdateCurrentUserPasswordRequest) (*UpdateCurrentUserPasswordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentUserPassword not implemented")
}
func (UnimplementedUserServer) UpdateCurrentUserSetting(context.Context, *UpdateCurrentUserSettingRequest) (*UpdateCurrentUserSettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentUserSetting not implemented")
}
func (UnimplementedUserServer) SendCurrentUserCaptcha(context.Context, *SendCurrentUserCaptchaRequest) (*SendCurrentUserCaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCurrentUserCaptcha not implemented")
}
func (UnimplementedUserServer) GetUserLoginCaptcha(context.Context, *emptypb.Empty) (*GetUserLoginCaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLoginCaptcha not implemented")
}
func (UnimplementedUserServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServer) UserLogout(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogout not implemented")
}
func (UnimplementedUserServer) UserRefreshToken(context.Context, *emptypb.Empty) (*UserRefreshTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRefreshToken not implemented")
}
func (UnimplementedUserServer) ListLoginLog(context.Context, *ListLoginLogRequest) (*ListLoginLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLoginLog not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}
func (UnimplementedUserServer) testEmbeddedByValue()              {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	// If the following call pancis, it indicates UnimplementedUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUserStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserStatus(ctx, req.(*UpdateUserStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetCurrentUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetCurrentUser(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ResetUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ResetUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ResetUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ResetUserPassword(ctx, req.(*ResetUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateCurrentUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateCurrentUser(ctx, req.(*UpdateCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateCurrentUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateCurrentUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateCurrentUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateCurrentUserRole(ctx, req.(*UpdateCurrentUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateCurrentUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateCurrentUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateCurrentUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateCurrentUserPassword(ctx, req.(*UpdateCurrentUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateCurrentUserSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentUserSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateCurrentUserSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateCurrentUserSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateCurrentUserSetting(ctx, req.(*UpdateCurrentUserSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendCurrentUserCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCurrentUserCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendCurrentUserCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SendCurrentUserCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendCurrentUserCaptcha(ctx, req.(*SendCurrentUserCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserLoginCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserLoginCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserLoginCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserLoginCaptcha(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLogout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogout(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserRefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserRefreshToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListLoginLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLoginLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListLoginLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListLoginLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListLoginLog(ctx, req.(*ListLoginLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wilson.api.manager.user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _User_ListUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserStatus",
			Handler:    _User_UpdateUserStatus_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _User_GetCurrentUser_Handler,
		},
		{
			MethodName: "ResetUserPassword",
			Handler:    _User_ResetUserPassword_Handler,
		},
		{
			MethodName: "UpdateCurrentUser",
			Handler:    _User_UpdateCurrentUser_Handler,
		},
		{
			MethodName: "UpdateCurrentUserRole",
			Handler:    _User_UpdateCurrentUserRole_Handler,
		},
		{
			MethodName: "UpdateCurrentUserPassword",
			Handler:    _User_UpdateCurrentUserPassword_Handler,
		},
		{
			MethodName: "UpdateCurrentUserSetting",
			Handler:    _User_UpdateCurrentUserSetting_Handler,
		},
		{
			MethodName: "SendCurrentUserCaptcha",
			Handler:    _User_SendCurrentUserCaptcha_Handler,
		},
		{
			MethodName: "GetUserLoginCaptcha",
			Handler:    _User_GetUserLoginCaptcha_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "UserLogout",
			Handler:    _User_UserLogout_Handler,
		},
		{
			MethodName: "UserRefreshToken",
			Handler:    _User_UserRefreshToken_Handler,
		},
		{
			MethodName: "ListLoginLog",
			Handler:    _User_ListLoginLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager/user/v1/manager_user_service.proto",
}
