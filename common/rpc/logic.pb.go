// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logic.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("logic.proto", fileDescriptor_60207fea82c31ca8) }

var fileDescriptor_60207fea82c31ca8 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x6f, 0x4b, 0xeb, 0x30,
	0x14, 0xc6, 0x19, 0x97, 0xdd, 0xbb, 0x9d, 0x76, 0x77, 0xbb, 0xb9, 0x88, 0xd0, 0x97, 0x82, 0x82,
	0x03, 0xe7, 0xe6, 0x54, 0x7c, 0xbb, 0x0d, 0x11, 0xff, 0x0c, 0xb1, 0x43, 0x04, 0xdf, 0x75, 0xd9,
	0x59, 0x0d, 0x6b, 0x9b, 0x9a, 0xc4, 0xc9, 0x3e, 0x8d, 0x5f, 0x55, 0x92, 0xa6, 0xae, 0xad, 0xbe,
	0xeb, 0xf3, 0x9c, 0xdf, 0x93, 0x9c, 0x73, 0x1a, 0x70, 0x22, 0x1e, 0x32, 0xda, 0x4b, 0x05, 0x57,
	0x9c, 0xfc, 0x12, 0x29, 0xf5, 0x9a, 0xb1, 0x0c, 0x33, 0xed, 0xb9, 0x94, 0xc7, 0x31, 0x4f, 0xac,
	0x72, 0x62, 0xbe, 0xc0, 0xc8, 0x8a, 0x66, 0x40, 0x57, 0xd9, 0xe7, 0xc9, 0x47, 0x1d, 0xdc, 0x3b,
	0x7d, 0xca, 0x0c, 0xc5, 0x9a, 0x51, 0x24, 0xc7, 0xd0, 0xf0, 0x31, 0x64, 0x52, 0xa1, 0x20, 0x9d,
	0x9e, 0x48, 0x69, 0x2f, 0x97, 0x3e, 0xbe, 0x7a, 0xff, 0x2a, 0x8e, 0x4c, 0xc9, 0x21, 0x38, 0x57,
	0xa8, 0x66, 0x28, 0x25, 0xe3, 0x89, 0x24, 0x60, 0x88, 0xeb, 0x44, 0x9d, 0x9f, 0x7a, 0xae, 0xf9,
	0xb6, 0xa5, 0x7e, 0x8d, 0x74, 0x01, 0xb6, 0x68, 0x89, 0xec, 0x14, 0x49, 0x73, 0xec, 0x05, 0xb4,
	0x26, 0x02, 0x03, 0x85, 0x39, 0xbe, 0x63, 0x90, 0x92, 0xa7, 0x3b, 0xfa, 0x9e, 0x3c, 0x02, 0xe7,
	0x86, 0xb3, 0x24, 0xcf, 0xfd, 0x37, 0x40, 0xc1, 0xd1, 0xa9, 0xa6, 0x31, 0xc7, 0x9c, 0x47, 0x1a,
	0x7f, 0x78, 0x63, 0xaa, 0x8c, 0x17, 0x9c, 0x0a, 0x3e, 0x80, 0x96, 0x8f, 0x49, 0x10, 0x57, 0xfa,
	0x2a, 0x79, 0x95, 0x48, 0x17, 0xfe, 0xfa, 0x48, 0x91, 0xad, 0x71, 0x8a, 0x52, 0x06, 0x21, 0x92,
	0x6c, 0x31, 0x56, 0x79, 0x85, 0x45, 0x90, 0x3d, 0x00, 0xcb, 0x8e, 0x26, 0xb7, 0xa4, 0x61, 0x2a,
	0x23, 0xba, 0xb2, 0xcc, 0x65, 0x9c, 0xaa, 0x0d, 0xd9, 0x37, 0x6b, 0x9c, 0x62, 0x3c, 0x47, 0x51,
	0x5e, 0x78, 0x76, 0xe9, 0xa3, 0x44, 0xd1, 0xaf, 0x91, 0x21, 0xb8, 0xb3, 0x4d, 0x42, 0xed, 0x2d,
	0xd2, 0x4e, 0x56, 0xb0, 0x74, 0x9b, 0xa5, 0x4e, 0xfa, 0x35, 0x32, 0x30, 0x7f, 0xf3, 0x2b, 0x43,
	0x4c, 0x79, 0xeb, 0xfc, 0x14, 0x39, 0x83, 0xf6, 0x13, 0x53, 0x2f, 0x0b, 0x11, 0xbc, 0xe7, 0xf3,
	0xed, 0x1a, 0xa4, 0xe2, 0x56, 0xb6, 0x72, 0x00, 0x7f, 0xee, 0x97, 0xcb, 0x88, 0x25, 0x48, 0xda,
	0xc6, 0xb5, 0x4a, 0x63, 0x85, 0x69, 0xc7, 0xf5, 0x67, 0xfd, 0xb2, 0xe7, 0xbf, 0xcd, 0x7b, 0x1d,
	0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x22, 0x91, 0x92, 0xf4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogicServiceClient is the client API for LogicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicServiceClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	GetSessions(ctx context.Context, in *Int64, opts ...grpc.CallOption) (LogicService_GetSessionsClient, error)
	GetSession(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SessionResp, error)
	CreateSession(ctx context.Context, in *CreateSessionReq, opts ...grpc.CallOption) (*SessionResp, error)
	JoinSession(ctx context.Context, in *JoinSessionReq, opts ...grpc.CallOption) (*Bool, error)
	QuitSession(ctx context.Context, in *QuitSessionReq, opts ...grpc.CallOption) (*Bool, error)
	RenameSession(ctx context.Context, in *RenameSessionReq, opts ...grpc.CallOption) (*Bool, error)
	ReceiveMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Int64, error)
	ReceiveACK(ctx context.Context, in *Ack, opts ...grpc.CallOption) (*Empty, error)
	GetMembers(ctx context.Context, in *Int64, opts ...grpc.CallOption) (LogicService_GetMembersClient, error)
	SyncMessages(ctx context.Context, in *SyncMessageReq, opts ...grpc.CallOption) (LogicService_SyncMessagesClient, error)
	GetMessages(ctx context.Context, in *GetMessageReq, opts ...grpc.CallOption) (LogicService_GetMessagesClient, error)
	WithdrawMessage(ctx context.Context, in *WithdrawMessageReq, opts ...grpc.CallOption) (*Bool, error)
	Offline(ctx context.Context, in *OfflineReq, opts ...grpc.CallOption) (*Empty, error)
}

type logicServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogicServiceClient(cc *grpc.ClientConn) LogicServiceClient {
	return &logicServiceClient{cc}
}

func (c *logicServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) GetSessions(ctx context.Context, in *Int64, opts ...grpc.CallOption) (LogicService_GetSessionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogicService_serviceDesc.Streams[0], "/rpc.LogicService/GetSessions", opts...)
	if err != nil {
		return nil, err
	}
	x := &logicServiceGetSessionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogicService_GetSessionsClient interface {
	Recv() (*Session, error)
	grpc.ClientStream
}

type logicServiceGetSessionsClient struct {
	grpc.ClientStream
}

func (x *logicServiceGetSessionsClient) Recv() (*Session, error) {
	m := new(Session)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logicServiceClient) GetSession(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SessionResp, error) {
	out := new(SessionResp)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) CreateSession(ctx context.Context, in *CreateSessionReq, opts ...grpc.CallOption) (*SessionResp, error) {
	out := new(SessionResp)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) JoinSession(ctx context.Context, in *JoinSessionReq, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/JoinSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) QuitSession(ctx context.Context, in *QuitSessionReq, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/QuitSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) RenameSession(ctx context.Context, in *RenameSessionReq, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/RenameSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) ReceiveMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Int64, error) {
	out := new(Int64)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/ReceiveMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) ReceiveACK(ctx context.Context, in *Ack, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/ReceiveACK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) GetMembers(ctx context.Context, in *Int64, opts ...grpc.CallOption) (LogicService_GetMembersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogicService_serviceDesc.Streams[1], "/rpc.LogicService/GetMembers", opts...)
	if err != nil {
		return nil, err
	}
	x := &logicServiceGetMembersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogicService_GetMembersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type logicServiceGetMembersClient struct {
	grpc.ClientStream
}

