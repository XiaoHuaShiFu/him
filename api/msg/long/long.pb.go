// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: api/msg/long/long.proto

package long

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

// 登录请求
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // 凭证
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_long_long_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_long_long_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_api_msg_long_long_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 登录响应
type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_long_long_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_long_long_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_api_msg_long_long_proto_rawDescGZIP(), []int{1}
}

// 发送消息请求
type SendMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender        *msg.Sender   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`                                     // 发送者
	Receiver      *msg.Receiver `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`                                 // 接收者
	SendTime      uint64        `protobuf:"varint,3,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`                // 发送时间
	CorrelationId uint64        `protobuf:"varint,4,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"` // 消息请求唯一标识
	// 消息内容
	//
	// Types that are assignable to Content:
	//	*SendMsgReq_TextMsg
	//	*SendMsgReq_ImageMsg
	//	*SendMsgReq_VoiceMsg
	//	*SendMsgReq_VideoMsg
	//	*SendMsgReq_FileMsg
	//	*SendMsgReq_WithdrawMsg
	//	*SendMsgReq_TipMsg
	//	*SendMsgReq_EventMsg
	Content isSendMsgReq_Content `protobuf_oneof:"content"`
}

func (x *SendMsgReq) Reset() {
	*x = SendMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_long_long_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgReq) ProtoMessage() {}

func (x *SendMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_long_long_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgReq.ProtoReflect.Descriptor instead.
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return file_api_msg_long_long_proto_rawDescGZIP(), []int{2}
}

func (x *SendMsgReq) GetSender() *msg.Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *SendMsgReq) GetReceiver() *msg.Receiver {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *SendMsgReq) GetSendTime() uint64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *SendMsgReq) GetCorrelationId() uint64 {
	if x != nil {
		return x.CorrelationId
	}
	return 0
}

func (m *SendMsgReq) GetContent() isSendMsgReq_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *SendMsgReq) GetTextMsg() *msg.TextMsg {
	if x, ok := x.GetContent().(*SendMsgReq_TextMsg); ok {
		return x.TextMsg
	}
	return nil
}

func (x *SendMsgReq) GetImageMsg() *msg.ImageMsg {
	if x, ok := x.GetContent().(*SendMsgReq_ImageMsg); ok {
		return x.ImageMsg
	}
	return nil
}

func (x *SendMsgReq) GetVoiceMsg() *msg.VoiceMsg {
	if x, ok := x.GetContent().(*SendMsgReq_VoiceMsg); ok {
		return x.VoiceMsg
	}
	return nil
}

func (x *SendMsgReq) GetVideoMsg() *msg.VideoMsg {
	if x, ok := x.GetContent().(*SendMsgReq_VideoMsg); ok {
		return x.VideoMsg
	}
	return nil
}

func (x *SendMsgReq) GetFileMsg() *msg.FileMsg {
	if x, ok := x.GetContent().(*SendMsgReq_FileMsg); ok {
		return x.FileMsg
	}
	return nil
}

func (x *SendMsgReq) GetWithdrawMsg() *msg.WithdrawMsg {
	if x, ok := x.GetContent().(*SendMsgReq_WithdrawMsg); ok {
		return x.WithdrawMsg
	}
	return nil
}

func (x *SendMsgReq) GetTipMsg() *msg.TipMsg {
	if x, ok := x.GetContent().(*SendMsgReq_TipMsg); ok {
		return x.TipMsg
	}
	return nil
}

func (x *SendMsgReq) GetEventMsg() *msg.EventMsg {
	if x, ok := x.GetContent().(*SendMsgReq_EventMsg); ok {
		return x.EventMsg
	}
	return nil
}

type isSendMsgReq_Content interface {
	isSendMsgReq_Content()
}

type SendMsgReq_TextMsg struct {
	TextMsg *msg.TextMsg `protobuf:"bytes,5,opt,name=text_msg,json=textMsg,proto3,oneof"`
}

type SendMsgReq_ImageMsg struct {
	ImageMsg *msg.ImageMsg `protobuf:"bytes,6,opt,name=image_msg,json=imageMsg,proto3,oneof"`
}

