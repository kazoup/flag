// Code generated by protoc-gen-go.
// source: github.com/Rakanixu/flag/srv/proto/flag/flag.proto
// DO NOT EDIT!

/*
Package go_micro_srv_flag is a generated protocol buffer package.

It is generated from these files:
	github.com/Rakanixu/flag/srv/proto/flag/flag.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	FlipRequest
	FlipResponse
	DeleteRequest
	DeleteResponse
	ListRequest
	ListResponse
*/
package go_micro_srv_flag

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type CreateRequest struct {
	Key         string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Value       bool   `protobuf:"varint,3,opt,name=value" json:"value,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ReadRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ReadResponse struct {
	Key         string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Value       bool   `protobuf:"varint,3,opt,name=value" json:"value,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type FlipRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *FlipRequest) Reset()                    { *m = FlipRequest{} }
func (m *FlipRequest) String() string            { return proto.CompactTextString(m) }
func (*FlipRequest) ProtoMessage()               {}
func (*FlipRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type FlipResponse struct {
}

func (m *FlipResponse) Reset()                    { *m = FlipResponse{} }
func (m *FlipResponse) String() string            { return proto.CompactTextString(m) }
func (*FlipResponse) ProtoMessage()               {}
func (*FlipResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ListResponse struct {
	Result []*ReadResponse `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListResponse) GetResult() []*ReadResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.flag.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.srv.flag.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.srv.flag.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.srv.flag.ReadResponse")
	proto.RegisterType((*FlipRequest)(nil), "go.micro.srv.flag.FlipRequest")
	proto.RegisterType((*FlipResponse)(nil), "go.micro.srv.flag.FlipResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.flag.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.srv.flag.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.srv.flag.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.srv.flag.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Flag service

type FlagClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Flip(ctx context.Context, in *FlipRequest, opts ...client.CallOption) (*FlipResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type flagClient struct {
	c           client.Client
	serviceName string
}

func NewFlagClient(serviceName string, c client.Client) FlagClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.flag"
	}
	return &flagClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *flagClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Flag.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flagClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Flag.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flagClient) Flip(ctx context.Context, in *FlipRequest, opts ...client.CallOption) (*FlipResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Flag.Flip", in)
	out := new(FlipResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flagClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Flag.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flagClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Flag.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Flag service

type FlagHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Flip(context.Context, *FlipRequest, *FlipResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterFlagHandler(s server.Server, hdlr FlagHandler) {
	s.Handle(s.NewHandler(&Flag{hdlr}))
}

type Flag struct {
	FlagHandler
}

func (h *Flag) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.FlagHandler.Create(ctx, in, out)
}

func (h *Flag) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.FlagHandler.Read(ctx, in, out)
}

func (h *Flag) Flip(ctx context.Context, in *FlipRequest, out *FlipResponse) error {
	return h.FlagHandler.Flip(ctx, in, out)
}

func (h *Flag) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.FlagHandler.Delete(ctx, in, out)
}

func (h *Flag) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.FlagHandler.List(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xcf, 0x4e, 0xf2, 0x40,
	0x10, 0x87, 0xaf, 0x7c, 0x44, 0xa7, 0x40, 0x70, 0xe3, 0x81, 0x78, 0x10, 0xe8, 0x89, 0xd3, 0x36,
	0xc1, 0x83, 0x0f, 0xa0, 0xc1, 0x98, 0x98, 0x98, 0xf4, 0xa4, 0xc7, 0x05, 0x46, 0xdc, 0x50, 0xd8,
	0xba, 0xbb, 0x25, 0xfa, 0x08, 0xbe, 0xb5, 0xdd, 0xee, 0x92, 0x14, 0xdd, 0xe2, 0xc5, 0x4b, 0xd3,
	0xe9, 0xcc, 0xfc, 0xe6, 0xf7, 0x27, 0x85, 0xe9, 0x8a, 0xeb, 0xd7, 0x7c, 0x4e, 0x17, 0x62, 0x13,
	0x27, 0x6c, 0xcd, 0xb6, 0xfc, 0x3d, 0x8f, 0x5f, 0x52, 0xb6, 0x8a, 0x95, 0xdc, 0xc5, 0x99, 0x14,
	0x5a, 0xd8, 0xd2, 0x3c, 0x68, 0x59, 0x93, 0xb3, 0x95, 0xa0, 0x1b, 0xbe, 0x90, 0x82, 0x16, 0x33,
	0xd4, 0x34, 0xa2, 0x67, 0xe8, 0xde, 0x48, 0x64, 0x1a, 0x13, 0x7c, 0xcb, 0x51, 0x69, 0xd2, 0x87,
	0x60, 0x8d, 0x1f, 0x83, 0xe6, 0xa8, 0x39, 0x39, 0x4d, 0xcc, 0x2b, 0x19, 0x41, 0xb8, 0x44, 0xb5,
	0x90, 0x3c, 0xd3, 0x5c, 0x6c, 0x07, 0xff, 0xca, 0x4e, 0xf5, 0x13, 0x39, 0x87, 0xff, 0x3b, 0x96,
	0xe6, 0x38, 0x08, 0x8a, 0xde, 0x49, 0x62, 0x8b, 0xa8, 0x0f, 0xbd, 0x3d, 0xb4, 0xca, 0xc4, 0x56,
	0x61, 0x34, 0x84, 0x30, 0x41, 0xb6, 0xac, 0x3d, 0x15, 0x3d, 0x41, 0xc7, 0x0e, 0xd8, 0x85, 0x3f,
	0x24, 0x53, 0x9c, 0x9e, 0xa5, 0x3c, 0xab, 0x3f, 0xdd, 0x83, 0x8e, 0x1d, 0x70, 0x5c, 0xc7, 0xd0,
	0xbd, 0xc5, 0x14, 0x8f, 0x18, 0x63, 0x04, 0xee, 0x47, 0xdc, 0x52, 0x17, 0xc2, 0x07, 0xae, 0xb4,
	0x5b, 0x89, 0xee, 0xa0, 0x63, 0x4b, 0x27, 0xe7, 0x1a, 0xda, 0x12, 0x55, 0x9e, 0xea, 0x02, 0x25,
	0x98, 0x84, 0xd3, 0x21, 0xfd, 0x11, 0x08, 0xad, 0xea, 0x4f, 0xdc, 0xf8, 0xf4, 0x33, 0x80, 0xd6,
	0xac, 0xe8, 0x92, 0x47, 0x68, 0x5b, 0x4f, 0xc9, 0xc8, 0xb3, 0x7b, 0x90, 0xe4, 0xc5, 0xf8, 0xc8,
	0x84, 0xe3, 0xdb, 0x20, 0xf7, 0xd0, 0x32, 0x17, 0xc9, 0x65, 0x2d, 0x15, 0x0b, 0xf6, 0x1b, 0x55,
	0x0b, 0x65, 0x1c, 0xf4, 0x42, 0x55, 0xbc, 0xf7, 0x42, 0x1d, 0x58, 0xdf, 0x30, 0x32, 0xad, 0xb3,
	0x5e, 0x99, 0x07, 0xb9, 0x78, 0x65, 0x7e, 0x8b, 0xa5, 0xe4, 0x66, 0x92, 0xf0, 0x72, 0xab, 0x24,
	0xe6, 0xe5, 0x56, 0x8d, 0x30, 0x6a, 0xcc, 0xdb, 0xe5, 0xbf, 0x74, 0xf5, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x58, 0x52, 0x3a, 0xc7, 0x81, 0x03, 0x00, 0x00,
}
