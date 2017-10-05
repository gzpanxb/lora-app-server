// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkServer.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateNetworkServerRequest struct {
	// Name of the network-server.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// hostname:ip of the network-server.
	Server string `protobuf:"bytes,2,opt,name=server" json:"server,omitempty"`
}

func (m *CreateNetworkServerRequest) Reset()                    { *m = CreateNetworkServerRequest{} }
func (m *CreateNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNetworkServerRequest) ProtoMessage()               {}
func (*CreateNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *CreateNetworkServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateNetworkServerRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type CreateNetworkServerResponse struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateNetworkServerResponse) Reset()                    { *m = CreateNetworkServerResponse{} }
func (m *CreateNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNetworkServerResponse) ProtoMessage()               {}
func (*CreateNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *CreateNetworkServerResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNetworkServerRequest struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetNetworkServerRequest) Reset()                    { *m = GetNetworkServerRequest{} }
func (m *GetNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNetworkServerRequest) ProtoMessage()               {}
func (*GetNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *GetNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNetworkServerResponse struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Timestamp when the record was created.
	CreatedAt string `protobuf:"bytes,2,opt,name=createdAt" json:"createdAt,omitempty"`
	// Timestamp when the record was last updated.
	UpdatedAt string `protobuf:"bytes,3,opt,name=updatedAt" json:"updatedAt,omitempty"`
	// Name of the network-server.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// hostname:ip of the network-server.
	Server string `protobuf:"bytes,5,opt,name=server" json:"server,omitempty"`
}

func (m *GetNetworkServerResponse) Reset()                    { *m = GetNetworkServerResponse{} }
func (m *GetNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNetworkServerResponse) ProtoMessage()               {}
func (*GetNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *GetNetworkServerResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetNetworkServerResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetNetworkServerResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *GetNetworkServerResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetNetworkServerResponse) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type UpdateNetworkServerRequest struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the network-server.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// hostname:ip of the network-server.
	Server string `protobuf:"bytes,3,opt,name=server" json:"server,omitempty"`
}

