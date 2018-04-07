// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello/hello.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	hello/hello.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Fuck    int64  `protobuf:"varint,2,opt,name=fuck" json:"fuck,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetFuck() int64 {
	if m != nil {
		return m.Fuck
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/hello.proto",
}

func init() { proto.RegisterFile("hello/hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x60, 0x4e, 0x79, 0x7e,
	0x51, 0x4e, 0x8a, 0x92, 0x09, 0x17, 0x8f, 0x07, 0x88, 0x17, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x22, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0x66, 0x0b, 0x09, 0x70, 0x31, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x07, 0x81,
	0x98, 0x4a, 0x56, 0x5c, 0x5c, 0x50, 0x5d, 0x05, 0x39, 0x95, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9,
	0xc5, 0xc5, 0x20, 0x35, 0x10, 0x6d, 0x30, 0x2e, 0xc8, 0xb4, 0xb4, 0xd2, 0xe4, 0x6c, 0xa8, 0x56,
	0x30, 0xdb, 0xc8, 0x93, 0x8b, 0xdd, 0xbd, 0x28, 0x35, 0xb5, 0x24, 0xb5, 0x48, 0xc8, 0x8e, 0x8b,
	0x23, 0x38, 0xb1, 0x12, 0x6c, 0x92, 0x90, 0x84, 0x1e, 0xc2, 0x55, 0x7a, 0xc8, 0x4e, 0x92, 0x12,
	0xc3, 0x22, 0x53, 0x90, 0x53, 0xa9, 0xc4, 0xe0, 0xc4, 0xb6, 0x88, 0x89, 0xd9, 0xc3, 0x27, 0x3c,
	0x89, 0x0d, 0xec, 0x2f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xe5, 0xbd, 0x10, 0xec,
	0x00, 0x00, 0x00,
}