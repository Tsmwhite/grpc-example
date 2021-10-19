// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: protobufs/member.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MemberSexEnum int32

const (
	MemberSexEnum_SEX_UNKNOWN MemberSexEnum = 0
	MemberSexEnum_SEX_BOY     MemberSexEnum = 1
	MemberSexEnum_SEX_GIRL    MemberSexEnum = 2
)

// Enum value maps for MemberSexEnum.
var (
	MemberSexEnum_name = map[int32]string{
		0: "SEX_UNKNOWN",
		1: "SEX_BOY",
		2: "SEX_GIRL",
	}
	MemberSexEnum_value = map[string]int32{
		"SEX_UNKNOWN": 0,
		"SEX_BOY":     1,
		"SEX_GIRL":    2,
	}
)

func (x MemberSexEnum) Enum() *MemberSexEnum {
	p := new(MemberSexEnum)
	*p = x
	return p
}

func (x MemberSexEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemberSexEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_protobufs_member_proto_enumTypes[0].Descriptor()
}

func (MemberSexEnum) Type() protoreflect.EnumType {
	return &file_protobufs_member_proto_enumTypes[0]
}

func (x MemberSexEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MemberSexEnum.Descriptor instead.
func (MemberSexEnum) EnumDescriptor() ([]byte, []int) {
	return file_protobufs_member_proto_rawDescGZIP(), []int{0}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int64         `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Sex  MemberSexEnum `protobuf:"varint,4,opt,name=sex,proto3,enum=member.MemberSexEnum" json:"sex,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_protobufs_member_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Member) GetSex() MemberSexEnum {
	if x != nil {
		return x.Sex
	}
	return MemberSexEnum_SEX_UNKNOWN
}

type GetUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserInfoRequest) Reset() {
	*x = GetUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRequest) ProtoMessage() {}

func (x *GetUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRequest.ProtoReflect.Descriptor instead.
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_protobufs_member_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err  string  `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Info *Member `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetUserInfoResponse) Reset() {
	*x = GetUserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResponse) ProtoMessage() {}

func (x *GetUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResponse.ProtoReflect.Descriptor instead.
func (*GetUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_protobufs_member_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserInfoResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetUserInfoResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *GetUserInfoResponse) GetInfo() *Member {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Page    int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size    int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Type    int64  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetUserListRequest) Reset() {
	*x = GetUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_member_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListRequest) ProtoMessage() {}

func (x *GetUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_member_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListRequest.ProtoReflect.Descriptor instead.
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return file_protobufs_member_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserListRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *GetUserListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetUserListRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetUserListRequest) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

type GetUserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err  string    `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Data []*Member `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserListResponse) Reset() {
	*x = GetUserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_member_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListResponse) ProtoMessage() {}

func (x *GetUserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_member_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListResponse.ProtoReflect.Descriptor instead.
func (*GetUserListResponse) Descriptor() ([]byte, []int) {
	return file_protobufs_member_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserListResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetUserListResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *GetUserListResponse) GetData() []*Member {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protobufs_member_proto protoreflect.FileDescriptor

var file_protobufs_member_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x67, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x27, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x78,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x03, 0x73, 0x65, 0x78, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x22, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0x6a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5f, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x78, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0f,
	0x0a, 0x0b, 0x53, 0x45, 0x58, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x45, 0x58, 0x5f, 0x42, 0x4f, 0x59, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x45, 0x58, 0x5f, 0x47, 0x49, 0x52, 0x4c, 0x10, 0x02, 0x32, 0x9f, 0x01, 0x0a, 0x0d, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protobufs_member_proto_rawDescOnce sync.Once
	file_protobufs_member_proto_rawDescData = file_protobufs_member_proto_rawDesc
)

func file_protobufs_member_proto_rawDescGZIP() []byte {
	file_protobufs_member_proto_rawDescOnce.Do(func() {
		file_protobufs_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobufs_member_proto_rawDescData)
	})
	return file_protobufs_member_proto_rawDescData
}

var file_protobufs_member_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobufs_member_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobufs_member_proto_goTypes = []interface{}{
	(MemberSexEnum)(0),          // 0: member.MemberSexEnum
	(*Member)(nil),              // 1: member.Member
	(*GetUserInfoRequest)(nil),  // 2: member.GetUserInfoRequest
	(*GetUserInfoResponse)(nil), // 3: member.GetUserInfoResponse
	(*GetUserListRequest)(nil),  // 4: member.GetUserListRequest
	(*GetUserListResponse)(nil), // 5: member.GetUserListResponse
}
var file_protobufs_member_proto_depIdxs = []int32{
	0, // 0: member.Member.sex:type_name -> member.MemberSexEnum
	1, // 1: member.GetUserInfoResponse.info:type_name -> member.Member
	1, // 2: member.GetUserListResponse.data:type_name -> member.Member
	2, // 3: member.MemberService.GetUserInfo:input_type -> member.GetUserInfoRequest
	4, // 4: member.MemberService.GetUserList:input_type -> member.GetUserListRequest
	3, // 5: member.MemberService.GetUserInfo:output_type -> member.GetUserInfoResponse
	5, // 6: member.MemberService.GetUserList:output_type -> member.GetUserListResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobufs_member_proto_init() }
func file_protobufs_member_proto_init() {
	if File_protobufs_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobufs_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobufs_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobufs_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobufs_member_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobufs_member_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobufs_member_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobufs_member_proto_goTypes,
		DependencyIndexes: file_protobufs_member_proto_depIdxs,
		EnumInfos:         file_protobufs_member_proto_enumTypes,
		MessageInfos:      file_protobufs_member_proto_msgTypes,
	}.Build()
	File_protobufs_member_proto = out.File
	file_protobufs_member_proto_rawDesc = nil
	file_protobufs_member_proto_goTypes = nil
	file_protobufs_member_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MemberServiceClient interface {
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
}

type memberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberServiceClient(cc grpc.ClientConnInterface) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/member.MemberService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, "/member.MemberService/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
type MemberServiceServer interface {
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
}

// UnimplementedMemberServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMemberServiceServer struct {
}

func (*UnimplementedMemberServiceServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (*UnimplementedMemberServiceServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}

func RegisterMemberServiceServer(s *grpc.Server, srv MemberServiceServer) {
	s.RegisterService(&_MemberService_serviceDesc, srv)
}

func _MemberService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MemberService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "member.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _MemberService_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _MemberService_GetUserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobufs/member.proto",
}
