// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MsgType int32

const (
	MsgType_MT_UNKNOWN MsgType = 0
	MsgType_MT_WORDS   MsgType = 1
	MsgType_MT_EMOJI   MsgType = 2
	MsgType_MT_VOICE   MsgType = 3
	MsgType_MT_IMAGE   MsgType = 4
	MsgType_MT_FILE    MsgType = 5
	MsgType_MT_GEO     MsgType = 6
	MsgType_MT_CUSTOM  MsgType = 7
)

var MsgType_name = map[int32]string{
	0: "MT_UNKNOWN",
	1: "MT_WORDS",
	2: "MT_EMOJI",
	3: "MT_VOICE",
	4: "MT_IMAGE",
	5: "MT_FILE",
	6: "MT_GEO",
	7: "MT_CUSTOM",
}

var MsgType_value = map[string]int32{
	"MT_UNKNOWN": 0,
	"MT_WORDS":   1,
	"MT_EMOJI":   2,
	"MT_VOICE":   3,
	"MT_IMAGE":   4,
	"MT_FILE":    5,
	"MT_GEO":     6,
	"MT_CUSTOM":  7,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

type MsgStatus int32

const (
	MsgStatus_MS_UNKNOWN  MsgStatus = 0
	MsgStatus_MS_NORMAL   MsgStatus = 1
	MsgStatus_MS_WITHDRAW MsgStatus = 2
)

var MsgStatus_name = map[int32]string{
	0: "MS_UNKNOWN",
	1: "MS_NORMAL",
	2: "MS_WITHDRAW",
}

var MsgStatus_value = map[string]int32{
	"MS_UNKNOWN":  0,
	"MS_NORMAL":   1,
	"MS_WITHDRAW": 2,
}

func (x MsgStatus) String() string {
	return proto.EnumName(MsgStatus_name, int32(x))
}

func (MsgStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

type Message struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SendId               int64     `protobuf:"varint,2,opt,name=sendId,proto3" json:"sendId,omitempty"`
	SessionId            int64     `protobuf:"varint,3,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Time                 int64     `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	RequestId            int64     `protobuf:"varint,5,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Status               MsgStatus `protobuf:"varint,6,opt,name=status,proto3,enum=rpc.MsgStatus" json:"status,omitempty"`
	Type                 MsgType   `protobuf:"varint,7,opt,name=type,proto3,enum=rpc.MsgType" json:"type,omitempty"`
	Content              []byte    `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
	SequenceNo           int64     `protobuf:"varint,9,opt,name=sequenceNo,proto3" json:"sequenceNo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message) GetSendId() int64 {
	if m != nil {
		return m.SendId
	}
	return 0
}

func (m *Message) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Message) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Message) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Message) GetStatus() MsgStatus {
	if m != nil {
		return m.Status
	}
	return MsgStatus_MS_UNKNOWN
}

func (m *Message) GetType() MsgType {
	if m != nil {
		return m.Type
	}
	return MsgType_MT_UNKNOWN
}

func (m *Message) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Message) GetSequenceNo() int64 {
	if m != nil {
		return m.SequenceNo
	}
	return 0
}

type Words struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Words) Reset()         { *m = Words{} }
func (m *Words) String() string { return proto.CompactTextString(m) }
func (*Words) ProtoMessage()    {}
func (*Words) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Words) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Words.Unmarshal(m, b)
}
func (m *Words) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Words.Marshal(b, m, deterministic)
}
func (m *Words) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Words.Merge(m, src)
}
func (m *Words) XXX_Size() int {
	return xxx_messageInfo_Words.Size(m)
}
func (m *Words) XXX_DiscardUnknown() {
	xxx_messageInfo_Words.DiscardUnknown(m)
}

var xxx_messageInfo_Words proto.InternalMessageInfo

func (m *Words) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Emoji struct {
	FaceId               int64    `protobuf:"varint,1,opt,name=face_id,json=faceId,proto3" json:"face_id,omitempty"`
	FaceUrl              string   `protobuf:"bytes,2,opt,name=face_url,json=faceUrl,proto3" json:"face_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Emoji) Reset()         { *m = Emoji{} }
func (m *Emoji) String() string { return proto.CompactTextString(m) }
func (*Emoji) ProtoMessage()    {}
func (*Emoji) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *Emoji) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Emoji.Unmarshal(m, b)
}
func (m *Emoji) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Emoji.Marshal(b, m, deterministic)
}
func (m *Emoji) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Emoji.Merge(m, src)
}
func (m *Emoji) XXX_Size() int {
	return xxx_messageInfo_Emoji.Size(m)
}
func (m *Emoji) XXX_DiscardUnknown() {
	xxx_messageInfo_Emoji.DiscardUnknown(m)
}