func (x *logicServiceGetMembersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logicServiceClient) SyncMessages(ctx context.Context, in *SyncMessageReq, opts ...grpc.CallOption) (LogicService_SyncMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogicService_serviceDesc.Streams[2], "/rpc.LogicService/SyncMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &logicServiceSyncMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogicService_SyncMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type logicServiceSyncMessagesClient struct {
	grpc.ClientStream
}

func (x *logicServiceSyncMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logicServiceClient) GetMessages(ctx context.Context, in *GetMessageReq, opts ...grpc.CallOption) (LogicService_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogicService_serviceDesc.Streams[3], "/rpc.LogicService/GetMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &logicServiceGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogicService_GetMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type logicServiceGetMessagesClient struct {
	grpc.ClientStream
}

func (x *logicServiceGetMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logicServiceClient) WithdrawMessage(ctx context.Context, in *WithdrawMessageReq, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/WithdrawMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicServiceClient) Offline(ctx context.Context, in *OfflineReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.LogicService/Offline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServiceServer is the server API for LogicService service.
type LogicServiceServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	GetSessions(*Int64, LogicService_GetSessionsServer) error
	GetSession(context.Context, *Int64) (*SessionResp, error)
	CreateSession(context.Context, *CreateSessionReq) (*SessionResp, error)
	JoinSession(context.Context, *JoinSessionReq) (*Bool, error)
	QuitSession(context.Context, *QuitSessionReq) (*Bool, error)
	RenameSession(context.Context, *RenameSessionReq) (*Bool, error)
	ReceiveMessage(context.Context, *Message) (*Int64, error)
	ReceiveACK(context.Context, *Ack) (*Empty, error)
	GetMembers(*Int64, LogicService_GetMembersServer) error
	SyncMessages(*SyncMessageReq, LogicService_SyncMessagesServer) error
	GetMessages(*GetMessageReq, LogicService_GetMessagesServer) error
	WithdrawMessage(context.Context, *WithdrawMessageReq) (*Bool, error)
	Offline(context.Context, *OfflineReq) (*Empty, error)
}

// UnimplementedLogicServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogicServiceServer struct {
}

func (*UnimplementedLogicServiceServer) Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedLogicServiceServer) GetSessions(req *Int64, srv LogicService_GetSessionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSessions not implemented")
}
func (*UnimplementedLogicServiceServer) GetSession(ctx context.Context, req *Int64) (*SessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (*UnimplementedLogicServiceServer) CreateSession(ctx context.Context, req *CreateSessionReq) (*SessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (*UnimplementedLogicServiceServer) JoinSession(ctx context.Context, req *JoinSessionReq) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinSession not implemented")
}
func (*UnimplementedLogicServiceServer) QuitSession(ctx context.Context, req *QuitSessionReq) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuitSession not implemented")
}
func (*UnimplementedLogicServiceServer) RenameSession(ctx context.Context, req *RenameSessionReq) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameSession not implemented")
}
func (*UnimplementedLogicServiceServer) ReceiveMessage(ctx context.Context, req *Message) (*Int64, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}
func (*UnimplementedLogicServiceServer) ReceiveACK(ctx context.Context, req *Ack) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveACK not implemented")
}
func (*UnimplementedLogicServiceServer) GetMembers(req *Int64, srv LogicService_GetMembersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}
func (*UnimplementedLogicServiceServer) SyncMessages(req *SyncMessageReq, srv LogicService_SyncMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncMessages not implemented")
}
func (*UnimplementedLogicServiceServer) GetMessages(req *GetMessageReq, srv LogicService_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (*UnimplementedLogicServiceServer) WithdrawMessage(ctx context.Context, req *WithdrawMessageReq) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawMessage not implemented")
}
func (*UnimplementedLogicServiceServer) Offline(ctx context.Context, req *OfflineReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offline not implemented")
}

