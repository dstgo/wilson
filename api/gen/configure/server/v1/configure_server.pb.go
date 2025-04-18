// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: configure/server/v1/configure_server.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListServerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          uint32                 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      uint32                 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Order         *string                `protobuf:"bytes,3,opt,name=order,proto3,oneof" json:"order,omitempty"`
	OrderBy       *string                `protobuf:"bytes,4,opt,name=orderBy,proto3,oneof" json:"orderBy,omitempty"`
	Keyword       *string                `protobuf:"bytes,5,opt,name=keyword,proto3,oneof" json:"keyword,omitempty"`
	Name          *string                `protobuf:"bytes,6,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Status        *bool                  `protobuf:"varint,7,opt,name=status,proto3,oneof" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListServerRequest) Reset() {
	*x = ListServerRequest{}
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServerRequest) ProtoMessage() {}

func (x *ListServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServerRequest.ProtoReflect.Descriptor instead.
func (*ListServerRequest) Descriptor() ([]byte, []int) {
	return file_configure_server_v1_configure_server_proto_rawDescGZIP(), []int{0}
}

func (x *ListServerRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListServerRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListServerRequest) GetOrder() string {
	if x != nil && x.Order != nil {
		return *x.Order
	}
	return ""
}

func (x *ListServerRequest) GetOrderBy() string {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return ""
}

func (x *ListServerRequest) GetKeyword() string {
	if x != nil && x.Keyword != nil {
		return *x.Keyword
	}
	return ""
}

func (x *ListServerRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ListServerRequest) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

type ListServerReply struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Total         uint32                    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List          []*ListServerReply_Server `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListServerReply) Reset() {
	*x = ListServerReply{}
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServerReply) ProtoMessage() {}

func (x *ListServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServerReply.ProtoReflect.Descriptor instead.
func (*ListServerReply) Descriptor() ([]byte, []int) {
	return file_configure_server_v1_configure_server_proto_rawDescGZIP(), []int{1}
}

func (x *ListServerReply) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListServerReply) GetList() []*ListServerReply_Server {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateServerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Keyword       string                 `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   *string                `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Status        *bool                  `protobuf:"varint,4,opt,name=status,proto3,oneof" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateServerRequest) Reset() {
	*x = CreateServerRequest{}
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerRequest) ProtoMessage() {}

func (x *CreateServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerRequest.ProtoReflect.Descriptor instead.
func (*CreateServerRequest) Descriptor() ([]byte, []int) {
	return file_configure_server_v1_configure_server_proto_rawDescGZIP(), []int{2}
}

func (x *CreateServerRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *CreateServerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateServerRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CreateServerRequest) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

type CreateServerReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateServerReply) Reset() {
	*x = CreateServerReply{}
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerReply) ProtoMessage() {}

func (x *CreateServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerReply.ProtoReflect.Descriptor instead.
func (*CreateServerReply) Descriptor() ([]byte, []int) {
	return file_configure_server_v1_configure_server_proto_rawDescGZIP(), []int{3}
}

func (x *CreateServerReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateServerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Keyword       string                 `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Status        *bool                  `protobuf:"varint,4,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Description   *string                `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateServerRequest) Reset() {
	*x = UpdateServerRequest{}
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServerRequest) ProtoMessage() {}

func (x *UpdateServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServerRequest.ProtoReflect.Descriptor instead.
func (*UpdateServerRequest) Descriptor() ([]byte, []int) {
	return file_configure_server_v1_configure_server_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateServerRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateServerRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *UpdateServerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateServerRequest) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

func (x *UpdateServerRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type UpdateServerReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateServerReply) Reset() {
	*x = UpdateServerReply{}
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServerReply) ProtoMessage() {}

func (x *UpdateServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServerReply.ProtoReflect.Descriptor instead.
func (*UpdateServerReply) Descriptor() ([]byte, []int) {
	return file_configure_server_v1_configure_server_proto_rawDescGZIP(), []int{5}
}

type DeleteServerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteServerRequest) Reset() {
	*x = DeleteServerRequest{}
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServerRequest) ProtoMessage() {}

func (x *DeleteServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServerRequest.ProtoReflect.Descriptor instead.
func (*DeleteServerRequest) Descriptor() ([]byte, []int) {
	return file_configure_server_v1_configure_server_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteServerRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteServerReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteServerReply) Reset() {
	*x = DeleteServerReply{}
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServerReply) ProtoMessage() {}

func (x *DeleteServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServerReply.ProtoReflect.Descriptor instead.
func (*DeleteServerReply) Descriptor() ([]byte, []int) {
	return file_configure_server_v1_configure_server_proto_rawDescGZIP(), []int{7}
}

type ListServerReply_Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Keyword       string                 `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description   *string                `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Status        *bool                  `protobuf:"varint,5,opt,name=status,proto3,oneof" json:"status,omitempty"`
	CreatedAt     uint32                 `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     uint32                 `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListServerReply_Server) Reset() {
	*x = ListServerReply_Server{}
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListServerReply_Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServerReply_Server) ProtoMessage() {}

func (x *ListServerReply_Server) ProtoReflect() protoreflect.Message {
	mi := &file_configure_server_v1_configure_server_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServerReply_Server.ProtoReflect.Descriptor instead.
func (*ListServerReply_Server) Descriptor() ([]byte, []int) {
	return file_configure_server_v1_configure_server_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListServerReply_Server) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListServerReply_Server) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListServerReply_Server) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListServerReply_Server) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ListServerReply_Server) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

