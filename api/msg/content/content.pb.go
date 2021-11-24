// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: api/msg/content/content.proto

package content

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

// 图片类型
type ImageType int32

const (
	ImageType_ImageTypeUnknown ImageType = 0 // 未知
	ImageType_ImageTypeJPG     ImageType = 1 // JPG格式
	ImageType_ImageTypePNG     ImageType = 2 // PNG格式
)

// Enum value maps for ImageType.
var (
	ImageType_name = map[int32]string{
		0: "ImageTypeUnknown",
		1: "ImageTypeJPG",
		2: "ImageTypePNG",
	}
	ImageType_value = map[string]int32{
		"ImageTypeUnknown": 0,
		"ImageTypeJPG":     1,
		"ImageTypePNG":     2,
	}
)

func (x ImageType) Enum() *ImageType {
	p := new(ImageType)
	*p = x
	return p
}

func (x ImageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_msg_content_content_proto_enumTypes[0].Descriptor()
}

func (ImageType) Type() protoreflect.EnumType {
	return &file_api_msg_content_content_proto_enumTypes[0]
}

func (x ImageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageType.Descriptor instead.
func (ImageType) EnumDescriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{0}
}

// 语音类型
type VoiceType int32

const (
	VoiceType_VoiceTypeUnknown VoiceType = 0 // 未知
	VoiceType_VoiceTypeMP3     VoiceType = 1 // MP3格式
)

// Enum value maps for VoiceType.
var (
	VoiceType_name = map[int32]string{
		0: "VoiceTypeUnknown",
		1: "VoiceTypeMP3",
	}
	VoiceType_value = map[string]int32{
		"VoiceTypeUnknown": 0,
		"VoiceTypeMP3":     1,
	}
)

func (x VoiceType) Enum() *VoiceType {
	p := new(VoiceType)
	*p = x
	return p
}

func (x VoiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_msg_content_content_proto_enumTypes[1].Descriptor()
}

func (VoiceType) Type() protoreflect.EnumType {
	return &file_api_msg_content_content_proto_enumTypes[1]
}

func (x VoiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VoiceType.Descriptor instead.
func (VoiceType) EnumDescriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{1}
}

// 视频类型
type VideoType int32

const (
	VideoType_VideoTypeUnknown VideoType = 0 // 未知
	VideoType_VideoTypeMP4     VideoType = 1 // MP4格式
)

// Enum value maps for VideoType.
var (
	VideoType_name = map[int32]string{
		0: "VideoTypeUnknown",
		1: "VideoTypeMP4",
	}
	VideoType_value = map[string]int32{
		"VideoTypeUnknown": 0,
		"VideoTypeMP4":     1,
	}
)

func (x VideoType) Enum() *VideoType {
	p := new(VideoType)
	*p = x
	return p
}

func (x VideoType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VideoType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_msg_content_content_proto_enumTypes[2].Descriptor()
}

func (VideoType) Type() protoreflect.EnumType {
	return &file_api_msg_content_content_proto_enumTypes[2]
}

func (x VideoType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VideoType.Descriptor instead.
func (VideoType) EnumDescriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{2}
}

// 提示类型
type TipType int32

const (
	TipType_TipTypeUnknown      TipType = 0 // 未知
	TipType_TipTypeText         TipType = 1 // 文本类型提示
	TipType_TipTypeNickNameText TipType = 2 // 昵称文本类型提示
)

// Enum value maps for TipType.
var (
	TipType_name = map[int32]string{
		0: "TipTypeUnknown",
		1: "TipTypeText",
		2: "TipTypeNickNameText",
	}
	TipType_value = map[string]int32{
		"TipTypeUnknown":      0,
		"TipTypeText":         1,
		"TipTypeNickNameText": 2,
	}
)

func (x TipType) Enum() *TipType {
	p := new(TipType)
	*p = x
	return p
}

func (x TipType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TipType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_msg_content_content_proto_enumTypes[3].Descriptor()
}

func (TipType) Type() protoreflect.EnumType {
	return &file_api_msg_content_content_proto_enumTypes[3]
}

func (x TipType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TipType.Descriptor instead.
func (TipType) EnumDescriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{3}
}

// 文件类型
type FileType int32

const (
	FileType_FileTypeUnknown FileType = 0 // 未知
	FileType_FileTypePDF     FileType = 1 // PDF格式
)

