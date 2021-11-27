// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: api/authnz/authnz.proto

package authnz

import (
	constant "github.com/xiaohuashifu/him/api/constant"
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

// UserType 用户类型
type UserType int32

const (
	UserType_USER_TYPE_UNSPECIFIED UserType = 0 // 未知
	UserType_USER_TYPE_USER        UserType = 1 // 用户
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "USER_TYPE_UNSPECIFIED",
		1: "USER_TYPE_USER",
	}
	UserType_value = map[string]int32{
		"USER_TYPE_UNSPECIFIED": 0,
		"USER_TYPE_USER":        1,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_authnz_authnz_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_api_authnz_authnz_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_api_authnz_authnz_proto_rawDescGZIP(), []int{0}
}

// Session 会话
type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64            `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                            // 用户编号
	UserType UserType          `protobuf:"varint,2,opt,name=user_type,json=userType,proto3,enum=authnz.UserType" json:"user_type,omitempty"` // 用户类型
	Terminal constant.Terminal `protobuf:"varint,3,opt,name=terminal,proto3,enum=constant.Terminal" json:"terminal,omitempty"`               // 终端
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_authnz_authnz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_api_authnz_authnz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_api_authnz_authnz_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Session) GetUserType() UserType {
	if x != nil {
		return x.UserType
	}
	return UserType_USER_TYPE_UNSPECIFIED
}

func (x *Session) GetTerminal() constant.Terminal {
	if x != nil {
		return x.Terminal
	}
	return constant.Terminal(0)
}

var File_api_authnz_authnz_proto protoreflect.FileDescriptor

var file_api_authnz_authnz_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x7a, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x7a, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81,
	0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x7a, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x2a, 0x39, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x6f,
	0x68, 0x75, 0x61, 0x73, 0x68, 0x69, 0x66, 0x75, 0x2f, 0x68, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_authnz_authnz_proto_rawDescOnce sync.Once
	file_api_authnz_authnz_proto_rawDescData = file_api_authnz_authnz_proto_rawDesc
)

func file_api_authnz_authnz_proto_rawDescGZIP() []byte {
	file_api_authnz_authnz_proto_rawDescOnce.Do(func() {
		file_api_authnz_authnz_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_authnz_authnz_proto_rawDescData)
	})
	return file_api_authnz_authnz_proto_rawDescData
}

var file_api_authnz_authnz_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_authnz_authnz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_authnz_authnz_proto_goTypes = []interface{}{
	(UserType)(0),          // 0: authnz.UserType
	(*Session)(nil),        // 1: authnz.Session
	(constant.Terminal)(0), // 2: constant.Terminal
}
var file_api_authnz_authnz_proto_depIdxs = []int32{
	0, // 0: authnz.Session.user_type:type_name -> authnz.UserType
	2, // 1: authnz.Session.terminal:type_name -> constant.Terminal
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_authnz_authnz_proto_init() }
func file_api_authnz_authnz_proto_init() {
	if File_api_authnz_authnz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_authnz_authnz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
			RawDescriptor: file_api_authnz_authnz_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_authnz_authnz_proto_goTypes,
		DependencyIndexes: file_api_authnz_authnz_proto_depIdxs,
		EnumInfos:         file_api_authnz_authnz_proto_enumTypes,
		MessageInfos:      file_api_authnz_authnz_proto_msgTypes,
	}.Build()
	File_api_authnz_authnz_proto = out.File
	file_api_authnz_authnz_proto_rawDesc = nil
	file_api_authnz_authnz_proto_goTypes = nil
	file_api_authnz_authnz_proto_depIdxs = nil
}
