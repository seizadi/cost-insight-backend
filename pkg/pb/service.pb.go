// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/seizadi/aws-cost/pkg/pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// TODO: Structure your own protobuf messages. Each protocol buffer message is a
// small logical record of information, containing a series of name-value pairs.
type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32bf3e45ff3a1a40, []int{0}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionResponse)(nil), "awscost.VersionResponse")
}

func init() {
	proto.RegisterFile("github.com/seizadi/aws-cost/pkg/pb/service.proto", fileDescriptor_32bf3e45ff3a1a40)
}

var fileDescriptor_32bf3e45ff3a1a40 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x55, 0x06, 0x02, 0x5e, 0x40, 0x19, 0x50, 0x14, 0x18, 0x50, 0x25, 0x04, 0x12, 0xc4,
	0x87, 0x60, 0x64, 0x2a, 0x08, 0xb1, 0xa2, 0x0e, 0x0c, 0x30, 0x39, 0xe9, 0x61, 0x2c, 0x5a, 0x9f,
	0xe5, 0x73, 0x13, 0xc2, 0xc8, 0x2b, 0xf0, 0x68, 0xbc, 0x02, 0x0f, 0x82, 0x9a, 0xb8, 0x52, 0x45,
	0x87, 0x6e, 0xf6, 0xfd, 0x77, 0xff, 0xa7, 0xff, 0x17, 0x97, 0xda, 0x84, 0xb7, 0x79, 0x29, 0x2b,
	0x9a, 0x01, 0xa3, 0xf9, 0x54, 0x13, 0x03, 0xaa, 0xe1, 0xa2, 0x22, 0x0e, 0xe0, 0xde, 0x35, 0xb8,
	0x12, 0x18, 0x7d, 0x6d, 0x2a, 0x94, 0xce, 0x53, 0xa0, 0x34, 0x51, 0x0d, 0x2f, 0xd4, 0xfc, 0x50,
	0x13, 0xe9, 0x29, 0x42, 0x37, 0x2e, 0xe7, 0xaf, 0x80, 0x33, 0x17, 0xda, 0x7e, 0x2b, 0x3f, 0x8a,
	0xa2, 0x72, 0x06, 0x94, 0xb5, 0x14, 0x54, 0x30, 0x64, 0x39, 0xaa, 0xa3, 0x15, 0x2a, 0xda, 0x9a,
	0x5a, 0xe7, 0xe9, 0xa3, 0xed, 0x9d, 0xaa, 0x42, 0xa3, 0x2d, 0x6a, 0x35, 0x35, 0x13, 0x15, 0x10,
	0xd6, 0x1e, 0xd1, 0xe2, 0x62, 0x65, 0x99, 0x1b, 0xa5, 0x35, 0x7a, 0x20, 0xd7, 0x41, 0xd6, 0x81,
	0xc3, 0x73, 0xb1, 0xf7, 0x84, 0x9e, 0x0d, 0xd9, 0x31, 0xb2, 0x23, 0xcb, 0x98, 0x66, 0x22, 0xa9,
	0xfb, 0x51, 0x36, 0x38, 0x1e, 0x9c, 0xed, 0x8e, 0x97, 0xdf, 0xab, 0x17, 0x91, 0x8c, 0x1a, 0xbe,
	0x23, 0x0e, 0xe9, 0xa3, 0x10, 0x0f, 0x18, 0xe2, 0x69, 0x7a, 0x20, 0xfb, 0x54, 0x72, 0x19, 0x59,
	0xde, 0x2f, 0x22, 0xe7, 0x99, 0x8c, 0x9d, 0xc8, 0x7f, 0x90, 0xe1, 0xfe, 0xd7, 0xcf, 0xef, 0xf7,
	0x96, 0x48, 0x77, 0x20, 0x9a, 0xdf, 0x9e, 0x3e, 0x9f, 0x6c, 0xae, 0xfc, 0xc6, 0x95, 0xe5, 0x76,
	0x07, 0xb9, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x58, 0x95, 0x78, 0x21, 0xa2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AwsCostClient is the client API for AwsCost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AwsCostClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type awsCostClient struct {
	cc *grpc.ClientConn
}

func NewAwsCostClient(cc *grpc.ClientConn) AwsCostClient {
	return &awsCostClient{cc}
}

func (c *awsCostClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/awscost.AwsCost/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AwsCostServer is the server API for AwsCost service.
type AwsCostServer interface {
	GetVersion(context.Context, *empty.Empty) (*VersionResponse, error)
}

func RegisterAwsCostServer(s *grpc.Server, srv AwsCostServer) {
	s.RegisterService(&_AwsCost_serviceDesc, srv)
}

func _AwsCost_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsCostServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/awscost.AwsCost/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsCostServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AwsCost_serviceDesc = grpc.ServiceDesc{
	ServiceName: "awscost.AwsCost",
	HandlerType: (*AwsCostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _AwsCost_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/seizadi/aws-cost/pkg/pb/service.proto",
}