// Enum value maps for FileType.
var (
	FileType_name = map[int32]string{
		0: "FileTypeUnknown",
		1: "FileTypePDF",
	}
	FileType_value = map[string]int32{
		"FileTypeUnknown": 0,
		"FileTypePDF":     1,
	}
)

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_msg_content_content_proto_enumTypes[4].Descriptor()
}

func (FileType) Type() protoreflect.EnumType {
	return &file_api_msg_content_content_proto_enumTypes[4]
}

func (x FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileType.Descriptor instead.
func (FileType) EnumDescriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{4}
}

// 图片
type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL    string    `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`                           // 地址
	Width  uint32    `protobuf:"varint,2,opt,name=Width,proto3" json:"Width,omitempty"`                      // 宽
	Height uint32    `protobuf:"varint,3,opt,name=Height,proto3" json:"Height,omitempty"`                    // 高
	Type   ImageType `protobuf:"varint,4,opt,name=Type,proto3,enum=message.ImageType" json:"Type,omitempty"` // 图片类型
	Size   uint32    `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`                        // 大小，单位B
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_content_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_content_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *Image) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Image) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Image) GetType() ImageType {
	if x != nil {
		return x.Type
	}
	return ImageType_ImageTypeUnknown
}

func (x *Image) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// 语音
type Voice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL      string    `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`                           // 地址
	Duration uint32    `protobuf:"varint,2,opt,name=Duration,proto3" json:"Duration,omitempty"`                // 时长，单位秒
	Type     VoiceType `protobuf:"varint,3,opt,name=Type,proto3,enum=message.VoiceType" json:"Type,omitempty"` // 语音类型
	Size     uint32    `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`                        // 大小，单位B
}

func (x *Voice) Reset() {
	*x = Voice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_content_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voice) ProtoMessage() {}

func (x *Voice) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_content_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voice.ProtoReflect.Descriptor instead.
func (*Voice) Descriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{1}
}

func (x *Voice) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *Voice) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Voice) GetType() VoiceType {
	if x != nil {
		return x.Type
	}
	return VoiceType_VoiceTypeUnknown
}

func (x *Voice) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// 视频
type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL      string    `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`                           // 地址
	Duration uint32    `protobuf:"varint,2,opt,name=Duration,proto3" json:"Duration,omitempty"`                // 时长，单位秒
	Type     VideoType `protobuf:"varint,3,opt,name=Type,proto3,enum=message.VideoType" json:"Type,omitempty"` // 视频类型
	Size     uint32    `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`                        // 大小，单位B
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_content_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_content_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{2}
}

func (x *Video) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *Video) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Video) GetType() VideoType {
	if x != nil {
		return x.Type
	}
	return VideoType_VideoTypeUnknown
}

func (x *Video) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// 文本提示
type TextTip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"` // 提示内容
}

func (x *TextTip) Reset() {
	*x = TextTip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_content_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextTip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextTip) ProtoMessage() {}

func (x *TextTip) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_content_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextTip.ProtoReflect.Descriptor instead.
func (*TextTip) Descriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{3}
}

func (x *TextTip) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// 昵称文本提示
type NickNameTextTip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format        string          `protobuf:"bytes,1,opt,name=Format,proto3" json:"Format,omitempty"`               // 格式字符串
	NickNameTexts []*NickNameText `protobuf:"bytes,2,rep,name=NickNameTexts,proto3" json:"NickNameTexts,omitempty"` // 昵称文本
}

func (x *NickNameTextTip) Reset() {
	*x = NickNameTextTip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_content_content_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NickNameTextTip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NickNameTextTip) ProtoMessage() {}

func (x *NickNameTextTip) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_content_content_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NickNameTextTip.ProtoReflect.Descriptor instead.
func (*NickNameTextTip) Descriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{4}
}

func (x *NickNameTextTip) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *NickNameTextTip) GetNickNameTexts() []*NickNameText {
	if x != nil {
		return x.NickNameTexts
	}
	return nil
}

// 昵称文本
type NickNameText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`    // 用户编号
	NickName string `protobuf:"bytes,2,opt,name=NickName,proto3" json:"NickName,omitempty"` // 昵称
}

func (x *NickNameText) Reset() {
	*x = NickNameText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_content_content_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NickNameText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NickNameText) ProtoMessage() {}

func (x *NickNameText) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_content_content_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NickNameText.ProtoReflect.Descriptor instead.
func (*NickNameText) Descriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{5}
}

func (x *NickNameText) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *NickNameText) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

// 文件
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL  string   `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`                          // 地址
	Name string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`                        // 文件名
	Type FileType `protobuf:"varint,3,opt,name=Type,proto3,enum=message.FileType" json:"Type,omitempty"` // 文件类型
	Size uint32   `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`                       // 大小，单位B
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_content_content_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_content_content_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_api_msg_content_content_proto_rawDescGZIP(), []int{6}
}

