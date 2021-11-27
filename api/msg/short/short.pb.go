// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: api/msg/short/short.proto

package short

import (
	msg "github.com/xiaohuashifu/him/api/msg"
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
type UploadReq_Type int32

const (
	UploadReq_TYPE_UNSPECIFIED UploadReq_Type = 0 // 未知
	UploadReq_TYPE_IMAGE       UploadReq_Type = 1 // 图片
	UploadReq_TYPE_VOICE       UploadReq_Type = 2 // 语音
	UploadReq_TYPE_VIDEO       UploadReq_Type = 3 // 视频
	UploadReq_TYPE_FILE        UploadReq_Type = 4 // 文件
)

// Enum value maps for UploadReq_Type.
var (
	UploadReq_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_IMAGE",
		2: "TYPE_VOICE",
		3: "TYPE_VIDEO",
		4: "TYPE_FILE",
	}
	UploadReq_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_IMAGE":       1,
		"TYPE_VOICE":       2,
		"TYPE_VIDEO":       3,
		"TYPE_FILE":        4,
	}
)

func (x UploadReq_Type) Enum() *UploadReq_Type {
	p := new(UploadReq_Type)
	*p = x
	return p
}

func (x UploadReq_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadReq_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_msg_short_short_proto_enumTypes[0].Descriptor()
}

func (UploadReq_Type) Type() protoreflect.EnumType {
	return &file_api_msg_short_short_proto_enumTypes[0]
}

func (x UploadReq_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadReq_Type.Descriptor instead.
func (UploadReq_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_msg_short_short_proto_rawDescGZIP(), []int{0, 0}
}

// 上传请求（通过字节流）
type UploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type UploadReq_Type `protobuf:"varint,1,opt,name=type,proto3,enum=msg.short.UploadReq_Type" json:"type,omitempty"` // 上传类型
}

func (x *UploadReq) Reset() {
	*x = UploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_short_short_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReq) ProtoMessage() {}

func (x *UploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_short_short_proto_msgTypes[0]
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
	return file_api_msg_short_short_proto_rawDescGZIP(), []int{0}
}

func (x *UploadReq) GetType() UploadReq_Type {
	if x != nil {
		return x.Type
	}
	return UploadReq_TYPE_UNSPECIFIED
}

// 上传响应
type UploadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 内容
	//
	// Types that are assignable to Content:
	//	*UploadResp_Image
	//	*UploadResp_Voice
	//	*UploadResp_Video
	//	*UploadResp_File
	Content isUploadResp_Content `protobuf_oneof:"Content"`
}

func (x *UploadResp) Reset() {
	*x = UploadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_short_short_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResp) ProtoMessage() {}

func (x *UploadResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_short_short_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResp.ProtoReflect.Descriptor instead.
func (*UploadResp) Descriptor() ([]byte, []int) {
	return file_api_msg_short_short_proto_rawDescGZIP(), []int{1}
}

func (m *UploadResp) GetContent() isUploadResp_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *UploadResp) GetImage() *msg.Image {
	if x, ok := x.GetContent().(*UploadResp_Image); ok {
		return x.Image
	}
	return nil
}

func (x *UploadResp) GetVoice() *msg.Voice {
	if x, ok := x.GetContent().(*UploadResp_Voice); ok {
		return x.Voice
	}
	return nil
}

func (x *UploadResp) GetVideo() *msg.Video {
	if x, ok := x.GetContent().(*UploadResp_Video); ok {
		return x.Video
	}
	return nil
}

func (x *UploadResp) GetFile() *msg.File {
	if x, ok := x.GetContent().(*UploadResp_File); ok {
		return x.File
	}
	return nil
}

type isUploadResp_Content interface {
	isUploadResp_Content()
}

type UploadResp_Image struct {
	Image *msg.Image `protobuf:"bytes,1,opt,name=image,proto3,oneof"` // 图片
}

type UploadResp_Voice struct {
	Voice *msg.Voice `protobuf:"bytes,2,opt,name=voice,proto3,oneof"` // 语音
}

type UploadResp_Video struct {
	Video *msg.Video `protobuf:"bytes,3,opt,name=video,proto3,oneof"` // 视频
}

type UploadResp_File struct {
	File *msg.File `protobuf:"bytes,4,opt,name=file,proto3,oneof"` // 文件
}

func (*UploadResp_Image) isUploadResp_Content() {}

func (*UploadResp_Voice) isUploadResp_Content() {}

func (*UploadResp_Video) isUploadResp_Content() {}

func (*UploadResp_File) isUploadResp_Content() {}

// 获取最后一个信箱消息编号请求
type GetLastMailBoxMsgIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户编号
}

func (x *GetLastMailBoxMsgIdReq) Reset() {
	*x = GetLastMailBoxMsgIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_short_short_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastMailBoxMsgIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastMailBoxMsgIdReq) ProtoMessage() {}

func (x *GetLastMailBoxMsgIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_short_short_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastMailBoxMsgIdReq.ProtoReflect.Descriptor instead.
func (*GetLastMailBoxMsgIdReq) Descriptor() ([]byte, []int) {
	return file_api_msg_short_short_proto_rawDescGZIP(), []int{2}
}

func (x *GetLastMailBoxMsgIdReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 获取最后一个信箱消息编号响应
type GetLastMailBoxMsgIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastMailBoxMsgId uint64 `protobuf:"varint,1,opt,name=last_mail_box_msg_id,json=lastMailBoxMsgId,proto3" json:"last_mail_box_msg_id,omitempty"` // 最后一个信箱消息编号
}

func (x *GetLastMailBoxMsgIdResp) Reset() {
	*x = GetLastMailBoxMsgIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_short_short_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastMailBoxMsgIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastMailBoxMsgIdResp) ProtoMessage() {}

func (x *GetLastMailBoxMsgIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_short_short_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastMailBoxMsgIdResp.ProtoReflect.Descriptor instead.
func (*GetLastMailBoxMsgIdResp) Descriptor() ([]byte, []int) {
	return file_api_msg_short_short_proto_rawDescGZIP(), []int{3}
}

func (x *GetLastMailBoxMsgIdResp) GetLastMailBoxMsgId() uint64 {
	if x != nil {
		return x.LastMailBoxMsgId
	}
	return 0
}

// 获取消息请求
type GetMsgsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailBoxMsgIdRanges []*GetMsgsReq_MailBoxMsgIdRange `protobuf:"bytes,1,rep,name=mail_box_msg_id_ranges,json=mailBoxMsgIdRanges,proto3" json:"mail_box_msg_id_ranges,omitempty"` // 信箱消息编号范围
}

func (x *GetMsgsReq) Reset() {
	*x = GetMsgsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_short_short_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMsgsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMsgsReq) ProtoMessage() {}

func (x *GetMsgsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_short_short_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMsgsReq.ProtoReflect.Descriptor instead.
func (*GetMsgsReq) Descriptor() ([]byte, []int) {
	return file_api_msg_short_short_proto_rawDescGZIP(), []int{4}
}

func (x *GetMsgsReq) GetMailBoxMsgIdRanges() []*GetMsgsReq_MailBoxMsgIdRange {
	if x != nil {
		return x.MailBoxMsgIdRanges
	}
	return nil
}

// 获取消息响应
type GetMsgsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgs []*msg.Msg `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"` // 消息
}

func (x *GetMsgsResp) Reset() {
	*x = GetMsgsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_short_short_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMsgsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMsgsResp) ProtoMessage() {}

func (x *GetMsgsResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_short_short_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMsgsResp.ProtoReflect.Descriptor instead.
func (*GetMsgsResp) Descriptor() ([]byte, []int) {
	return file_api_msg_short_short_proto_rawDescGZIP(), []int{5}
}

func (x *GetMsgsResp) GetMsgs() []*msg.Msg {
	if x != nil {
		return x.Msgs
	}
	return nil
}

// @存储模型
// 使用对象存储服务存储
type ObjectStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"` // HTTP头设置ContentType这样可以直接打开
	Url         string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`                                    // URL用于访问
	Content     []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`                            // 对象
}

func (x *ObjectStore) Reset() {
	*x = ObjectStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_short_short_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectStore) ProtoMessage() {}

func (x *ObjectStore) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_short_short_proto_msgTypes[6]
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
	return file_api_msg_short_short_proto_rawDescGZIP(), []int{6}
}

func (x *ObjectStore) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *ObjectStore) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ObjectStore) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// 信箱消息编号范围
type GetMsgsReq_MailBoxMsgIdRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartId uint64 `protobuf:"varint,1,opt,name=start_id,json=startId,proto3" json:"start_id,omitempty"` // 起始编号
	EndId   uint64 `protobuf:"varint,2,opt,name=end_id,json=endId,proto3" json:"end_id,omitempty"`       // 结尾编号
}

func (x *GetMsgsReq_MailBoxMsgIdRange) Reset() {
	*x = GetMsgsReq_MailBoxMsgIdRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_short_short_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMsgsReq_MailBoxMsgIdRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMsgsReq_MailBoxMsgIdRange) ProtoMessage() {}

func (x *GetMsgsReq_MailBoxMsgIdRange) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_short_short_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMsgsReq_MailBoxMsgIdRange.ProtoReflect.Descriptor instead.
func (*GetMsgsReq_MailBoxMsgIdRange) Descriptor() ([]byte, []int) {
	return file_api_msg_short_short_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetMsgsReq_MailBoxMsgIdRange) GetStartId() uint64 {
	if x != nil {
		return x.StartId
	}
	return 0
}

func (x *GetMsgsReq_MailBoxMsgIdRange) GetEndId() uint64 {
	if x != nil {
		return x.EndId
	}
	return 0
}

