// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: photo.proto

package photo

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

// PhotoerClient is the client API for Photoer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhotoerClient interface {
	SendPhoto(ctx context.Context, in *PhotoRequest, opts ...grpc.CallOption) (*PhotoReply, error)
}

type photoerClient struct {
	cc grpc.ClientConnInterface
}

func NewPhotoerClient(cc grpc.ClientConnInterface) PhotoerClient {
	return &photoerClient{cc}
}

func (c *photoerClient) SendPhoto(ctx context.Context, in *PhotoRequest, opts ...grpc.CallOption) (*PhotoReply, error) {
	out := new(PhotoReply)
	err := c.cc.Invoke(ctx, "/photo.Photoer/SendPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhotoerServer is the server API for Photoer service.
// All implementations must embed UnimplementedPhotoerServer
// for forward compatibility
type PhotoerServer interface {
	SendPhoto(context.Context, *PhotoRequest) (*PhotoReply, error)
	mustEmbedUnimplementedPhotoerServer()
}

// UnimplementedPhotoerServer must be embedded to have forward compatible implementations.
type UnimplementedPhotoerServer struct {
}

func (UnimplementedPhotoerServer) SendPhoto(context.Context, *PhotoRequest) (*PhotoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoto not implemented")
}
func (UnimplementedPhotoerServer) mustEmbedUnimplementedPhotoerServer() {}

// UnsafePhotoerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhotoerServer will
// result in compilation errors.
type UnsafePhotoerServer interface {
	mustEmbedUnimplementedPhotoerServer()
}

func RegisterPhotoerServer(s grpc.ServiceRegistrar, srv PhotoerServer) {
	s.RegisterService(&Photoer_ServiceDesc, srv)
}

func _Photoer_SendPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoerServer).SendPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/photo.Photoer/SendPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoerServer).SendPhoto(ctx, req.(*PhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Photoer_ServiceDesc is the grpc.ServiceDesc for Photoer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Photoer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "photo.Photoer",
	HandlerType: (*PhotoerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPhoto",
			Handler:    _Photoer_SendPhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "photo.proto",
}