var xxx_messageInfo_Emoji proto.InternalMessageInfo

func (m *Emoji) GetFaceId() int64 {
	if m != nil {
		return m.FaceId
	}
	return 0
}

func (m *Emoji) GetFaceUrl() string {
	if m != nil {
		return m.FaceUrl
	}
	return ""
}

type Voice struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Duration             int32    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Voice) Reset()         { *m = Voice{} }
func (m *Voice) String() string { return proto.CompactTextString(m) }
func (*Voice) ProtoMessage()    {}
func (*Voice) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *Voice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Voice.Unmarshal(m, b)
}
func (m *Voice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Voice.Marshal(b, m, deterministic)
}
func (m *Voice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Voice.Merge(m, src)
}
func (m *Voice) XXX_Size() int {
	return xxx_messageInfo_Voice.Size(m)
}
func (m *Voice) XXX_DiscardUnknown() {
	xxx_messageInfo_Voice.DiscardUnknown(m)
}

var xxx_messageInfo_Voice proto.InternalMessageInfo

func (m *Voice) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Voice) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Voice) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Voice) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Image struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	ThumbnailUrl         string   `protobuf:"bytes,5,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

type File struct {
	Id                   int64    `protobuf:"varint,12,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,13,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64    `protobuf:"varint,14,opt,name=size,proto3" json:"size,omitempty"`
	Url                  string   `protobuf:"bytes,15,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Geo struct {
	Desc                 string   `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	Lat                  float64  `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  float64  `protobuf:"fixed64,3,opt,name=lng,proto3" json:"lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Geo) Reset()         { *m = Geo{} }
func (m *Geo) String() string { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()    {}
func (*Geo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{6}
}

func (m *Geo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Geo.Unmarshal(m, b)
}
func (m *Geo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Geo.Marshal(b, m, deterministic)
}
func (m *Geo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Geo.Merge(m, src)
}
func (m *Geo) XXX_Size() int {
	return xxx_messageInfo_Geo.Size(m)
}
func (m *Geo) XXX_DiscardUnknown() {
	xxx_messageInfo_Geo.DiscardUnknown(m)
}

var xxx_messageInfo_Geo proto.InternalMessageInfo

func (m *Geo) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Geo) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Geo) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

type Custom struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Custom) Reset()         { *m = Custom{} }
func (m *Custom) String() string { return proto.CompactTextString(m) }
func (*Custom) ProtoMessage()    {}
func (*Custom) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{7}
}

func (m *Custom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Custom.Unmarshal(m, b)
}
func (m *Custom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Custom.Marshal(b, m, deterministic)
}
func (m *Custom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Custom.Merge(m, src)
}
func (m *Custom) XXX_Size() int {
	return xxx_messageInfo_Custom.Size(m)
}
func (m *Custom) XXX_DiscardUnknown() {
	xxx_messageInfo_Custom.DiscardUnknown(m)
}

var xxx_messageInfo_Custom proto.InternalMessageInfo

