// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user.proto

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
	UserCurrInfo(ctx context.Context, in *UserCurrInfoRequest, opts ...grpc.CallOption) (*UserCurrInfoResponse, error)
	SendSmsCode(ctx context.Context, in *SmsRequest, opts ...grpc.CallOption) (*SmsResponse, error)
	IsFocueUser(ctx context.Context, in *IsFocusUserRequest, opts ...grpc.CallOption) (*IsFocusUserResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.User/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	out := new(GenerateTokenResponse)
	err := c.cc.Invoke(ctx, "/user.User/generateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserCurrInfo(ctx context.Context, in *UserCurrInfoRequest, opts ...grpc.CallOption) (*UserCurrInfoResponse, error) {
	out := new(UserCurrInfoResponse)
	err := c.cc.Invoke(ctx, "/user.User/UserCurrInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SendSmsCode(ctx context.Context, in *SmsRequest, opts ...grpc.CallOption) (*SmsResponse, error) {
	out := new(SmsResponse)
	err := c.cc.Invoke(ctx, "/user.User/SendSmsCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) IsFocueUser(ctx context.Context, in *IsFocusUserRequest, opts ...grpc.CallOption) (*IsFocusUserResponse, error) {
	out := new(IsFocusUserResponse)
	err := c.cc.Invoke(ctx, "/user.User/IsFocueUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	UserLogin(context.Context, *LoginRequest) (*LoginResponse, error)
	GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error)
	UserCurrInfo(context.Context, *UserCurrInfoRequest) (*UserCurrInfoResponse, error)
	SendSmsCode(context.Context, *SmsRequest) (*SmsResponse, error)
	IsFocueUser(context.Context, *IsFocusUserRequest) (*IsFocusUserResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) UserLogin(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServer) GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedUserServer) UserCurrInfo(context.Context, *UserCurrInfoRequest) (*UserCurrInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCurrInfo not implemented")
}
func (UnimplementedUserServer) SendSmsCode(context.Context, *SmsRequest) (*SmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmsCode not implemented")
}
func (UnimplementedUserServer) IsFocueUser(context.Context, *IsFocusUserRequest) (*IsFocusUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFocueUser not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/generateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GenerateToken(ctx, req.(*GenerateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserCurrInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCurrInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserCurrInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserCurrInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserCurrInfo(ctx, req.(*UserCurrInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SendSmsCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SendSmsCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SendSmsCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SendSmsCode(ctx, req.(*SmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_IsFocueUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFocusUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).IsFocueUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/IsFocueUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).IsFocueUser(ctx, req.(*IsFocusUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "generateToken",
			Handler:    _User_GenerateToken_Handler,
		},
		{
			MethodName: "UserCurrInfo",
			Handler:    _User_UserCurrInfo_Handler,
		},
		{
			MethodName: "SendSmsCode",
			Handler:    _User_SendSmsCode_Handler,
		},
		{
			MethodName: "IsFocueUser",
			Handler:    _User_IsFocueUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
