// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: manager/department/v1/manager_department.proto

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

type GetDepartmentRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Params:
	//
	//	*GetDepartmentRequest_Id
	//	*GetDepartmentRequest_Keyword
	Params        isGetDepartmentRequest_Params `protobuf_oneof:"params"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDepartmentRequest) Reset() {
	*x = GetDepartmentRequest{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepartmentRequest) ProtoMessage() {}

func (x *GetDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepartmentRequest.ProtoReflect.Descriptor instead.
func (*GetDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{0}
}

func (x *GetDepartmentRequest) GetParams() isGetDepartmentRequest_Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GetDepartmentRequest) GetId() uint32 {
	if x != nil {
		if x, ok := x.Params.(*GetDepartmentRequest_Id); ok {
			return x.Id
		}
	}
	return 0
}

func (x *GetDepartmentRequest) GetKeyword() string {
	if x != nil {
		if x, ok := x.Params.(*GetDepartmentRequest_Keyword); ok {
			return x.Keyword
		}
	}
	return ""
}

type isGetDepartmentRequest_Params interface {
	isGetDepartmentRequest_Params()
}

type GetDepartmentRequest_Id struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type GetDepartmentRequest_Keyword struct {
	Keyword string `protobuf:"bytes,2,opt,name=keyword,proto3,oneof"`
}

func (*GetDepartmentRequest_Id) isGetDepartmentRequest_Params() {}

func (*GetDepartmentRequest_Keyword) isGetDepartmentRequest_Params() {}

type GetDepartmentReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId      uint32                 `protobuf:"varint,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Keyword       string                 `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Description   *string                `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	CreatedAt     uint32                 `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     uint32                 `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDepartmentReply) Reset() {
	*x = GetDepartmentReply{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDepartmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepartmentReply) ProtoMessage() {}

func (x *GetDepartmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepartmentReply.ProtoReflect.Descriptor instead.
func (*GetDepartmentReply) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{1}
}

func (x *GetDepartmentReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetDepartmentReply) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *GetDepartmentReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetDepartmentReply) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *GetDepartmentReply) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *GetDepartmentReply) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GetDepartmentReply) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type ListDepartmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          *string                `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Keyword       *string                `protobuf:"bytes,4,opt,name=keyword,proto3,oneof" json:"keyword,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDepartmentRequest) Reset() {
	*x = ListDepartmentRequest{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDepartmentRequest) ProtoMessage() {}

func (x *ListDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDepartmentRequest.ProtoReflect.Descriptor instead.
func (*ListDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{2}
}

func (x *ListDepartmentRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ListDepartmentRequest) GetKeyword() string {
	if x != nil && x.Keyword != nil {
		return *x.Keyword
	}
	return ""
}

type ListDepartmentReply struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	List          []*ListDepartmentReply_Department `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDepartmentReply) Reset() {
	*x = ListDepartmentReply{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDepartmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDepartmentReply) ProtoMessage() {}

func (x *ListDepartmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDepartmentReply.ProtoReflect.Descriptor instead.
func (*ListDepartmentReply) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{3}
}

func (x *ListDepartmentReply) GetList() []*ListDepartmentReply_Department {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateDepartmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParentId      uint32                 `protobuf:"varint,1,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Keyword       string                 `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Description   *string                `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDepartmentRequest) Reset() {
	*x = CreateDepartmentRequest{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDepartmentRequest) ProtoMessage() {}

func (x *CreateDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDepartmentRequest.ProtoReflect.Descriptor instead.
func (*CreateDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDepartmentRequest) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CreateDepartmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDepartmentRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *CreateDepartmentRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type CreateDepartmentReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDepartmentReply) Reset() {
	*x = CreateDepartmentReply{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDepartmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDepartmentReply) ProtoMessage() {}

func (x *CreateDepartmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDepartmentReply.ProtoReflect.Descriptor instead.
func (*CreateDepartmentReply) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDepartmentReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateDepartmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId      uint32                 `protobuf:"varint,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description   *string                `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDepartmentRequest) Reset() {
	*x = UpdateDepartmentRequest{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDepartmentRequest) ProtoMessage() {}

func (x *UpdateDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDepartmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDepartmentRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDepartmentRequest) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *UpdateDepartmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDepartmentRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type UpdateDepartmentReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDepartmentReply) Reset() {
	*x = UpdateDepartmentReply{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDepartmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDepartmentReply) ProtoMessage() {}

func (x *UpdateDepartmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDepartmentReply.ProtoReflect.Descriptor instead.
func (*UpdateDepartmentReply) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{7}
}

type DeleteDepartmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDepartmentRequest) Reset() {
	*x = DeleteDepartmentRequest{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDepartmentRequest) ProtoMessage() {}

func (x *DeleteDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDepartmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDepartmentRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteDepartmentReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDepartmentReply) Reset() {
	*x = DeleteDepartmentReply{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDepartmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDepartmentReply) ProtoMessage() {}

func (x *DeleteDepartmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDepartmentReply.ProtoReflect.Descriptor instead.
func (*DeleteDepartmentReply) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{9}
}

type ListDepartmentReply_Department struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Id            uint32                            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId      uint32                            `protobuf:"varint,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Name          string                            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Keyword       string                            `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Description   *string                           `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	CreatedAt     uint32                            `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     uint32                            `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Children      []*ListDepartmentReply_Department `protobuf:"bytes,8,rep,name=children,proto3" json:"children,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDepartmentReply_Department) Reset() {
	*x = ListDepartmentReply_Department{}
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDepartmentReply_Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDepartmentReply_Department) ProtoMessage() {}

func (x *ListDepartmentReply_Department) ProtoReflect() protoreflect.Message {
	mi := &file_manager_department_v1_manager_department_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDepartmentReply_Department.ProtoReflect.Descriptor instead.
func (*ListDepartmentReply_Department) Descriptor() ([]byte, []int) {
	return file_manager_department_v1_manager_department_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListDepartmentReply_Department) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListDepartmentReply_Department) GetParentId() uint32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *ListDepartmentReply_Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListDepartmentReply_Department) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListDepartmentReply_Department) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ListDepartmentReply_Department) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ListDepartmentReply_Department) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ListDepartmentReply_Department) GetChildren() []*ListDepartmentReply_Department {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_manager_department_v1_manager_department_proto protoreflect.FileDescriptor

var file_manager_department_v1_manager_department_proto_rawDesc = string([]byte{
	0x0a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x77, 0x69, 0x6c, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xe1, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x64, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xa5, 0x03, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x54, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x77, 0x69, 0x6c, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0xb7, 0x02, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x5c, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x77, 0x69, 0x6c, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xb5, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xab, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x32, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_manager_department_v1_manager_department_proto_rawDescOnce sync.Once
	file_manager_department_v1_manager_department_proto_rawDescData []byte
)

func file_manager_department_v1_manager_department_proto_rawDescGZIP() []byte {
	file_manager_department_v1_manager_department_proto_rawDescOnce.Do(func() {
		file_manager_department_v1_manager_department_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_manager_department_v1_manager_department_proto_rawDesc), len(file_manager_department_v1_manager_department_proto_rawDesc)))
	})
	return file_manager_department_v1_manager_department_proto_rawDescData
}