func (m *Custom) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterEnum("rpc.MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("rpc.MsgStatus", MsgStatus_name, MsgStatus_value)
	proto.RegisterType((*Message)(nil), "rpc.Message")
	proto.RegisterType((*Words)(nil), "rpc.Words")
	proto.RegisterType((*Emoji)(nil), "rpc.Emoji")
	proto.RegisterType((*Voice)(nil), "rpc.Voice")
	proto.RegisterType((*Image)(nil), "rpc.Image")
	proto.RegisterType((*File)(nil), "rpc.File")
	proto.RegisterType((*Geo)(nil), "rpc.Geo")
	proto.RegisterType((*Custom)(nil), "rpc.Custom")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x89, 0x3f, 0xe2, 0x69, 0x9a, 0x5a, 0x2b, 0x04, 0xe6, 0x43, 0x28, 0x32, 0x12, 0xaa,
	0x7a, 0xe8, 0x01, 0x8e, 0x15, 0x87, 0xd2, 0xa6, 0xc5, 0xd0, 0x8d, 0xd1, 0xc6, 0x6d, 0x04, 0x17,
	0xcb, 0xb5, 0x97, 0x64, 0x51, 0x6c, 0x07, 0xef, 0x46, 0x50, 0xc4, 0x95, 0xff, 0x8d, 0x76, 0xec,
	0xb8, 0x41, 0xdc, 0xe6, 0xbd, 0x99, 0x9d, 0xd9, 0xf7, 0x76, 0x16, 0xdc, 0x42, 0x2e, 0x8e, 0xd7,
	0x75, 0xa5, 0x2a, 0xd2, 0xaf, 0xd7, 0x59, 0xf0, 0xa7, 0x07, 0x0e, 0xe5, 0x52, 0xa6, 0x0b, 0x4e,
	0x46, 0xd0, 0x13, 0xb9, 0x6f, 0x8c, 0x8d, 0xc3, 0x3e, 0xeb, 0x89, 0x9c, 0x3c, 0x02, 0x5b, 0xf2,
	0x32, 0x0f, 0x73, 0xbf, 0x87, 0x5c, 0x8b, 0xc8, 0x73, 0x70, 0x25, 0x97, 0x52, 0x54, 0x65, 0x98,
	0xfb, 0x7d, 0x4c, 0xdd, 0x13, 0x84, 0x80, 0xa9, 0x44, 0xc1, 0x7d, 0x13, 0x13, 0x18, 0xeb, 0x13,
	0x35, 0xff, 0xbe, 0xe1, 0x52, 0x85, 0xb9, 0x6f, 0x35, 0x27, 0x3a, 0x82, 0xbc, 0x02, 0x5b, 0xaa,
	0x54, 0x6d, 0xa4, 0x6f, 0x8f, 0x8d, 0xc3, 0xd1, 0xeb, 0xd1, 0x71, 0xbd, 0xce, 0x8e, 0xa9, 0x5c,
	0xcc, 0x90, 0x65, 0x6d, 0x96, 0x8c, 0xc1, 0x54, 0x77, 0x6b, 0xee, 0x3b, 0x58, 0x35, 0xdc, 0x56,
	0xc5, 0x77, 0x6b, 0xce, 0x30, 0x43, 0x7c, 0x70, 0xb2, 0xaa, 0x54, 0xbc, 0x54, 0xfe, 0x60, 0x6c,
	0x1c, 0x0e, 0xd9, 0x16, 0x92, 0x17, 0x00, 0x52, 0x0f, 0x2c, 0x33, 0x3e, 0xad, 0x7c, 0x17, 0xaf,
	0xb0, 0xc3, 0x04, 0xcf, 0xc0, 0x9a, 0x57, 0x75, 0x2e, 0xf1, 0xfa, 0xfc, 0xa7, 0x42, 0x1b, 0x5c,
	0x86, 0x71, 0x70, 0x02, 0xd6, 0xa4, 0xa8, 0xbe, 0x09, 0xf2, 0x18, 0x9c, 0xaf, 0x69, 0xc6, 0x93,
	0xce, 0x26, 0x5b, 0xc3, 0x30, 0x27, 0x4f, 0x60, 0x80, 0x89, 0x4d, 0xbd, 0x42, 0xb3, 0x5c, 0x86,
	0x85, 0xd7, 0xf5, 0x2a, 0xf8, 0x0c, 0xd6, 0x4d, 0x25, 0xb2, 0x5d, 0x7b, 0x5d, 0xb4, 0x97, 0x80,
	0x29, 0xc5, 0x2f, 0x8e, 0xf5, 0x16, 0xc3, 0x98, 0x3c, 0x85, 0x41, 0xbe, 0xa9, 0x53, 0x25, 0xaa,
	0x12, 0x9d, 0xb5, 0x58, 0x87, 0x89, 0x07, 0x7d, 0xdd, 0xde, 0xc4, 0x06, 0x3a, 0x0c, 0x7e, 0x83,
	0x15, 0x16, 0xff, 0xbe, 0x5c, 0xd3, 0xfa, 0x21, 0x58, 0x3f, 0x44, 0xae, 0x96, 0x6d, 0xef, 0x06,
	0xe8, 0xf7, 0x5c, 0x72, 0xb1, 0x58, 0xaa, 0xb6, 0x75, 0x8b, 0xfe, 0x6f, 0x4c, 0x5e, 0xc2, 0xbe,
	0x5a, 0x6e, 0x8a, 0xdb, 0x32, 0x15, 0x2b, 0xd4, 0x64, 0x61, 0x6e, 0xd8, 0x91, 0x5a, 0xd8, 0x27,
	0x30, 0x2f, 0xc4, 0x6a, 0x3b, 0x7c, 0xd8, 0xad, 0x0d, 0x01, 0xb3, 0x4c, 0x0b, 0xee, 0xef, 0x37,
	0x0e, 0xea, 0xb8, 0xd3, 0x3a, 0x6a, 0x96, 0x02, 0xb5, 0xb6, 0x63, 0x0f, 0xee, 0xf5, 0xbc, 0x85,
	0xfe, 0x25, 0xaf, 0x74, 0x71, 0xce, 0x65, 0xb6, 0x7d, 0x02, 0x1d, 0xeb, 0xe2, 0x55, 0xaa, 0x50,
	0x8f, 0xc1, 0x74, 0x88, 0x4c, 0xb9, 0x40, 0x29, 0x9a, 0x29, 0x17, 0x41, 0x00, 0xf6, 0xd9, 0x46,
	0xaa, 0xaa, 0xd8, 0xdd, 0x83, 0xa6, 0xc9, 0x16, 0x1e, 0xdd, 0x81, 0xd3, 0xae, 0x0c, 0x19, 0x01,
	0xd0, 0x38, 0xb9, 0x9e, 0x7e, 0x9c, 0x46, 0xf3, 0xa9, 0xf7, 0x80, 0x0c, 0x61, 0x40, 0xe3, 0x64,
	0x1e, 0xb1, 0xf3, 0x99, 0x67, 0xb4, 0x68, 0x42, 0xa3, 0x0f, 0xa1, 0xd7, 0x6b, 0xd1, 0x4d, 0x14,
	0x9e, 0x4d, 0xbc, 0x7e, 0x8b, 0x42, 0x7a, 0x7a, 0x39, 0xf1, 0x4c, 0xb2, 0x07, 0x0e, 0x8d, 0x93,
	0x8b, 0xf0, 0x6a, 0xe2, 0x59, 0x04, 0xc0, 0xa6, 0x71, 0x72, 0x39, 0x89, 0x3c, 0x9b, 0xec, 0x83,
	0x4b, 0xe3, 0xe4, 0xec, 0x7a, 0x16, 0x47, 0xd4, 0x73, 0x8e, 0x4e, 0xc0, 0xed, 0x76, 0x1a, 0x87,
	0xcf, 0x76, 0x86, 0xeb, 0xda, 0x59, 0x32, 0x8d, 0x18, 0x3d, 0xbd, 0xf2, 0x0c, 0x72, 0x00, 0x7b,
	0x74, 0x96, 0xcc, 0xc3, 0xf8, 0xfd, 0x39, 0x3b, 0x9d, 0x7b, 0xbd, 0x77, 0xd6, 0x17, 0xfd, 0x5d,
	0x6f, 0x6d, 0xfc, 0xba, 0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x58, 0x02, 0xed, 0xbc, 0xc7,
	0x03, 0x00, 0x00,
}
