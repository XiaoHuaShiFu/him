// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authz

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

// AuthzServiceClient is the client API for AuthzService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthzServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	SendSmVerCodeForLogin(ctx context.Context, in *SendSmVerCodeForLoginReq, opts ...grpc.CallOption) (*SendSmVerCodeForLoginResp, error)
	SetPwdLogin(ctx context.Context, in *SetPwdLoginReq, opts ...grpc.CallOption) (*SetPwdLoginResp, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error)
}

type authzServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzServiceClient(cc grpc.ClientConnInterface) AuthzServiceClient {
	return &authzServiceClient{cc}
}

func (c *authzServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/authnz.authz.AuthzService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authzServiceClient) SendSmVerCodeForLogin(ctx context.Context, in *SendSmVerCodeForLoginReq, opts ...grpc.CallOption) (*SendSmVerCodeForLoginResp, error) {
	out := new(SendSmVerCodeForLoginResp)
	err := c.cc.Invoke(ctx, "/authnz.authz.AuthzService/SendSmVerCodeForLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authzServiceClient) SetPwdLogin(ctx context.Context, in *SetPwdLoginReq, opts ...grpc.CallOption) (*SetPwdLoginResp, error) {
	out := new(SetPwdLoginResp)
	err := c.cc.Invoke(ctx, "/authnz.authz.AuthzService/SetPwdLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authzServiceClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error) {
	out := new(LogoutResp)
	err := c.cc.Invoke(ctx, "/authnz.authz.AuthzService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzServiceServer is the server API for AuthzService service.
// All implementations must embed UnimplementedAuthzServiceServer
// for forward compatibility
type AuthzServiceServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	SendSmVerCodeForLogin(context.Context, *SendSmVerCodeForLoginReq) (*SendSmVerCodeForLoginResp, error)
	SetPwdLogin(context.Context, *SetPwdLoginReq) (*SetPwdLoginResp, error)
	Logout(context.Context, *LogoutReq) (*LogoutResp, error)
	mustEmbedUnimplementedAuthzServiceServer()
}

// UnimplementedAuthzServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthzServiceServer struct {
}

func (UnimplementedAuthzServiceServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthzServiceServer) SendSmVerCodeForLogin(context.Context, *SendSmVerCodeForLoginReq) (*SendSmVerCodeForLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmVerCodeForLogin not implemented")
}
func (UnimplementedAuthzServiceServer) SetPwdLogin(context.Context, *SetPwdLoginReq) (*SetPwdLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPwdLogin not implemented")
}
func (UnimplementedAuthzServiceServer) Logout(context.Context, *LogoutReq) (*LogoutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthzServiceServer) mustEmbedUnimplementedAuthzServiceServer() {}

// UnsafeAuthzServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthzServiceServer will
// result in compilation errors.
type UnsafeAuthzServiceServer interface {
	mustEmbedUnimplementedAuthzServiceServer()
}

func RegisterAuthzServiceServer(s grpc.ServiceRegistrar, srv AuthzServiceServer) {
	s.RegisterService(&AuthzService_ServiceDesc, srv)
}

func _AuthzService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authnz.authz.AuthzService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthzService_SendSmVerCodeForLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmVerCodeForLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServiceServer).SendSmVerCodeForLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authnz.authz.AuthzService/SendSmVerCodeForLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServiceServer).SendSmVerCodeForLogin(ctx, req.(*SendSmVerCodeForLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthzService_SetPwdLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPwdLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServiceServer).SetPwdLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authnz.authz.AuthzService/SetPwdLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServiceServer).SetPwdLogin(ctx, req.(*SetPwdLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthzService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authnz.authz.AuthzService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServiceServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthzService_ServiceDesc is the grpc.ServiceDesc for AuthzService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthzService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authnz.authz.AuthzService",
	HandlerType: (*AuthzServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthzService_Login_Handler,
		},
		{
			MethodName: "SendSmVerCodeForLogin",
			Handler:    _AuthzService_SendSmVerCodeForLogin_Handler,
		},
		{
			MethodName: "SetPwdLogin",
			Handler:    _AuthzService_SetPwdLogin_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthzService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/authnz/authz/authz.proto",
}