type SendMsgReq_VoiceMsg struct {
	VoiceMsg *msg.VoiceMsg `protobuf:"bytes,7,opt,name=voice_msg,json=voiceMsg,proto3,oneof"`
}

type SendMsgReq_VideoMsg struct {
	VideoMsg *msg.VideoMsg `protobuf:"bytes,8,opt,name=video_msg,json=videoMsg,proto3,oneof"`
}

type SendMsgReq_FileMsg struct {
	FileMsg *msg.FileMsg `protobuf:"bytes,9,opt,name=file_msg,json=fileMsg,proto3,oneof"`
}

type SendMsgReq_WithdrawMsg struct {
	WithdrawMsg *msg.WithdrawMsg `protobuf:"bytes,10,opt,name=withdraw_msg,json=withdrawMsg,proto3,oneof"`
}

type SendMsgReq_TipMsg struct {
	TipMsg *msg.TipMsg `protobuf:"bytes,11,opt,name=tip_msg,json=tipMsg,proto3,oneof"`
}

type SendMsgReq_EventMsg struct {
	EventMsg *msg.EventMsg `protobuf:"bytes,12,opt,name=event_msg,json=eventMsg,proto3,oneof"`
}

func (*SendMsgReq_TextMsg) isSendMsgReq_Content() {}

func (*SendMsgReq_ImageMsg) isSendMsgReq_Content() {}

func (*SendMsgReq_VoiceMsg) isSendMsgReq_Content() {}

func (*SendMsgReq_VideoMsg) isSendMsgReq_Content() {}

func (*SendMsgReq_FileMsg) isSendMsgReq_Content() {}

func (*SendMsgReq_WithdrawMsg) isSendMsgReq_Content() {}

func (*SendMsgReq_TipMsg) isSendMsgReq_Content() {}

func (*SendMsgReq_EventMsg) isSendMsgReq_Content() {}

// 发送消息响应
type SendMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId uint64 `protobuf:"varint,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`  // 消息请求唯一标识
	MailBoxMsgId  uint64 `protobuf:"varint,2,opt,name=mail_box_msg_id,json=mailBoxMsgId,proto3" json:"mail_box_msg_id,omitempty"` // 信箱里消息编号
	MsgId         uint64 `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`                          // 消息编号，递增
}

func (x *SendMsgResp) Reset() {
	*x = SendMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_long_long_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgResp) ProtoMessage() {}

func (x *SendMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_long_long_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgResp.ProtoReflect.Descriptor instead.
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return file_api_msg_long_long_proto_rawDescGZIP(), []int{3}
}

func (x *SendMsgResp) GetCorrelationId() uint64 {
	if x != nil {
		return x.CorrelationId
	}
	return 0
}

func (x *SendMsgResp) GetMailBoxMsgId() uint64 {
	if x != nil {
		return x.MailBoxMsgId
	}
	return 0
}

func (x *SendMsgResp) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

var File_api_msg_long_long_proto protoreflect.FileDescriptor

var file_api_msg_long_long_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x2f, 0x6c,
	0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x73, 0x67, 0x2e, 0x6c,
	0x6f, 0x6e, 0x67, 0x1a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0b, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x98, 0x04, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x07, 0x74, 0x65, 0x78, 0x74,
	0x4d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x73,
	0x67, 0x12, 0x2c, 0x0a, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x08, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x12,
	0x2c, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x73,
	0x67, 0x48, 0x00, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x73, 0x67, 0x12, 0x29, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x35, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x4d, 0x73, 0x67,
	0x48, 0x00, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x4d, 0x73, 0x67, 0x12,
	0x26, 0x0a, 0x07, 0x74, 0x69, 0x70, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x69, 0x70, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52,
	0x06, 0x74, 0x69, 0x70, 0x4d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x4d, 0x73, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x72, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x62,
	0x6f, 0x78, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d,
	0x73, 0x67, 0x49, 0x64, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x6f, 0x68, 0x75, 0x61, 0x73, 0x68, 0x69, 0x66, 0x75, 0x2f,
	0x68, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x6c, 0x6f, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_msg_long_long_proto_rawDescOnce sync.Once
	file_api_msg_long_long_proto_rawDescData = file_api_msg_long_long_proto_rawDesc
)

