// Code generated by protoc-gen-go. DO NOT EDIT.
// source: servDiscover.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	servDiscover.proto

It has these top-level messages:
	GPRCRequest
	GPRCReply
*/
package pb

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
type GPRCRequest struct {
	Getver int32 `protobuf:"varint,1,opt,name=getver" json:"getver,omitempty"`
}

func (m *GPRCRequest) Reset()                    { *m = GPRCRequest{} }
func (m *GPRCRequest) String() string            { return proto.CompactTextString(m) }
func (*GPRCRequest) ProtoMessage()               {}
func (*GPRCRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GPRCRequest) GetGetver() int32 {
	if m != nil {
		return m.Getver
	}
	return 0
}

// The response message containing the greetings
type GPRCReply struct {
	//    repeated KV kvs = 1;
	Softver  string `protobuf:"bytes,1,opt,name=softver" json:"softver,omitempty"`
	Softname string `protobuf:"bytes,2,opt,name=softname" json:"softname,omitempty"`
}

func (m *GPRCReply) Reset()                    { *m = GPRCReply{} }
func (m *GPRCReply) String() string            { return proto.CompactTextString(m) }
func (*GPRCReply) ProtoMessage()               {}
func (*GPRCReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GPRCReply) GetSoftver() string {
	if m != nil {
		return m.Softver
	}
	return ""
}

func (m *GPRCReply) GetSoftname() string {
	if m != nil {
		return m.Softname
	}
	return ""
}

func init() {
	proto.RegisterType((*GPRCRequest)(nil), "pb.GPRCRequest")
	proto.RegisterType((*GPRCReply)(nil), "pb.GPRCReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GetVer service

type GetVerClient interface {
	// Sends a greeting
	Getvers(ctx context.Context, in *GPRCRequest, opts ...grpc.CallOption) (*GPRCReply, error)
}

type getVerClient struct {
	cc *grpc.ClientConn
}

func NewGetVerClient(cc *grpc.ClientConn) GetVerClient {
	return &getVerClient{cc}
}

func (c *getVerClient) Getvers(ctx context.Context, in *GPRCRequest, opts ...grpc.CallOption) (*GPRCReply, error) {
	out := new(GPRCReply)
	err := grpc.Invoke(ctx, "/pb.GetVer/Getvers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetVer service

type GetVerServer interface {
	// Sends a greeting
	Getvers(context.Context, *GPRCRequest) (*GPRCReply, error)
}

func RegisterGetVerServer(s *grpc.Server, srv GetVerServer) {
	s.RegisterService(&_GetVer_serviceDesc, srv)
}

func _GetVer_Getvers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GPRCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetVerServer).Getvers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GetVer/Getvers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetVerServer).Getvers(ctx, req.(*GPRCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetVer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GetVer",
	HandlerType: (*GetVerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Getvers",
			Handler:    _GetVer_Getvers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servDiscover.proto",
}

func init() { proto.RegisterFile("servDiscover.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0x2d, 0x2a,
	0x73, 0xc9, 0x2c, 0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2a, 0x48, 0x52, 0x52, 0xe5, 0xe2, 0x76, 0x0f, 0x08, 0x72, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe3, 0x62, 0x4b, 0x4f, 0x2d, 0x29, 0x4b, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0d, 0x82, 0xf2, 0x94, 0x1c, 0xb9, 0x38, 0x21, 0xca, 0x0a, 0x72, 0x2a, 0x85, 0x24, 0xb8,
	0xd8, 0x8b, 0xf3, 0xd3, 0xe0, 0xaa, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x29, 0x2e, 0x0e, 0x10, 0x33,
	0x2f, 0x31, 0x37, 0x55, 0x82, 0x09, 0x2c, 0x05, 0xe7, 0x1b, 0x99, 0x72, 0xb1, 0xb9, 0xa7, 0x96,
	0x84, 0xa5, 0x16, 0x09, 0x69, 0x73, 0xb1, 0xbb, 0x83, 0x8d, 0x2d, 0x16, 0xe2, 0xd7, 0x2b, 0x48,
	0xd2, 0x43, 0x72, 0x80, 0x14, 0x2f, 0x42, 0xa0, 0x20, 0xa7, 0x52, 0x89, 0x21, 0x89, 0x0d, 0xec,
	0x56, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xb3, 0x91, 0x1d, 0xc1, 0x00, 0x00, 0x00,
}
