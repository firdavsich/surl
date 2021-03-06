// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SurlClient is the client API for Surl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SurlClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type surlClient struct {
	cc grpc.ClientConnInterface
}

func NewSurlClient(cc grpc.ClientConnInterface) SurlClient {
	return &surlClient{cc}
}

func (c *surlClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/api.Surl/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surlClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.Surl/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SurlServer is the server API for Surl service.
// All implementations must embed UnimplementedSurlServer
// for forward compatibility
type SurlServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedSurlServer()
}

// UnimplementedSurlServer must be embedded to have forward compatible implementations.
type UnimplementedSurlServer struct {
}

func (UnimplementedSurlServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSurlServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSurlServer) mustEmbedUnimplementedSurlServer() {}

// UnsafeSurlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SurlServer will
// result in compilation errors.
type UnsafeSurlServer interface {
	mustEmbedUnimplementedSurlServer()
}

func RegisterSurlServer(s grpc.ServiceRegistrar, srv SurlServer) {
	s.RegisterService(&Surl_ServiceDesc, srv)
}

func _Surl_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurlServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Surl/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurlServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Surl_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurlServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Surl/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurlServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Surl_ServiceDesc is the grpc.ServiceDesc for Surl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Surl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Surl",
	HandlerType: (*SurlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Surl_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Surl_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "surl.proto",
}