func file_api_msg_long_long_proto_rawDescGZIP() []byte {
	file_api_msg_long_long_proto_rawDescOnce.Do(func() {
		file_api_msg_long_long_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_msg_long_long_proto_rawDescData)
	})
	return file_api_msg_long_long_proto_rawDescData
}

var file_api_msg_long_long_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_msg_long_long_proto_goTypes = []interface{}{
	(*LoginReq)(nil),        // 0: msg.long.LoginReq
	(*LoginResp)(nil),       // 1: msg.long.LoginResp
	(*SendMsgReq)(nil),      // 2: msg.long.SendMsgReq
	(*SendMsgResp)(nil),     // 3: msg.long.SendMsgResp
	(*msg.Sender)(nil),      // 4: msg.Sender
	(*msg.Receiver)(nil),    // 5: msg.Receiver
	(*msg.TextMsg)(nil),     // 6: msg.TextMsg
	(*msg.ImageMsg)(nil),    // 7: msg.ImageMsg
	(*msg.VoiceMsg)(nil),    // 8: msg.VoiceMsg
	(*msg.VideoMsg)(nil),    // 9: msg.VideoMsg
	(*msg.FileMsg)(nil),     // 10: msg.FileMsg
	(*msg.WithdrawMsg)(nil), // 11: msg.WithdrawMsg
	(*msg.TipMsg)(nil),      // 12: msg.TipMsg
	(*msg.EventMsg)(nil),    // 13: msg.EventMsg
}
var file_api_msg_long_long_proto_depIdxs = []int32{
	4,  // 0: msg.long.SendMsgReq.sender:type_name -> msg.Sender
	5,  // 1: msg.long.SendMsgReq.receiver:type_name -> msg.Receiver
	6,  // 2: msg.long.SendMsgReq.text_msg:type_name -> msg.TextMsg
	7,  // 3: msg.long.SendMsgReq.image_msg:type_name -> msg.ImageMsg
	8,  // 4: msg.long.SendMsgReq.voice_msg:type_name -> msg.VoiceMsg
	9,  // 5: msg.long.SendMsgReq.video_msg:type_name -> msg.VideoMsg
	10, // 6: msg.long.SendMsgReq.file_msg:type_name -> msg.FileMsg
	11, // 7: msg.long.SendMsgReq.withdraw_msg:type_name -> msg.WithdrawMsg
	12, // 8: msg.long.SendMsgReq.tip_msg:type_name -> msg.TipMsg
	13, // 9: msg.long.SendMsgReq.event_msg:type_name -> msg.EventMsg
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_msg_long_long_proto_init() }
func file_api_msg_long_long_proto_init() {
	if File_api_msg_long_long_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_msg_long_long_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_api_msg_long_long_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_api_msg_long_long_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgReq); i {
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
		file_api_msg_long_long_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgResp); i {
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
	file_api_msg_long_long_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SendMsgReq_TextMsg)(nil),
		(*SendMsgReq_ImageMsg)(nil),
		(*SendMsgReq_VoiceMsg)(nil),
		(*SendMsgReq_VideoMsg)(nil),
		(*SendMsgReq_FileMsg)(nil),
		(*SendMsgReq_WithdrawMsg)(nil),
		(*SendMsgReq_TipMsg)(nil),
		(*SendMsgReq_EventMsg)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_msg_long_long_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_msg_long_long_proto_goTypes,
		DependencyIndexes: file_api_msg_long_long_proto_depIdxs,
		MessageInfos:      file_api_msg_long_long_proto_msgTypes,
	}.Build()
	File_api_msg_long_long_proto = out.File
	file_api_msg_long_long_proto_rawDesc = nil
	file_api_msg_long_long_proto_goTypes = nil
	file_api_msg_long_long_proto_depIdxs = nil
}