func (x *ListServerReply_Server) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ListServerReply_Server) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_configure_server_v1_configure_server_proto protoreflect.FileDescriptor

var file_configure_server_v1_configure_server_proto_rawDesc = string([]byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x77, 0x69,
	0x6c, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x02, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x20, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a,
	0x04, 0x18, 0x32, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x2b, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10,
	0xfa, 0x42, 0x0d, 0x72, 0x0b, 0x52, 0x03, 0x61, 0x73, 0x63, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x48, 0x00, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x52, 0x02, 0x69, 0x64, 0x48, 0x01, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0xd7, 0x02, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x4a, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x77, 0x69, 0x6c, 0x73,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0xe1, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbb, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2e, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x18, 0x5a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_configure_server_v1_configure_server_proto_rawDescOnce sync.Once
	file_configure_server_v1_configure_server_proto_rawDescData []byte
)

func file_configure_server_v1_configure_server_proto_rawDescGZIP() []byte {
	file_configure_server_v1_configure_server_proto_rawDescOnce.Do(func() {
		file_configure_server_v1_configure_server_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_configure_server_v1_configure_server_proto_rawDesc), len(file_configure_server_v1_configure_server_proto_rawDesc)))
	})
	return file_configure_server_v1_configure_server_proto_rawDescData
}

var file_configure_server_v1_configure_server_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_configure_server_v1_configure_server_proto_goTypes = []any{
	(*ListServerRequest)(nil),      // 0: wilson.api.configure.server.v1.ListServerRequest
	(*ListServerReply)(nil),        // 1: wilson.api.configure.server.v1.ListServerReply
	(*CreateServerRequest)(nil),    // 2: wilson.api.configure.server.v1.CreateServerRequest
	(*CreateServerReply)(nil),      // 3: wilson.api.configure.server.v1.CreateServerReply
	(*UpdateServerRequest)(nil),    // 4: wilson.api.configure.server.v1.UpdateServerRequest
	(*UpdateServerReply)(nil),      // 5: wilson.api.configure.server.v1.UpdateServerReply
	(*DeleteServerRequest)(nil),    // 6: wilson.api.configure.server.v1.DeleteServerRequest
	(*DeleteServerReply)(nil),      // 7: wilson.api.configure.server.v1.DeleteServerReply
	(*ListServerReply_Server)(nil), // 8: wilson.api.configure.server.v1.ListServerReply.Server
}
var file_configure_server_v1_configure_server_proto_depIdxs = []int32{
	8, // 0: wilson.api.configure.server.v1.ListServerReply.list:type_name -> wilson.api.configure.server.v1.ListServerReply.Server
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_configure_server_v1_configure_server_proto_init() }
func file_configure_server_v1_configure_server_proto_init() {
	if File_configure_server_v1_configure_server_proto != nil {
		return
	}
	file_configure_server_v1_configure_server_proto_msgTypes[0].OneofWrappers = []any{}
	file_configure_server_v1_configure_server_proto_msgTypes[2].OneofWrappers = []any{}
	file_configure_server_v1_configure_server_proto_msgTypes[4].OneofWrappers = []any{}
	file_configure_server_v1_configure_server_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_configure_server_v1_configure_server_proto_rawDesc), len(file_configure_server_v1_configure_server_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_configure_server_v1_configure_server_proto_goTypes,
		DependencyIndexes: file_configure_server_v1_configure_server_proto_depIdxs,
		MessageInfos:      file_configure_server_v1_configure_server_proto_msgTypes,
	}.Build()
	File_configure_server_v1_configure_server_proto = out.File
	file_configure_server_v1_configure_server_proto_goTypes = nil
	file_configure_server_v1_configure_server_proto_depIdxs = nil
}