var file_manager_department_v1_manager_department_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_manager_department_v1_manager_department_proto_goTypes = []any{
	(*GetDepartmentRequest)(nil),           // 0: wilson.api.manager.department.v1.GetDepartmentRequest
	(*GetDepartmentReply)(nil),             // 1: wilson.api.manager.department.v1.GetDepartmentReply
	(*ListDepartmentRequest)(nil),          // 2: wilson.api.manager.department.v1.ListDepartmentRequest
	(*ListDepartmentReply)(nil),            // 3: wilson.api.manager.department.v1.ListDepartmentReply
	(*CreateDepartmentRequest)(nil),        // 4: wilson.api.manager.department.v1.CreateDepartmentRequest
	(*CreateDepartmentReply)(nil),          // 5: wilson.api.manager.department.v1.CreateDepartmentReply
	(*UpdateDepartmentRequest)(nil),        // 6: wilson.api.manager.department.v1.UpdateDepartmentRequest
	(*UpdateDepartmentReply)(nil),          // 7: wilson.api.manager.department.v1.UpdateDepartmentReply
	(*DeleteDepartmentRequest)(nil),        // 8: wilson.api.manager.department.v1.DeleteDepartmentRequest
	(*DeleteDepartmentReply)(nil),          // 9: wilson.api.manager.department.v1.DeleteDepartmentReply
	(*ListDepartmentReply_Department)(nil), // 10: wilson.api.manager.department.v1.ListDepartmentReply.Department
}
var file_manager_department_v1_manager_department_proto_depIdxs = []int32{
	10, // 0: wilson.api.manager.department.v1.ListDepartmentReply.list:type_name -> wilson.api.manager.department.v1.ListDepartmentReply.Department
	10, // 1: wilson.api.manager.department.v1.ListDepartmentReply.Department.children:type_name -> wilson.api.manager.department.v1.ListDepartmentReply.Department
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_manager_department_v1_manager_department_proto_init() }
func file_manager_department_v1_manager_department_proto_init() {
	if File_manager_department_v1_manager_department_proto != nil {
		return
	}
	file_manager_department_v1_manager_department_proto_msgTypes[0].OneofWrappers = []any{
		(*GetDepartmentRequest_Id)(nil),
		(*GetDepartmentRequest_Keyword)(nil),
	}
	file_manager_department_v1_manager_department_proto_msgTypes[1].OneofWrappers = []any{}
	file_manager_department_v1_manager_department_proto_msgTypes[2].OneofWrappers = []any{}
	file_manager_department_v1_manager_department_proto_msgTypes[4].OneofWrappers = []any{}
	file_manager_department_v1_manager_department_proto_msgTypes[6].OneofWrappers = []any{}
	file_manager_department_v1_manager_department_proto_msgTypes[10].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_manager_department_v1_manager_department_proto_rawDesc), len(file_manager_department_v1_manager_department_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manager_department_v1_manager_department_proto_goTypes,
		DependencyIndexes: file_manager_department_v1_manager_department_proto_depIdxs,
		MessageInfos:      file_manager_department_v1_manager_department_proto_msgTypes,
	}.Build()
	File_manager_department_v1_manager_department_proto = out.File
	file_manager_department_v1_manager_department_proto_goTypes = nil
	file_manager_department_v1_manager_department_proto_depIdxs = nil
}