func (x *File) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetType() FileType {
	if x != nil {
		return x.Type
	}
	return FileType_FileTypeUnknown
}

func (x *File) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_api_msg_content_content_proto protoreflect.FileDescriptor

var file_api_msg_content_content_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x05, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x71,
	0x0a, 0x05, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x71, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52,
	0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x54, 0x65, 0x78, 0x74, 0x54, 0x69, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x66, 0x0a, 0x0f, 0x4e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x78, 0x74, 0x54, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x54, 0x65, 0x78, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x0d, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x73, 0x22, 0x42, 0x0a, 0x0c, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x2a, 0x45,
	0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4a, 0x50,
	0x47, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x4e, 0x47, 0x10, 0x02, 0x2a, 0x33, 0x0a, 0x09, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x50, 0x33, 0x10, 0x01, 0x2a, 0x33, 0x0a, 0x09, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x50, 0x34, 0x10, 0x01, 0x2a,
	0x47, 0x0a, 0x07, 0x54, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x69,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x54, 0x65, 0x78, 0x74, 0x10, 0x01, 0x12,
	0x17, 0x0a, 0x13, 0x54, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x54, 0x65, 0x78, 0x74, 0x10, 0x02, 0x2a, 0x30, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x44, 0x46, 0x10, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x6f, 0x68, 0x75, 0x61,
	0x73, 0x68, 0x69, 0x66, 0x75, 0x2f, 0x68, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73,
	0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_msg_content_content_proto_rawDescOnce sync.Once
	file_api_msg_content_content_proto_rawDescData = file_api_msg_content_content_proto_rawDesc
)

func file_api_msg_content_content_proto_rawDescGZIP() []byte {
	file_api_msg_content_content_proto_rawDescOnce.Do(func() {
		file_api_msg_content_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_msg_content_content_proto_rawDescData)
	})
	return file_api_msg_content_content_proto_rawDescData
}

var file_api_msg_content_content_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_api_msg_content_content_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_msg_content_content_proto_goTypes = []interface{}{
	(ImageType)(0),          // 0: message.ImageType
	(VoiceType)(0),          // 1: message.VoiceType
	(VideoType)(0),          // 2: message.VideoType
	(TipType)(0),            // 3: message.TipType
	(FileType)(0),           // 4: message.FileType
	(*Image)(nil),           // 5: message.Image
	(*Voice)(nil),           // 6: message.Voice
	(*Video)(nil),           // 7: message.Video
	(*TextTip)(nil),         // 8: message.TextTip
	(*NickNameTextTip)(nil), // 9: message.NickNameTextTip
	(*NickNameText)(nil),    // 10: message.NickNameText
	(*File)(nil),            // 11: message.File
}
var file_api_msg_content_content_proto_depIdxs = []int32{
	0,  // 0: message.Image.Type:type_name -> message.ImageType
	1,  // 1: message.Voice.Type:type_name -> message.VoiceType
	2,  // 2: message.Video.Type:type_name -> message.VideoType
	10, // 3: message.NickNameTextTip.NickNameTexts:type_name -> message.NickNameText
	4,  // 4: message.File.Type:type_name -> message.FileType
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_msg_content_content_proto_init() }
func file_api_msg_content_content_proto_init() {
	if File_api_msg_content_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_msg_content_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_api_msg_content_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Voice); i {
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
		file_api_msg_content_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_api_msg_content_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextTip); i {
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
		file_api_msg_content_content_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NickNameTextTip); i {
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
		file_api_msg_content_content_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NickNameText); i {
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
		file_api_msg_content_content_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_api_msg_content_content_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_msg_content_content_proto_goTypes,
		DependencyIndexes: file_api_msg_content_content_proto_depIdxs,
		EnumInfos:         file_api_msg_content_content_proto_enumTypes,
		MessageInfos:      file_api_msg_content_content_proto_msgTypes,
	}.Build()
	File_api_msg_content_content_proto = out.File
	file_api_msg_content_content_proto_rawDesc = nil
	file_api_msg_content_content_proto_goTypes = nil
	file_api_msg_content_content_proto_depIdxs = nil
}