func (m *UpdateNetworkServerRequest) Reset()                    { *m = UpdateNetworkServerRequest{} }
func (m *UpdateNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNetworkServerRequest) ProtoMessage()               {}
func (*UpdateNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *UpdateNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateNetworkServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateNetworkServerRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type UpdateNetworkServerResponse struct {
}

func (m *UpdateNetworkServerResponse) Reset()                    { *m = UpdateNetworkServerResponse{} }
func (m *UpdateNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateNetworkServerResponse) ProtoMessage()               {}
func (*UpdateNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

type DeleteNetworkServerRequest struct {
	// ID of the network-server.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteNetworkServerRequest) Reset()                    { *m = DeleteNetworkServerRequest{} }
func (m *DeleteNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNetworkServerRequest) ProtoMessage()               {}
func (*DeleteNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *DeleteNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteNetworkServerResponse struct {
}

func (m *DeleteNetworkServerResponse) Reset()                    { *m = DeleteNetworkServerResponse{} }
func (m *DeleteNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNetworkServerResponse) ProtoMessage()               {}
func (*DeleteNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

type ListNetworkServerRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListNetworkServerRequest) Reset()                    { *m = ListNetworkServerRequest{} }
func (m *ListNetworkServerRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNetworkServerRequest) ProtoMessage()               {}
func (*ListNetworkServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *ListNetworkServerRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListNetworkServerRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListNetworkServerResponse struct {
	// Total number of network-servers.
	TotalCount int64 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// Network-servers within the result-set.
	Result []*GetNetworkServerResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListNetworkServerResponse) Reset()                    { *m = ListNetworkServerResponse{} }
func (m *ListNetworkServerResponse) String() string            { return proto.CompactTextString(m) }
func (*ListNetworkServerResponse) ProtoMessage()               {}
func (*ListNetworkServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *ListNetworkServerResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListNetworkServerResponse) GetResult() []*GetNetworkServerResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateNetworkServerRequest)(nil), "api.CreateNetworkServerRequest")
	proto.RegisterType((*CreateNetworkServerResponse)(nil), "api.CreateNetworkServerResponse")
	proto.RegisterType((*GetNetworkServerRequest)(nil), "api.GetNetworkServerRequest")
	proto.RegisterType((*GetNetworkServerResponse)(nil), "api.GetNetworkServerResponse")
	proto.RegisterType((*UpdateNetworkServerRequest)(nil), "api.UpdateNetworkServerRequest")
	proto.RegisterType((*UpdateNetworkServerResponse)(nil), "api.UpdateNetworkServerResponse")
	proto.RegisterType((*DeleteNetworkServerRequest)(nil), "api.DeleteNetworkServerRequest")
	proto.RegisterType((*DeleteNetworkServerResponse)(nil), "api.DeleteNetworkServerResponse")
	proto.RegisterType((*ListNetworkServerRequest)(nil), "api.ListNetworkServerRequest")
	proto.RegisterType((*ListNetworkServerResponse)(nil), "api.ListNetworkServerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkServer service

type NetworkServerClient interface {
	// Create creates the given network-server.
	Create(ctx context.Context, in *CreateNetworkServerRequest, opts ...grpc.CallOption) (*CreateNetworkServerResponse, error)
	// Get returns the network-server matching the given id.
	Get(ctx context.Context, in *GetNetworkServerRequest, opts ...grpc.CallOption) (*GetNetworkServerResponse, error)
	// Update updates the given network-server.
	Update(ctx context.Context, in *UpdateNetworkServerRequest, opts ...grpc.CallOption) (*UpdateNetworkServerResponse, error)
	// Delete deletes the network-server matching the given id.
	Delete(ctx context.Context, in *DeleteNetworkServerRequest, opts ...grpc.CallOption) (*DeleteNetworkServerResponse, error)
	// List lists the available network-servers.
	List(ctx context.Context, in *ListNetworkServerRequest, opts ...grpc.CallOption) (*ListNetworkServerResponse, error)
}

type networkServerClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServerClient(cc *grpc.ClientConn) NetworkServerClient {
	return &networkServerClient{cc}
}

func (c *networkServerClient) Create(ctx context.Context, in *CreateNetworkServerRequest, opts ...grpc.CallOption) (*CreateNetworkServerResponse, error) {
	out := new(CreateNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Get(ctx context.Context, in *GetNetworkServerRequest, opts ...grpc.CallOption) (*GetNetworkServerResponse, error) {
	out := new(GetNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Update(ctx context.Context, in *UpdateNetworkServerRequest, opts ...grpc.CallOption) (*UpdateNetworkServerResponse, error) {
	out := new(UpdateNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) Delete(ctx context.Context, in *DeleteNetworkServerRequest, opts ...grpc.CallOption) (*DeleteNetworkServerResponse, error) {
	out := new(DeleteNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) List(ctx context.Context, in *ListNetworkServerRequest, opts ...grpc.CallOption) (*ListNetworkServerResponse, error) {
	out := new(ListNetworkServerResponse)
	err := grpc.Invoke(ctx, "/api.NetworkServer/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkServer service

type NetworkServerServer interface {
	// Create creates the given network-server.
	Create(context.Context, *CreateNetworkServerRequest) (*CreateNetworkServerResponse, error)
	// Get returns the network-server matching the given id.
	Get(context.Context, *GetNetworkServerRequest) (*GetNetworkServerResponse, error)
	// Update updates the given network-server.
	Update(context.Context, *UpdateNetworkServerRequest) (*UpdateNetworkServerResponse, error)
	// Delete deletes the network-server matching the given id.
	Delete(context.Context, *DeleteNetworkServerRequest) (*DeleteNetworkServerResponse, error)
	// List lists the available network-servers.
	List(context.Context, *ListNetworkServerRequest) (*ListNetworkServerResponse, error)
}

func RegisterNetworkServerServer(s *grpc.Server, srv NetworkServerServer) {
	s.RegisterService(&_NetworkServer_serviceDesc, srv)
}

func _NetworkServer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Create(ctx, req.(*CreateNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Get(ctx, req.(*GetNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Update(ctx, req.(*UpdateNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).Delete(ctx, req.(*DeleteNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServer/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).List(ctx, req.(*ListNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.NetworkServer",
	HandlerType: (*NetworkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NetworkServer_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NetworkServer_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NetworkServer_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetworkServer_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetworkServer_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkServer.proto",
}

func init() { proto.RegisterFile("networkServer.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0xd5, 0x26, 0x69, 0xa4, 0x0e, 0x82, 0x83, 0xa9, 0x4a, 0x36, 0x9b, 0x6e, 0x4b, 0xc4, 0xa1,
	0x54, 0x74, 0x57, 0x2a, 0xe2, 0xc2, 0x0d, 0x15, 0xa9, 0x3d, 0x20, 0x0e, 0x41, 0x48, 0x5c, 0x4d,
	0x33, 0xad, 0x2c, 0x52, 0x3b, 0x8d, 0x1d, 0x38, 0x20, 0x2e, 0x7c, 0x02, 0xfc, 0x06, 0x7f, 0xc3,
	0x2f, 0xf0, 0x21, 0x28, 0xf6, 0xb0, 0x74, 0xc1, 0x8e, 0xf6, 0xb6, 0x9e, 0x79, 0x79, 0x6f, 0x9e,
	0xe7, 0x79, 0xe1, 0xbe, 0x44, 0xf3, 0x49, 0x75, 0x1f, 0xde, 0x60, 0xf7, 0x11, 0xbb, 0x45, 0xdb,
	0x29, 0xa3, 0x58, 0xcc, 0x5b, 0x91, 0x17, 0x57, 0x4a, 0x5d, 0x35, 0xb8, 0xe4, 0xad, 0x58, 0x72,
	0x29, 0x95, 0xe1, 0x46, 0x28, 0xa9, 0x1d, 0xa4, 0x3c, 0x87, 0xfc, 0xb4, 0x43, 0x6e, 0xf0, 0xf5,
	0xed, 0xef, 0x2b, 0xbc, 0xe9, 0x51, 0x1b, 0xc6, 0x20, 0x91, 0xfc, 0x1a, 0xb3, 0xc9, 0xc1, 0xe4,
	0x70, 0xbb, 0xb2, 0xbf, 0xd9, 0x2e, 0xa4, 0xda, 0x82, 0xb2, 0xc8, 0x56, 0xe9, 0x54, 0x1e, 0xc3,
	0xcc, 0xcb, 0xa4, 0x5b, 0x25, 0x35, 0xb2, 0x7b, 0x10, 0x89, 0xda, 0x12, 0xc5, 0x55, 0x24, 0xea,
	0xf2, 0x31, 0x3c, 0x38, 0x43, 0xe3, 0x55, 0xfd, 0x17, 0xfa, 0x6d, 0x02, 0xd9, 0xff, 0x58, 0x3f,
	0x2f, 0x2b, 0x60, 0xfb, 0xc2, 0x8e, 0x51, 0xbf, 0x30, 0x34, 0xe1, 0xdf, 0xc2, 0xd0, 0xed, 0xdb,
	0x9a, 0xba, 0xb1, 0xeb, 0xae, 0x0a, 0x2b, 0xbb, 0x89, 0xd7, 0xee, 0xd6, 0x9a, 0xdd, 0x77, 0x90,
	0xbf, 0xb5, 0x1f, 0x6e, 0x62, 0x61, 0xc5, 0x1c, 0x79, 0x99, 0xe3, 0x35, 0xe6, 0x3d, 0x98, 0x79,
	0x99, 0x9d, 0xe1, 0xf2, 0x09, 0xe4, 0x2f, 0xb1, 0xc1, 0xcd, 0x84, 0x07, 0x32, 0x2f, 0x9a, 0xc8,
	0xce, 0x21, 0x7b, 0x25, 0xb4, 0x7f, 0x0d, 0x3b, 0xb0, 0xd5, 0x88, 0x6b, 0x61, 0x88, 0xcd, 0x1d,
	0x86, 0xa9, 0xd5, 0xe5, 0xa5, 0x46, 0x77, 0xb9, 0x71, 0x45, 0xa7, 0xb2, 0x83, 0xa9, 0x87, 0x89,
	0x96, 0x34, 0x07, 0x30, 0xca, 0xf0, 0xe6, 0x54, 0xf5, 0xf2, 0x0f, 0xdf, 0xad, 0x0a, 0x7b, 0x06,
	0x69, 0x87, 0xba, 0x6f, 0x06, 0xd2, 0xf8, 0xf0, 0xce, 0xc9, 0xde, 0x82, 0xb7, 0x62, 0x11, 0xda,
	0x79, 0x45, 0xe0, 0x93, 0x1f, 0x09, 0xdc, 0x5d, 0x43, 0xb0, 0x06, 0x52, 0x17, 0x42, 0xb6, 0x6f,
	0x29, 0xc2, 0xd9, 0xce, 0x0f, 0xc2, 0x00, 0xba, 0x9c, 0xfd, 0xaf, 0x3f, 0x7f, 0x7d, 0x8f, 0xa6,
	0xe5, 0x8e, 0x7d, 0x3b, 0xf4, 0xc0, 0x8e, 0xdd, 0x96, 0xf4, 0xf3, 0xc9, 0x11, 0x43, 0x88, 0xcf,
	0xd0, 0xb0, 0x22, 0x30, 0xad, 0xd3, 0x19, 0xf7, 0x52, 0x3e, 0xb4, 0x22, 0x33, 0x36, 0xf5, 0x89,
	0x2c, 0x3f, 0x8b, 0xfa, 0x0b, 0xbb, 0x81, 0xd4, 0x05, 0x82, 0x4c, 0x85, 0x73, 0x47, 0xa6, 0xc6,
	0xe2, 0xf3, 0xc8, 0xea, 0xcd, 0xf3, 0xb0, 0xde, 0xe0, 0x4c, 0x42, 0xea, 0x62, 0x43, 0x92, 0xe1,
	0xc4, 0x91, 0xe4, 0x58, 0xc8, 0xc8, 0xe2, 0xd1, 0x88, 0xc5, 0x0b, 0x48, 0x86, 0xf4, 0x30, 0x77,
	0x59, 0xa1, 0x48, 0xe6, 0xf3, 0x50, 0x9b, 0x94, 0x0a, 0xab, 0xb4, 0xcb, 0xbc, 0x1b, 0x7b, 0x9f,
	0xda, 0xbf, 0xbc, 0xa7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x95, 0xe4, 0xe0, 0x2c, 0x05,
	0x00, 0x00,
}
