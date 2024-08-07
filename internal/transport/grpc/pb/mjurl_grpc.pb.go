// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: mjurl.proto

package pb

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

// MJUrlClient is the client API for MJUrl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MJUrlClient interface {
	GetURL(ctx context.Context, in *ShortURL, opts ...grpc.CallOption) (*LongURL, error)
	ShortenURL(ctx context.Context, in *LongURL, opts ...grpc.CallOption) (*ShortURL, error)
}

type mJUrlClient struct {
	cc grpc.ClientConnInterface
}

func NewMJUrlClient(cc grpc.ClientConnInterface) MJUrlClient {
	return &mJUrlClient{cc}
}

func (c *mJUrlClient) GetURL(ctx context.Context, in *ShortURL, opts ...grpc.CallOption) (*LongURL, error) {
	out := new(LongURL)
	err := c.cc.Invoke(ctx, "/MJUrl/GetURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mJUrlClient) ShortenURL(ctx context.Context, in *LongURL, opts ...grpc.CallOption) (*ShortURL, error) {
	out := new(ShortURL)
	err := c.cc.Invoke(ctx, "/MJUrl/ShortenURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MJUrlServer is the server API for MJUrl service.
// All implementations must embed UnimplementedMJUrlServer
// for forward compatibility
type MJUrlServer interface {
	GetURL(context.Context, *ShortURL) (*LongURL, error)
	ShortenURL(context.Context, *LongURL) (*ShortURL, error)
	mustEmbedUnimplementedMJUrlServer()
}

// UnimplementedMJUrlServer must be embedded to have forward compatible implementations.
type UnimplementedMJUrlServer struct {
}

func (UnimplementedMJUrlServer) GetURL(context.Context, *ShortURL) (*LongURL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURL not implemented")
}
func (UnimplementedMJUrlServer) ShortenURL(context.Context, *LongURL) (*ShortURL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortenURL not implemented")
}
func (UnimplementedMJUrlServer) mustEmbedUnimplementedMJUrlServer() {}

// UnsafeMJUrlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MJUrlServer will
// result in compilation errors.
type UnsafeMJUrlServer interface {
	mustEmbedUnimplementedMJUrlServer()
}

func RegisterMJUrlServer(s grpc.ServiceRegistrar, srv MJUrlServer) {
	s.RegisterService(&MJUrl_ServiceDesc, srv)
}

func _MJUrl_GetURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortURL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MJUrlServer).GetURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MJUrl/GetURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MJUrlServer).GetURL(ctx, req.(*ShortURL))
	}
	return interceptor(ctx, in, info, handler)
}

func _MJUrl_ShortenURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LongURL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MJUrlServer).ShortenURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MJUrl/ShortenURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MJUrlServer).ShortenURL(ctx, req.(*LongURL))
	}
	return interceptor(ctx, in, info, handler)
}

// MJUrl_ServiceDesc is the grpc.ServiceDesc for MJUrl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MJUrl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MJUrl",
	HandlerType: (*MJUrlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetURL",
			Handler:    _MJUrl_GetURL_Handler,
		},
		{
			MethodName: "ShortenURL",
			Handler:    _MJUrl_ShortenURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mjurl.proto",
}
