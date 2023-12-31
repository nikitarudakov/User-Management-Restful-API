// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: protos/user.proto

package user

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

// UserControllerClient is the client API for UserController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserControllerClient interface {
	DeleteUserAndProfile(ctx context.Context, in *Username, opts ...grpc.CallOption) (*Empty, error)
	ListProfiles(ctx context.Context, in *Page, opts ...grpc.CallOption) (*Profiles, error)
	UpdateUserAndProfile(ctx context.Context, in *UpdateArg, opts ...grpc.CallOption) (*Empty, error)
}

type userControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserControllerClient(cc grpc.ClientConnInterface) UserControllerClient {
	return &userControllerClient{cc}
}

func (c *userControllerClient) DeleteUserAndProfile(ctx context.Context, in *Username, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.UserController/DeleteUserAndProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) ListProfiles(ctx context.Context, in *Page, opts ...grpc.CallOption) (*Profiles, error) {
	out := new(Profiles)
	err := c.cc.Invoke(ctx, "/user.UserController/ListProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) UpdateUserAndProfile(ctx context.Context, in *UpdateArg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.UserController/UpdateUserAndProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserControllerServer is the server API for UserController service.
// All implementations must embed UnimplementedUserControllerServer
// for forward compatibility
type UserControllerServer interface {
	DeleteUserAndProfile(context.Context, *Username) (*Empty, error)
	ListProfiles(context.Context, *Page) (*Profiles, error)
	UpdateUserAndProfile(context.Context, *UpdateArg) (*Empty, error)
	mustEmbedUnimplementedUserControllerServer()
}

// UnimplementedUserControllerServer must be embedded to have forward compatible implementations.
type UnimplementedUserControllerServer struct {
}

func (UnimplementedUserControllerServer) DeleteUserAndProfile(context.Context, *Username) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAndProfile not implemented")
}
func (UnimplementedUserControllerServer) ListProfiles(context.Context, *Page) (*Profiles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProfiles not implemented")
}
func (UnimplementedUserControllerServer) UpdateUserAndProfile(context.Context, *UpdateArg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAndProfile not implemented")
}
func (UnimplementedUserControllerServer) mustEmbedUnimplementedUserControllerServer() {}

// UnsafeUserControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserControllerServer will
// result in compilation errors.
type UnsafeUserControllerServer interface {
	mustEmbedUnimplementedUserControllerServer()
}

func RegisterUserControllerServer(s grpc.ServiceRegistrar, srv UserControllerServer) {
	s.RegisterService(&UserController_ServiceDesc, srv)
}

func _UserController_DeleteUserAndProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).DeleteUserAndProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserController/DeleteUserAndProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).DeleteUserAndProfile(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_ListProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).ListProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserController/ListProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).ListProfiles(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_UpdateUserAndProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).UpdateUserAndProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserController/UpdateUserAndProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).UpdateUserAndProfile(ctx, req.(*UpdateArg))
	}
	return interceptor(ctx, in, info, handler)
}

// UserController_ServiceDesc is the grpc.ServiceDesc for UserController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserController",
	HandlerType: (*UserControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteUserAndProfile",
			Handler:    _UserController_DeleteUserAndProfile_Handler,
		},
		{
			MethodName: "ListProfiles",
			Handler:    _UserController_ListProfiles_Handler,
		},
		{
			MethodName: "UpdateUserAndProfile",
			Handler:    _UserController_UpdateUserAndProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/user.proto",
}
