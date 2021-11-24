// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: api/msg/upload/upload.proto

package upload

import (
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

// 上传类型
type UploadType int32

const (
	UploadType_UploadTypeUnknown UploadType = 0 // 未知
	UploadType_UploadTypeImage   UploadType = 1 // 图片
	UploadType_UploadTypeVoice   UploadType = 2 // 语音
	UploadType_UploadTypeVideo   UploadType = 3 // 视频
	UploadType_UploadTypeFile    UploadType = 4 // 文件
)

// Enum value maps for UploadType.
var (
	UploadType_name = map[int32]string{
		0: "UploadTypeUnknown",
		1: "UploadTypeImage",
		2: "UploadTypeVoice",
		3: "UploadTypeVideo",
		4: "UploadTypeFile",
	}
	UploadType_value = map[string]int32{
		"UploadTypeUnknown": 0,
		"UploadTypeImage":   1,
		"UploadTypeVoice":   2,
		"UploadTypeVideo":   3,
		"UploadTypeFile":    4,
	}
)

func (x UploadType) Enum() *UploadType {
	p := new(UploadType)
	*p = x
	return p
}

func (x UploadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_msg_upload_upload_proto_enumTypes[0].Descriptor()
}

func (UploadType) Type() protoreflect.EnumType {
	return &file_api_msg_upload_upload_proto_enumTypes[0]
}

func (x UploadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadType.Descriptor instead.
func (UploadType) EnumDescriptor() ([]byte, []int) {
	return file_api_msg_upload_upload_proto_rawDescGZIP(), []int{0}
}

// 上传请求（通过字节流）
type UploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type UploadType `protobuf:"varint,1,opt,name=Type,proto3,enum=msg.upload.UploadType" json:"Type,omitempty"` // 上传类型
}

func (x *UploadReq) Reset() {
	*x = UploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_upload_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReq) ProtoMessage() {}

func (x *UploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_upload_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadReq.ProtoReflect.Descriptor instead.
func (*UploadReq) Descriptor() ([]byte, []int) {
	return file_api_msg_upload_upload_proto_rawDescGZIP(), []int{0}
}

func (x *UploadReq) GetType() UploadType {
	if x != nil {
		return x.Type
	}
	return UploadType_UploadTypeUnknown
}

// 上传响应
type UploadRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"` // 内容，比如Image
}

func (x *UploadRsp) Reset() {
	*x = UploadRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_upload_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRsp) ProtoMessage() {}

func (x *UploadRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_upload_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRsp.ProtoReflect.Descriptor instead.
func (*UploadRsp) Descriptor() ([]byte, []int) {
	return file_api_msg_upload_upload_proto_rawDescGZIP(), []int{1}
}

func (x *UploadRsp) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// @存储模型
// 使用对象存储服务存储
type ObjectStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType string `protobuf:"bytes,1,opt,name=ContentType,proto3" json:"ContentType,omitempty"` // HTTP头设置ContentType这样可以直接打开
	URL         string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`                 // URL用于访问
	Content     []byte `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`         // 对象
}

func (x *ObjectStore) Reset() {
	*x = ObjectStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_upload_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStore) ProtoMessage() {}

func (x *ObjectStore) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_upload_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectStore.ProtoReflect.Descriptor instead.
func (*ObjectStore) Descriptor() ([]byte, []int) {
	return file_api_msg_upload_upload_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectStore) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *ObjectStore) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *ObjectStore) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_api_msg_upload_upload_proto protoreflect.FileDescriptor

var file_api_msg_upload_upload_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d,
	0x73, 0x67, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x37, 0x0a, 0x09, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x0b, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52,
	0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x76, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x04, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61,
	0x6f, 0x68, 0x75, 0x61, 0x73, 0x68, 0x69, 0x66, 0x75, 0x2f, 0x68, 0x69, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_msg_upload_upload_proto_rawDescOnce sync.Once
	file_api_msg_upload_upload_proto_rawDescData = file_api_msg_upload_upload_proto_rawDesc
)

func file_api_msg_upload_upload_proto_rawDescGZIP() []byte {
	file_api_msg_upload_upload_proto_rawDescOnce.Do(func() {
		file_api_msg_upload_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_msg_upload_upload_proto_rawDescData)
	})
	return file_api_msg_upload_upload_proto_rawDescData
}

var file_api_msg_upload_upload_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_msg_upload_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_msg_upload_upload_proto_goTypes = []interface{}{
	(UploadType)(0),     // 0: msg.upload.UploadType
	(*UploadReq)(nil),   // 1: msg.upload.UploadReq
	(*UploadRsp)(nil),   // 2: msg.upload.UploadRsp
	(*ObjectStore)(nil), // 3: msg.upload.ObjectStore
}
var file_api_msg_upload_upload_proto_depIdxs = []int32{
	0, // 0: msg.upload.UploadReq.Type:type_name -> msg.upload.UploadType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_msg_upload_upload_proto_init() }
func file_api_msg_upload_upload_proto_init() {
	if File_api_msg_upload_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_msg_upload_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadReq); i {
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
		file_api_msg_upload_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRsp); i {
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
		file_api_msg_upload_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectStore); i {
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
			RawDescriptor: file_api_msg_upload_upload_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_msg_upload_upload_proto_goTypes,
		DependencyIndexes: file_api_msg_upload_upload_proto_depIdxs,
		EnumInfos:         file_api_msg_upload_upload_proto_enumTypes,
		MessageInfos:      file_api_msg_upload_upload_proto_msgTypes,
	}.Build()
	File_api_msg_upload_upload_proto = out.File
	file_api_msg_upload_upload_proto_rawDesc = nil
	file_api_msg_upload_upload_proto_goTypes = nil
	file_api_msg_upload_upload_proto_depIdxs = nil
}
