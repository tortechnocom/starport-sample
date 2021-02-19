// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: starportsample/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("starportsample/query.proto", fileDescriptor_6286bf60d9e8499a) }

var fileDescriptor_6286bf60d9e8499a = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xcf, 0x31, 0x6a, 0xc4, 0x30,
	0x10, 0x05, 0x50, 0xbb, 0x48, 0x02, 0x2e, 0x53, 0x9a, 0xa0, 0x03, 0x18, 0xe2, 0xc1, 0x49, 0x9f,
	0x22, 0x37, 0x48, 0x17, 0xd2, 0x8d, 0x84, 0x90, 0x05, 0x96, 0x46, 0x91, 0xc6, 0x21, 0xbe, 0x45,
	0x8e, 0x95, 0xd2, 0xe5, 0x96, 0x8b, 0x7d, 0x91, 0x65, 0x6d, 0x2f, 0xec, 0xba, 0x1c, 0x78, 0xcc,
	0xff, 0xbf, 0x28, 0x13, 0x63, 0x0c, 0x14, 0x39, 0xa1, 0x0b, 0x9d, 0x86, 0xef, 0x5e, 0xc7, 0xa1,
	0x0e, 0x91, 0x98, 0x1e, 0x2b, 0xa6, 0xc8, 0x5a, 0xb5, 0x9e, 0x14, 0xb9, 0xfa, 0x16, 0xee, 0xce,
	0xf2, 0xc9, 0x10, 0x99, 0x4e, 0x03, 0x06, 0x0b, 0xe8, 0x3d, 0x31, 0xb2, 0x25, 0x9f, 0xd6, 0x4f,
	0x65, 0xa5, 0x28, 0x39, 0x4a, 0x20, 0x31, 0x6d, 0x11, 0xf0, 0xd3, 0x48, 0xcd, 0xd8, 0x40, 0x40,
	0x63, 0xfd, 0x82, 0x57, 0xfb, 0xf2, 0x50, 0xdc, 0x7d, 0x9c, 0xc5, 0xfb, 0xe7, 0xff, 0x24, 0xf2,
	0x71, 0x12, 0xf9, 0x71, 0x12, 0xf9, 0xdf, 0x2c, 0xb2, 0x71, 0x16, 0xd9, 0x61, 0x16, 0xd9, 0xd7,
	0x9b, 0xb1, 0xdc, 0xf6, 0xb2, 0x56, 0xe4, 0xe0, 0xba, 0x23, 0x5c, 0x4a, 0x3d, 0x6f, 0x6b, 0x7e,
	0x61, 0x37, 0x8f, 0x87, 0xa0, 0x93, 0xbc, 0x5f, 0x92, 0x5e, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x9c, 0x9a, 0x56, 0xfd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tortechnocom.starportsample.starportsample.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "starportsample/query.proto",
}