func RegisterLogicServiceServer(s *grpc.Server, srv LogicServiceServer) {
	s.RegisterService(&_LogicService_serviceDesc, srv)
}

func _LogicService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_GetSessions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Int64)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogicServiceServer).GetSessions(m, &logicServiceGetSessionsServer{stream})
}

type LogicService_GetSessionsServer interface {
	Send(*Session) error
	grpc.ServerStream
}

type logicServiceGetSessionsServer struct {
	grpc.ServerStream
}

func (x *logicServiceGetSessionsServer) Send(m *Session) error {
	return x.ServerStream.SendMsg(m)
}

func _LogicService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).GetSession(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).CreateSession(ctx, req.(*CreateSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_JoinSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).JoinSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/JoinSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).JoinSession(ctx, req.(*JoinSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_QuitSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuitSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).QuitSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/QuitSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).QuitSession(ctx, req.(*QuitSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_RenameSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).RenameSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/RenameSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).RenameSession(ctx, req.(*RenameSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_ReceiveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).ReceiveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/ReceiveMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).ReceiveMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_ReceiveACK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).ReceiveACK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/ReceiveACK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).ReceiveACK(ctx, req.(*Ack))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_GetMembers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Int64)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogicServiceServer).GetMembers(m, &logicServiceGetMembersServer{stream})
}

type LogicService_GetMembersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type logicServiceGetMembersServer struct {
	grpc.ServerStream
}

func (x *logicServiceGetMembersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _LogicService_SyncMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncMessageReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogicServiceServer).SyncMessages(m, &logicServiceSyncMessagesServer{stream})
}

type LogicService_SyncMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type logicServiceSyncMessagesServer struct {
	grpc.ServerStream
}

func (x *logicServiceSyncMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _LogicService_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMessageReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogicServiceServer).GetMessages(m, &logicServiceGetMessagesServer{stream})
}

type LogicService_GetMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type logicServiceGetMessagesServer struct {
	grpc.ServerStream
}

func (x *logicServiceGetMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _LogicService_WithdrawMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).WithdrawMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/WithdrawMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).WithdrawMessage(ctx, req.(*WithdrawMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicService_Offline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfflineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServiceServer).Offline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicService/Offline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServiceServer).Offline(ctx, req.(*OfflineReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.LogicService",
	HandlerType: (*LogicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _LogicService_Register_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _LogicService_GetSession_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _LogicService_CreateSession_Handler,
		},
		{
			MethodName: "JoinSession",
			Handler:    _LogicService_JoinSession_Handler,
		},
		{
			MethodName: "QuitSession",
			Handler:    _LogicService_QuitSession_Handler,
		},
		{
			MethodName: "RenameSession",
			Handler:    _LogicService_RenameSession_Handler,
		},
		{
			MethodName: "ReceiveMessage",
			Handler:    _LogicService_ReceiveMessage_Handler,
		},
		{
			MethodName: "ReceiveACK",
			Handler:    _LogicService_ReceiveACK_Handler,
		},
		{
			MethodName: "WithdrawMessage",
			Handler:    _LogicService_WithdrawMessage_Handler,
		},
		{
			MethodName: "Offline",
			Handler:    _LogicService_Offline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSessions",
			Handler:       _LogicService_GetSessions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMembers",
			Handler:       _LogicService_GetMembers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SyncMessages",
			Handler:       _LogicService_SyncMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMessages",
			Handler:       _LogicService_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "logic.proto",
}
