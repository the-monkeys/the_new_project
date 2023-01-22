// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: services/api_gateway/pkg/blogsandposts/pb/blogsandposts.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBlogReq_Ownership int32

const (
	CreateBlogReq_THE_USER    CreateBlogReq_Ownership = 0
	CreateBlogReq_THE_MONKEYS CreateBlogReq_Ownership = 1
	CreateBlogReq_THE_PARTNER CreateBlogReq_Ownership = 2
)

// Enum value maps for CreateBlogReq_Ownership.
var (
	CreateBlogReq_Ownership_name = map[int32]string{
		0: "THE_USER",
		1: "THE_MONKEYS",
		2: "THE_PARTNER",
	}
	CreateBlogReq_Ownership_value = map[string]int32{
		"THE_USER":    0,
		"THE_MONKEYS": 1,
		"THE_PARTNER": 2,
	}
)

func (x CreateBlogReq_Ownership) Enum() *CreateBlogReq_Ownership {
	p := new(CreateBlogReq_Ownership)
	*p = x
	return p
}

func (x CreateBlogReq_Ownership) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateBlogReq_Ownership) Descriptor() protoreflect.EnumDescriptor {
	return file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_enumTypes[0].Descriptor()
}

func (CreateBlogReq_Ownership) Type() protoreflect.EnumType {
	return &file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_enumTypes[0]
}

func (x CreateBlogReq_Ownership) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateBlogReq_Ownership.Descriptor instead.
func (CreateBlogReq_Ownership) EnumDescriptor() ([]byte, []int) {
	return file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescGZIP(), []int{0, 0}
}

type CreateBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string                  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content    string                  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorName string                  `protobuf:"bytes,4,opt,name=authorName,proto3" json:"authorName,omitempty"`
	AuthorId   string                  `protobuf:"bytes,5,opt,name=authorId,proto3" json:"authorId,omitempty"`
	Published  bool                    `protobuf:"varint,6,opt,name=published,proto3" json:"published,omitempty"`
	Tags       []string                `protobuf:"bytes,7,rep,name=Tags,proto3" json:"Tags,omitempty"`
	Ownership  CreateBlogReq_Ownership `protobuf:"varint,8,opt,name=ownership,proto3,enum=auth.CreateBlogReq_Ownership" json:"ownership,omitempty"`
	CanEdit    bool                    `protobuf:"varint,9,opt,name=canEdit,proto3" json:"canEdit,omitempty"`
	FolderPath string                  `protobuf:"bytes,10,opt,name=folderPath,proto3" json:"folderPath,omitempty"`
	CreateTime *timestamppb.Timestamp  `protobuf:"bytes,11,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime *timestamppb.Timestamp  `protobuf:"bytes,12,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *CreateBlogReq) Reset() {
	*x = CreateBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogReq) ProtoMessage() {}

func (x *CreateBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogReq.ProtoReflect.Descriptor instead.
func (*CreateBlogReq) Descriptor() ([]byte, []int) {
	return file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBlogReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateBlogReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBlogReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateBlogReq) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *CreateBlogReq) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *CreateBlogReq) GetPublished() bool {
	if x != nil {
		return x.Published
	}
	return false
}

func (x *CreateBlogReq) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateBlogReq) GetOwnership() CreateBlogReq_Ownership {
	if x != nil {
		return x.Ownership
	}
	return CreateBlogReq_THE_USER
}

func (x *CreateBlogReq) GetCanEdit() bool {
	if x != nil {
		return x.CanEdit
	}
	return false
}

func (x *CreateBlogReq) GetFolderPath() string {
	if x != nil {
		return x.FolderPath
	}
	return ""
}

func (x *CreateBlogReq) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CreateBlogReq) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type CreateBlogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateBlogRes) Reset() {
	*x = CreateBlogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogRes) ProtoMessage() {}

func (x *CreateBlogRes) ProtoReflect() protoreflect.Message {
	mi := &file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogRes.ProtoReflect.Descriptor instead.
func (*CreateBlogRes) Descriptor() ([]byte, []int) {
	return file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBlogRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateBlogRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateBlogRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto protoreflect.FileDescriptor

var file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73,
	0x61, 0x6e, 0x64, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x73, 0x61, 0x6e, 0x64, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x03, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x2e, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3a, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x09, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x48, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x48, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x4b, 0x45, 0x59, 0x53, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x48, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x4e, 0x45, 0x52,
	0x10, 0x02, 0x22, 0x4f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x50, 0x0a, 0x13, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x41, 0x6e, 0x64, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x61, 0x6e, 0x64, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescOnce sync.Once
	file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescData = file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDesc
)

func file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescGZIP() []byte {
	file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescOnce.Do(func() {
		file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescData)
	})
	return file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDescData
}

var file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_goTypes = []interface{}{
	(CreateBlogReq_Ownership)(0),  // 0: auth.CreateBlogReq.Ownership
	(*CreateBlogReq)(nil),         // 1: auth.CreateBlogReq
	(*CreateBlogRes)(nil),         // 2: auth.CreateBlogRes
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_depIdxs = []int32{
	0, // 0: auth.CreateBlogReq.ownership:type_name -> auth.CreateBlogReq.Ownership
	3, // 1: auth.CreateBlogReq.createTime:type_name -> google.protobuf.Timestamp
	3, // 2: auth.CreateBlogReq.updateTime:type_name -> google.protobuf.Timestamp
	1, // 3: auth.BlogsAndPostService.CreateABlog:input_type -> auth.CreateBlogReq
	2, // 4: auth.BlogsAndPostService.CreateABlog:output_type -> auth.CreateBlogRes
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_init() }
func file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_init() {
	if File_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogReq); i {
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
		file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogRes); i {
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
			RawDescriptor: file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_goTypes,
		DependencyIndexes: file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_depIdxs,
		EnumInfos:         file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_enumTypes,
		MessageInfos:      file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_msgTypes,
	}.Build()
	File_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto = out.File
	file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_rawDesc = nil
	file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_goTypes = nil
	file_services_api_gateway_pkg_blogsandposts_pb_blogsandposts_proto_depIdxs = nil
}
