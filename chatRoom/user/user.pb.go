// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.6
// source: user.proto

package user

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

//注册信息
type UserRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"` //用户昵称
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` //密码
}

func (x *UserRegister) Reset() {
	*x = UserRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegister) ProtoMessage() {}

func (x *UserRegister) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegister.ProtoReflect.Descriptor instead.
func (*UserRegister) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegister) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserRegister) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`    //用户账号
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` //用户密码
}

func (x *UserLogin) Reset() {
	*x = UserLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogin) ProtoMessage() {}

func (x *UserLogin) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogin.ProtoReflect.Descriptor instead.
func (*UserLogin) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserLogin) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserLogin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`    //用户账号
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"` //用户昵称
	Status   int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`    //用户状态 0未登录状态 1登录状态
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserDetail) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserDetail) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserDetail) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UserLogout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"` //用户账号
}

func (x *UserLogout) Reset() {
	*x = UserLogout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLogout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogout) ProtoMessage() {}

func (x *UserLogout) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogout.ProtoReflect.Descriptor instead.
func (*UserLogout) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserLogout) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Needusers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Pagenums int32 `protobuf:"varint,2,opt,name=pagenums,proto3" json:"pagenums,omitempty"`
}

func (x *Needusers) Reset() {
	*x = Needusers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Needusers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Needusers) ProtoMessage() {}

func (x *Needusers) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Needusers.ProtoReflect.Descriptor instead.
func (*Needusers) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *Needusers) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Needusers) GetPagenums() int32 {
	if x != nil {
		return x.Pagenums
	}
	return 0
}

type UserInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Users []*UserDetail `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserInfos) Reset() {
	*x = UserInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfos) ProtoMessage() {}

func (x *UserInfos) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfos.ProtoReflect.Descriptor instead.
func (*UserInfos) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfos) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserInfos) GetUsers() []*UserDetail {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x3f, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x58, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x24, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x09, 0x4e, 0x65, 0x65, 0x64, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x22, 0x44, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xa2, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x12, 0x0d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x1a, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x21,
	0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x0a, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x22, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x0b, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x1a, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x0a, 0x2e, 0x4e, 0x65, 0x65, 0x64, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x1a, 0x0a, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_proto_goTypes = []interface{}{
	(*UserRegister)(nil), // 0: UserRegister
	(*UserLogin)(nil),    // 1: UserLogin
	(*UserDetail)(nil),   // 2: UserDetail
	(*UserLogout)(nil),   // 3: UserLogout
	(*Needusers)(nil),    // 4: Needusers
	(*UserInfos)(nil),    // 5: UserInfos
}
var file_user_proto_depIdxs = []int32{
	2, // 0: UserInfos.users:type_name -> UserDetail
	0, // 1: UserService.SignUp:input_type -> UserRegister
	1, // 2: UserService.SignIn:input_type -> UserLogin
	3, // 3: UserService.Logout:input_type -> UserLogout
	4, // 4: UserService.GetUserInfos:input_type -> Needusers
	2, // 5: UserService.SignUp:output_type -> UserDetail
	2, // 6: UserService.SignIn:output_type -> UserDetail
	2, // 7: UserService.Logout:output_type -> UserDetail
	5, // 8: UserService.GetUserInfos:output_type -> UserInfos
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegister); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLogin); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLogout); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Needusers); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfos); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	SignUp(ctx context.Context, in *UserRegister, opts ...grpc.CallOption) (*UserDetail, error)
	SignIn(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserDetail, error)
	Logout(ctx context.Context, in *UserLogout, opts ...grpc.CallOption) (*UserDetail, error)
	GetUserInfos(ctx context.Context, in *Needusers, opts ...grpc.CallOption) (*UserInfos, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) SignUp(ctx context.Context, in *UserRegister, opts ...grpc.CallOption) (*UserDetail, error) {
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, "/UserService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignIn(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserDetail, error) {
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, "/UserService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Logout(ctx context.Context, in *UserLogout, opts ...grpc.CallOption) (*UserDetail, error) {
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, "/UserService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfos(ctx context.Context, in *Needusers, opts ...grpc.CallOption) (*UserInfos, error) {
	out := new(UserInfos)
	err := c.cc.Invoke(ctx, "/UserService/GetUserInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	SignUp(context.Context, *UserRegister) (*UserDetail, error)
	SignIn(context.Context, *UserLogin) (*UserDetail, error)
	Logout(context.Context, *UserLogout) (*UserDetail, error)
	GetUserInfos(context.Context, *Needusers) (*UserInfos, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) SignUp(context.Context, *UserRegister) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedUserServiceServer) SignIn(context.Context, *UserLogin) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedUserServiceServer) Logout(context.Context, *UserLogout) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedUserServiceServer) GetUserInfos(context.Context, *Needusers) (*UserInfos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfos not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignUp(ctx, req.(*UserRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignIn(ctx, req.(*UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Logout(ctx, req.(*UserLogout))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Needusers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/GetUserInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfos(ctx, req.(*Needusers))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _UserService_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _UserService_SignIn_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserService_Logout_Handler,
		},
		{
			MethodName: "GetUserInfos",
			Handler:    _UserService_GetUserInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}