var File_api_msg_short_short_proto protoreflect.FileDescriptor

var file_api_msg_short_short_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2f,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x73, 0x67,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x1a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f,
	0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x09, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x4f, 0x49,
	0x43, 0x45, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x44,
	0x45, 0x4f, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c,
	0x45, 0x10, 0x04, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x48, 0x00, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1f,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x4d, 0x73, 0x67, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x49, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x4d,
	0x73, 0x67, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x61, 0x69, 0x6c,
	0x42, 0x6f, 0x78, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x4d, 0x73, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12, 0x5b, 0x0a, 0x16, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x62, 0x6f, 0x78, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x4d,
	0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x12, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x1a, 0x45, 0x0a, 0x11, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x4d,
	0x73, 0x67, 0x49, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x4d, 0x73, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x04, 0x6d, 0x73,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d,
	0x73, 0x67, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x22, 0x5c, 0x0a, 0x0b, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x6f, 0x68, 0x75, 0x61, 0x73, 0x68, 0x69, 0x66,
	0x75, 0x2f, 0x68, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_msg_short_short_proto_rawDescOnce sync.Once
	file_api_msg_short_short_proto_rawDescData = file_api_msg_short_short_proto_rawDesc
)

func file_api_msg_short_short_proto_rawDescGZIP() []byte {
	file_api_msg_short_short_proto_rawDescOnce.Do(func() {
		file_api_msg_short_short_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_msg_short_short_proto_rawDescData)
	})
	return file_api_msg_short_short_proto_rawDescData
}

var file_api_msg_short_short_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_msg_short_short_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_msg_short_short_proto_goTypes = []interface{}{
	(UploadReq_Type)(0),                  // 0: msg.short.UploadReq.Type
	(*UploadReq)(nil),                    // 1: msg.short.UploadReq
	(*UploadResp)(nil),                   // 2: msg.short.UploadResp
	(*GetLastMailBoxMsgIdReq)(nil),       // 3: msg.short.GetLastMailBoxMsgIdReq
	(*GetLastMailBoxMsgIdResp)(nil),      // 4: msg.short.GetLastMailBoxMsgIdResp
	(*GetMsgsReq)(nil),                   // 5: msg.short.GetMsgsReq
	(*GetMsgsResp)(nil),                  // 6: msg.short.GetMsgsResp
	(*ObjectStore)(nil),                  // 7: msg.short.ObjectStore
	(*GetMsgsReq_MailBoxMsgIdRange)(nil), // 8: msg.short.GetMsgsReq.MailBoxMsgIdRange
	(*msg.Image)(nil),                    // 9: msg.Image
	(*msg.Voice)(nil),                    // 10: msg.Voice
	(*msg.Video)(nil),                    // 11: msg.Video
	(*msg.File)(nil),                     // 12: msg.File
	(*msg.Msg)(nil),                      // 13: msg.Msg
}
var file_api_msg_short_short_proto_depIdxs = []int32{
	0,  // 0: msg.short.UploadReq.type:type_name -> msg.short.UploadReq.Type
	9,  // 1: msg.short.UploadResp.image:type_name -> msg.Image
	10, // 2: msg.short.UploadResp.voice:type_name -> msg.Voice
	11, // 3: msg.short.UploadResp.video:type_name -> msg.Video
	12, // 4: msg.short.UploadResp.file:type_name -> msg.File
	8,  // 5: msg.short.GetMsgsReq.mail_box_msg_id_ranges:type_name -> msg.short.GetMsgsReq.MailBoxMsgIdRange
	13, // 6: msg.short.GetMsgsResp.msgs:type_name -> msg.Msg
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_msg_short_short_proto_init() }
func file_api_msg_short_short_proto_init() {
	if File_api_msg_short_short_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_msg_short_short_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_msg_short_short_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResp); i {
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
		file_api_msg_short_short_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastMailBoxMsgIdReq); i {
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
		file_api_msg_short_short_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastMailBoxMsgIdResp); i {
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
		file_api_msg_short_short_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMsgsReq); i {
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
		file_api_msg_short_short_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMsgsResp); i {
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
		file_api_msg_short_short_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_msg_short_short_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMsgsReq_MailBoxMsgIdRange); i {
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
	file_api_msg_short_short_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UploadResp_Image)(nil),
		(*UploadResp_Voice)(nil),
		(*UploadResp_Video)(nil),
		(*UploadResp_File)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_msg_short_short_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_msg_short_short_proto_goTypes,
		DependencyIndexes: file_api_msg_short_short_proto_depIdxs,
		EnumInfos:         file_api_msg_short_short_proto_enumTypes,
		MessageInfos:      file_api_msg_short_short_proto_msgTypes,
	}.Build()
	File_api_msg_short_short_proto = out.File
	file_api_msg_short_short_proto_rawDesc = nil
	file_api_msg_short_short_proto_goTypes = nil
	file_api_msg_short_short_proto_depIdxs = nil
}
