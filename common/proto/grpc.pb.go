// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RPCRequest struct {
	AccountID            uint32   `protobuf:"varint,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	MessageID            uint32   `protobuf:"varint,2,opt,name=messageID,proto3" json:"messageID,omitempty"`
	MessageBody          []byte   `protobuf:"bytes,3,opt,name=messageBody,proto3" json:"messageBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCRequest) Reset()         { *m = RPCRequest{} }
func (m *RPCRequest) String() string { return proto.CompactTextString(m) }
func (*RPCRequest) ProtoMessage()    {}
func (*RPCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_1a142ccc69e146e8, []int{0}
}
func (m *RPCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCRequest.Unmarshal(m, b)
}
func (m *RPCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCRequest.Marshal(b, m, deterministic)
}
func (dst *RPCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCRequest.Merge(dst, src)
}
func (m *RPCRequest) XXX_Size() int {
	return xxx_messageInfo_RPCRequest.Size(m)
}
func (m *RPCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPCRequest proto.InternalMessageInfo

func (m *RPCRequest) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *RPCRequest) GetMessageID() uint32 {
	if m != nil {
		return m.MessageID
	}
	return 0
}

func (m *RPCRequest) GetMessageBody() []byte {
	if m != nil {
		return m.MessageBody
	}
	return nil
}

type RPCReply struct {
	AccountID            uint32   `protobuf:"varint,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	MessageID            uint32   `protobuf:"varint,2,opt,name=messageID,proto3" json:"messageID,omitempty"`
	MessageBody          []byte   `protobuf:"bytes,3,opt,name=messageBody,proto3" json:"messageBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCReply) Reset()         { *m = RPCReply{} }
func (m *RPCReply) String() string { return proto.CompactTextString(m) }
func (*RPCReply) ProtoMessage()    {}
func (*RPCReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_1a142ccc69e146e8, []int{1}
}
func (m *RPCReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCReply.Unmarshal(m, b)
}
func (m *RPCReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCReply.Marshal(b, m, deterministic)
}
func (dst *RPCReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCReply.Merge(dst, src)
}
func (m *RPCReply) XXX_Size() int {
	return xxx_messageInfo_RPCReply.Size(m)
}
func (m *RPCReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCReply.DiscardUnknown(m)
}

var xxx_messageInfo_RPCReply proto.InternalMessageInfo

func (m *RPCReply) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *RPCReply) GetMessageID() uint32 {
	if m != nil {
		return m.MessageID
	}
	return 0
}

func (m *RPCReply) GetMessageBody() []byte {
	if m != nil {
		return m.MessageBody
	}
	return nil
}

func init() {
	proto.RegisterType((*RPCRequest)(nil), "proto.RPCRequest")
	proto.RegisterType((*RPCReply)(nil), "proto.RPCReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClient interface {
	Forward(ctx context.Context, opts ...grpc.CallOption) (Stream_ForwardClient, error)
}

type streamClient struct {
	cc *grpc.ClientConn
}

func NewStreamClient(cc *grpc.ClientConn) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) Forward(ctx context.Context, opts ...grpc.CallOption) (Stream_ForwardClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[0], "/proto.Stream/Forward", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamForwardClient{stream}
	return x, nil
}

type Stream_ForwardClient interface {
	Send(*RPCRequest) error
	Recv() (*RPCReply, error)
	grpc.ClientStream
}

type streamForwardClient struct {
	grpc.ClientStream
}

func (x *streamForwardClient) Send(m *RPCRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamForwardClient) Recv() (*RPCReply, error) {
	m := new(RPCReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServer is the server API for Stream service.
type StreamServer interface {
	Forward(Stream_ForwardServer) error
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
	s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_Forward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServer).Forward(&streamForwardServer{stream})
}

type Stream_ForwardServer interface {
	Send(*RPCReply) error
	Recv() (*RPCRequest, error)
	grpc.ServerStream
}

type streamForwardServer struct {
	grpc.ServerStream
}

func (x *streamForwardServer) Send(m *RPCReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamForwardServer) Recv() (*RPCRequest, error) {
	m := new(RPCRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Stream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Forward",
			Handler:       _Stream_Forward_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_grpc_1a142ccc69e146e8) }

var fileDescriptor_grpc_1a142ccc69e146e8 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x59, 0x5c, 0x5c, 0x41, 0x01,
	0xce, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x32, 0x5c, 0x9c, 0x89, 0xc9, 0xc9, 0xf9,
	0xa5, 0x79, 0x25, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x08, 0x01, 0x90, 0x6c,
	0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0xa7, 0x8b, 0x04, 0x13, 0x44, 0x16, 0x2e, 0x20, 0xa4,
	0xc0, 0xc5, 0x0d, 0xe5, 0x38, 0xe5, 0xa7, 0x54, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x04, 0x21,
	0x0b, 0x29, 0x65, 0x70, 0x71, 0x80, 0xed, 0x2a, 0xc8, 0xa9, 0xa4, 0xad, 0x4d, 0x46, 0xb6, 0x5c,
	0x6c, 0xc1, 0x25, 0x45, 0xa9, 0x89, 0xb9, 0x42, 0xc6, 0x5c, 0xec, 0x6e, 0xf9, 0x45, 0xe5, 0x89,
	0x45, 0x29, 0x42, 0x82, 0x10, 0x9f, 0xeb, 0x21, 0xfc, 0x2b, 0xc5, 0x8f, 0x2c, 0x54, 0x90, 0x53,
	0xa9, 0xc4, 0xa0, 0xc1, 0x68, 0xc0, 0xe8, 0xc4, 0x14, 0xc0, 0x98, 0xc4, 0x06, 0x96, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x3b, 0x5e, 0xf8, 0x34, 0x01, 0x00, 0x00,
}