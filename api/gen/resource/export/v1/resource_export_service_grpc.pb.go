// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: resource/export/v1/resource_export_service.proto

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
	Export_GetExport_FullMethodName    = "/wilson.api.resource.export.v1.Export/GetExport"
	Export_ListExport_FullMethodName   = "/wilson.api.resource.export.v1.Export/ListExport"
	Export_ExportFile_FullMethodName   = "/wilson.api.resource.export.v1.Export/ExportFile"
	Export_ExportExcel_FullMethodName  = "/wilson.api.resource.export.v1.Export/ExportExcel"
	Export_DeleteExport_FullMethodName = "/wilson.api.resource.export.v1.Export/DeleteExport"
)

// ExportClient is the client API for Export service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExportClient interface {
	// GetExport
	GetExport(ctx context.Context, in *GetExportRequest, opts ...grpc.CallOption) (*GetExportReply, error)
	// ListExport
	ListExport(ctx context.Context, in *ListExportRequest, opts ...grpc.CallOption) (*ListExportReply, error)
	// ExportFile
	ExportFile(ctx context.Context, in *ExportFileRequest, opts ...grpc.CallOption) (*ExportFileReply, error)
	// ExportExcel
	ExportExcel(ctx context.Context, in *ExportExcelRequest, opts ...grpc.CallOption) (*ExportExcelReply, error)
	// DeleteExport
	DeleteExport(ctx context.Context, in *DeleteExportRequest, opts ...grpc.CallOption) (*DeleteExportReply, error)
}

type exportClient struct {
	cc grpc.ClientConnInterface
}

func NewExportClient(cc grpc.ClientConnInterface) ExportClient {
	return &exportClient{cc}
}

func (c *exportClient) GetExport(ctx context.Context, in *GetExportRequest, opts ...grpc.CallOption) (*GetExportReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetExportReply)
	err := c.cc.Invoke(ctx, Export_GetExport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exportClient) ListExport(ctx context.Context, in *ListExportRequest, opts ...grpc.CallOption) (*ListExportReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListExportReply)
	err := c.cc.Invoke(ctx, Export_ListExport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exportClient) ExportFile(ctx context.Context, in *ExportFileRequest, opts ...grpc.CallOption) (*ExportFileReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExportFileReply)
	err := c.cc.Invoke(ctx, Export_ExportFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exportClient) ExportExcel(ctx context.Context, in *ExportExcelRequest, opts ...grpc.CallOption) (*ExportExcelReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExportExcelReply)
	err := c.cc.Invoke(ctx, Export_ExportExcel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exportClient) DeleteExport(ctx context.Context, in *DeleteExportRequest, opts ...grpc.CallOption) (*DeleteExportReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteExportReply)
	err := c.cc.Invoke(ctx, Export_DeleteExport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExportServer is the server API for Export service.
// All implementations must embed UnimplementedExportServer
// for forward compatibility.
type ExportServer interface {
	// GetExport
	GetExport(context.Context, *GetExportRequest) (*GetExportReply, error)
	// ListExport
	ListExport(context.Context, *ListExportRequest) (*ListExportReply, error)
	// ExportFile
	ExportFile(context.Context, *ExportFileRequest) (*ExportFileReply, error)
	// ExportExcel
	ExportExcel(context.Context, *ExportExcelRequest) (*ExportExcelReply, error)
	// DeleteExport
	DeleteExport(context.Context, *DeleteExportRequest) (*DeleteExportReply, error)
	mustEmbedUnimplementedExportServer()
}

// UnimplementedExportServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExportServer struct{}

func (UnimplementedExportServer) GetExport(context.Context, *GetExportRequest) (*GetExportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExport not implemented")
}
func (UnimplementedExportServer) ListExport(context.Context, *ListExportRequest) (*ListExportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExport not implemented")
}
func (UnimplementedExportServer) ExportFile(context.Context, *ExportFileRequest) (*ExportFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportFile not implemented")
}
func (UnimplementedExportServer) ExportExcel(context.Context, *ExportExcelRequest) (*ExportExcelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportExcel not implemented")
}
func (UnimplementedExportServer) DeleteExport(context.Context, *DeleteExportRequest) (*DeleteExportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExport not implemented")
}
func (UnimplementedExportServer) mustEmbedUnimplementedExportServer() {}
func (UnimplementedExportServer) testEmbeddedByValue()                {}

// UnsafeExportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExportServer will
// result in compilation errors.
type UnsafeExportServer interface {
	mustEmbedUnimplementedExportServer()
}

func RegisterExportServer(s grpc.ServiceRegistrar, srv ExportServer) {
	// If the following call pancis, it indicates UnimplementedExportServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Export_ServiceDesc, srv)
}

func _Export_GetExport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExportServer).GetExport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Export_GetExport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExportServer).GetExport(ctx, req.(*GetExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Export_ListExport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExportServer).ListExport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Export_ListExport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExportServer).ListExport(ctx, req.(*ListExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Export_ExportFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExportServer).ExportFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Export_ExportFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExportServer).ExportFile(ctx, req.(*ExportFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Export_ExportExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportExcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExportServer).ExportExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Export_ExportExcel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExportServer).ExportExcel(ctx, req.(*ExportExcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Export_DeleteExport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExportServer).DeleteExport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Export_DeleteExport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExportServer).DeleteExport(ctx, req.(*DeleteExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Export_ServiceDesc is the grpc.ServiceDesc for Export service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Export_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wilson.api.resource.export.v1.Export",
	HandlerType: (*ExportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExport",
			Handler:    _Export_GetExport_Handler,
		},
		{
			MethodName: "ListExport",
			Handler:    _Export_ListExport_Handler,
		},
		{
			MethodName: "ExportFile",
			Handler:    _Export_ExportFile_Handler,
		},
		{
			MethodName: "ExportExcel",
			Handler:    _Export_ExportExcel_Handler,
		},
		{
			MethodName: "DeleteExport",
			Handler:    _Export_DeleteExport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource/export/v1/resource_export_service.proto",
}
