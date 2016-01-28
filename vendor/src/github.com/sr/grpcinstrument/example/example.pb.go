// Code generated by protoc-gen-go.
// source: example/example.proto
// DO NOT EDIT!

package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "example.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "example.HelloResponse")
}

var fileDescriptor0 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x0c, 0x17, 0x8f, 0x47, 0x6a, 0x4e, 0x4e, 0x7e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x10, 0x0f, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x92, 0x02,
	0x17, 0x2f, 0x54, 0xb6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x88, 0x9f, 0x8b, 0x3d, 0x37, 0xb5,
	0xb8, 0x38, 0x31, 0x1d, 0xaa, 0xc2, 0xc8, 0x91, 0x8b, 0xdd, 0x15, 0x62, 0x94, 0x90, 0x19, 0x17,
	0x2b, 0x58, 0xb1, 0x90, 0xa8, 0x1e, 0xcc, 0x32, 0x64, 0xa3, 0xa5, 0xc4, 0xd0, 0x85, 0x21, 0x66,
	0x26, 0xb1, 0x81, 0x9d, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x8b, 0x36, 0x3a, 0xab,
	0x00, 0x00, 0x00,
}