// Code generated by protoc-gen-go.
// source: gcloud/gcloud.proto
// DO NOT EDIT!

/*
Package gcloud is a generated protocol buffer package.

It is generated from these files:
	gcloud/gcloud.proto

It has these top-level messages:
	Instance
	ListInstancesRequest
	ListInstancesResponse
*/
package gcloud

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/sr/operator/src/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Instance struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Zone   string `protobuf:"bytes,4,opt,name=zone" json:"zone,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListInstancesRequest struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id" json:"project_id,omitempty"`
}

func (m *ListInstancesRequest) Reset()                    { *m = ListInstancesRequest{} }
func (m *ListInstancesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListInstancesRequest) ProtoMessage()               {}
func (*ListInstancesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListInstancesResponse struct {
	Output  *proto1.Output `protobuf:"bytes,1,opt,name=output" json:"output,omitempty"`
	Objects []*Instance    `protobuf:"bytes,2,rep,name=objects" json:"objects,omitempty"`
}

func (m *ListInstancesResponse) Reset()                    { *m = ListInstancesResponse{} }
func (m *ListInstancesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListInstancesResponse) ProtoMessage()               {}
func (*ListInstancesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListInstancesResponse) GetOutput() *proto1.Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *ListInstancesResponse) GetObjects() []*Instance {
	if m != nil {
		return m.Objects
	}
	return nil
}

func init() {
	proto.RegisterType((*Instance)(nil), "gcloud.Instance")
	proto.RegisterType((*ListInstancesRequest)(nil), "gcloud.ListInstancesRequest")
	proto.RegisterType((*ListInstancesResponse)(nil), "gcloud.ListInstancesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for GCloudService service

type GCloudServiceClient interface {
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
}

type gCloudServiceClient struct {
	cc *grpc.ClientConn
}

func NewGCloudServiceClient(cc *grpc.ClientConn) GCloudServiceClient {
	return &gCloudServiceClient{cc}
}

func (c *gCloudServiceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := grpc.Invoke(ctx, "/gcloud.GCloudService/ListInstances", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GCloudService service

type GCloudServiceServer interface {
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
}

func RegisterGCloudServiceServer(s *grpc.Server, srv GCloudServiceServer) {
	s.RegisterService(&_GCloudService_serviceDesc, srv)
}

func _GCloudService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GCloudServiceServer).ListInstances(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _GCloudService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gcloud.GCloudService",
	HandlerType: (*GCloudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInstances",
			Handler:    _GCloudService_ListInstances_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4f, 0xce, 0xc9,
	0x2f, 0x4d, 0xd1, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94,
	0x08, 0x98, 0xab, 0x9f, 0x5f, 0x90, 0x5a, 0x94, 0x58, 0x92, 0x5f, 0x04, 0x91, 0x55, 0x72, 0xe2,
	0xe2, 0xf0, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0x15, 0xe2, 0xe2, 0x62, 0xca, 0x4c, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x14, 0xe2, 0xe1, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02,
	0xf3, 0xf8, 0xb8, 0xd8, 0x80, 0x6a, 0x4a, 0x4a, 0x8b, 0x25, 0x98, 0x61, 0xb2, 0x55, 0xf9, 0x79,
	0xa9, 0x12, 0x2c, 0x20, 0x9e, 0x92, 0x16, 0x97, 0x88, 0x4f, 0x66, 0x71, 0x09, 0xcc, 0x9c, 0xe2,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0xa0, 0x81, 0x40, 0x4b, 0xb2, 0x52, 0x93, 0x4b,
	0xe2, 0x61, 0xe6, 0x2a, 0x45, 0x72, 0x89, 0xa2, 0xa9, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x92, 0xe5, 0x62, 0xcb, 0x2f, 0x2d, 0x29, 0x28, 0x2d, 0x01, 0x2b, 0xe4, 0x36, 0xe2, 0x85, 0x38,
	0x50, 0xcf, 0x1f, 0x2c, 0x28, 0xa4, 0xc8, 0xc5, 0x9e, 0x9f, 0x04, 0x32, 0xaa, 0x18, 0xe8, 0x24,
	0x66, 0xa0, 0xbc, 0x80, 0x1e, 0xd4, 0x97, 0x30, 0xa3, 0x8c, 0x62, 0xb9, 0x78, 0xdd, 0x9d, 0x41,
	0x42, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0x40, 0xff, 0xf8, 0x70, 0xf1, 0xa2, 0xd8, 0x25, 0x24, 0x03,
	0xd3, 0x83, 0xcd, 0xb9, 0x52, 0xb2, 0x38, 0x64, 0x21, 0x0e, 0x4c, 0x62, 0x03, 0xbb, 0xc7, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x2f, 0xf0, 0xdc, 0x65, 0x01, 0x00, 0x00,
}