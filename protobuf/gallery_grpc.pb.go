// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: gallery.proto

package proto

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

// GaleryServiceClient is the client API for GaleryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GaleryServiceClient interface {
	SaveLink(ctx context.Context, in *SaveLinkRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	GetLinkByUserId(ctx context.Context, in *GetLinkByUserIdRequest, opts ...grpc.CallOption) (*GetLinkByUserIdResponse, error)
	UpdateLink(ctx context.Context, in *UpdateLinkRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type galeryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGaleryServiceClient(cc grpc.ClientConnInterface) GaleryServiceClient {
	return &galeryServiceClient{cc}
}

func (c *galeryServiceClient) SaveLink(ctx context.Context, in *SaveLinkRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/GaleryService/SaveLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galeryServiceClient) GetLinkByUserId(ctx context.Context, in *GetLinkByUserIdRequest, opts ...grpc.CallOption) (*GetLinkByUserIdResponse, error) {
	out := new(GetLinkByUserIdResponse)
	err := c.cc.Invoke(ctx, "/GaleryService/GetLinkByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galeryServiceClient) UpdateLink(ctx context.Context, in *UpdateLinkRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/GaleryService/UpdateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *galeryServiceClient) DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/GaleryService/DeleteLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GaleryServiceServer is the server API for GaleryService service.
// All implementations must embed UnimplementedGaleryServiceServer
// for forward compatibility
type GaleryServiceServer interface {
	SaveLink(context.Context, *SaveLinkRequest) (*CommonResponse, error)
	GetLinkByUserId(context.Context, *GetLinkByUserIdRequest) (*GetLinkByUserIdResponse, error)
	UpdateLink(context.Context, *UpdateLinkRequest) (*CommonResponse, error)
	DeleteLink(context.Context, *DeleteLinkRequest) (*CommonResponse, error)
	mustEmbedUnimplementedGaleryServiceServer()
}

// UnimplementedGaleryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGaleryServiceServer struct {
}

func (UnimplementedGaleryServiceServer) SaveLink(context.Context, *SaveLinkRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLink not implemented")
}
func (UnimplementedGaleryServiceServer) GetLinkByUserId(context.Context, *GetLinkByUserIdRequest) (*GetLinkByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkByUserId not implemented")
}
func (UnimplementedGaleryServiceServer) UpdateLink(context.Context, *UpdateLinkRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLink not implemented")
}
func (UnimplementedGaleryServiceServer) DeleteLink(context.Context, *DeleteLinkRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLink not implemented")
}
func (UnimplementedGaleryServiceServer) mustEmbedUnimplementedGaleryServiceServer() {}

// UnsafeGaleryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GaleryServiceServer will
// result in compilation errors.
type UnsafeGaleryServiceServer interface {
	mustEmbedUnimplementedGaleryServiceServer()
}

func RegisterGaleryServiceServer(s grpc.ServiceRegistrar, srv GaleryServiceServer) {
	s.RegisterService(&GaleryService_ServiceDesc, srv)
}

func _GaleryService_SaveLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaleryServiceServer).SaveLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GaleryService/SaveLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaleryServiceServer).SaveLink(ctx, req.(*SaveLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GaleryService_GetLinkByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaleryServiceServer).GetLinkByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GaleryService/GetLinkByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaleryServiceServer).GetLinkByUserId(ctx, req.(*GetLinkByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GaleryService_UpdateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaleryServiceServer).UpdateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GaleryService/UpdateLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaleryServiceServer).UpdateLink(ctx, req.(*UpdateLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GaleryService_DeleteLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaleryServiceServer).DeleteLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GaleryService/DeleteLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaleryServiceServer).DeleteLink(ctx, req.(*DeleteLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GaleryService_ServiceDesc is the grpc.ServiceDesc for GaleryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GaleryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GaleryService",
	HandlerType: (*GaleryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveLink",
			Handler:    _GaleryService_SaveLink_Handler,
		},
		{
			MethodName: "GetLinkByUserId",
			Handler:    _GaleryService_GetLinkByUserId_Handler,
		},
		{
			MethodName: "UpdateLink",
			Handler:    _GaleryService_UpdateLink_Handler,
		},
		{
			MethodName: "DeleteLink",
			Handler:    _GaleryService_DeleteLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gallery.proto",
}
