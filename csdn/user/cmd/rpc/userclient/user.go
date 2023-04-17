// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CancelFocusUserRequest    = user.CancelFocusUserRequest
	CancelFocusUserResponse   = user.CancelFocusUserResponse
	FocusUserRequest          = user.FocusUserRequest
	FocusUserResponse         = user.FocusUserResponse
	GenerateTokenRequest      = user.GenerateTokenRequest
	GenerateTokenResponse     = user.GenerateTokenResponse
	GetUserAddressRequest     = user.GetUserAddressRequest
	GetUserAddressResponse    = user.GetUserAddressResponse
	IsFocusUserRequest        = user.IsFocusUserRequest
	IsFocusUserResponse       = user.IsFocusUserResponse
	LoginRequest              = user.LoginRequest
	LoginResponse             = user.LoginResponse
	SmsRequest                = user.SmsRequest
	SmsResponse               = user.SmsResponse
	UpdateUserAddressRequest  = user.UpdateUserAddressRequest
	UpdateUserAddressResponse = user.UpdateUserAddressResponse
	UserAddress               = user.UserAddress
	UserAddressRequest        = user.UserAddressRequest
	UserAddressResponse       = user.UserAddressResponse
	UserCurrInfoRequest       = user.UserCurrInfoRequest
	UserCurrInfoResponse      = user.UserCurrInfoResponse
	UserFansListRequest       = user.UserFansListRequest
	UserFansListResponse      = user.UserFansListResponse
	UserFocus                 = user.UserFocus
	UserFocusListRequest      = user.UserFocusListRequest
	UserFocusListResponse     = user.UserFocusListResponse
	UserInfoEditRequest       = user.UserInfoEditRequest
	UserInfoEditResponse      = user.UserInfoEditResponse

	User interface {
		UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
		UserCurrInfo(ctx context.Context, in *UserCurrInfoRequest, opts ...grpc.CallOption) (*UserCurrInfoResponse, error)
		SendSmsCode(ctx context.Context, in *SmsRequest, opts ...grpc.CallOption) (*SmsResponse, error)
		IsFocueUser(ctx context.Context, in *IsFocusUserRequest, opts ...grpc.CallOption) (*IsFocusUserResponse, error)
		FocueUser(ctx context.Context, in *FocusUserRequest, opts ...grpc.CallOption) (*FocusUserResponse, error)
		CancelFocueUser(ctx context.Context, in *CancelFocusUserRequest, opts ...grpc.CallOption) (*CancelFocusUserResponse, error)
		UserFocusList(ctx context.Context, in *UserFocusListRequest, opts ...grpc.CallOption) (*UserFocusListResponse, error)
		UserFansList(ctx context.Context, in *UserFansListRequest, opts ...grpc.CallOption) (*UserFansListResponse, error)
		UserInfoEdit(ctx context.Context, in *UserInfoEditRequest, opts ...grpc.CallOption) (*UserInfoEditResponse, error)
		UserAddress(ctx context.Context, in *UserAddressRequest, opts ...grpc.CallOption) (*UserAddressResponse, error)
		UserUpdateAddress(ctx context.Context, in *UpdateUserAddressRequest, opts ...grpc.CallOption) (*UpdateUserAddressResponse, error)
		UserGetAddress(ctx context.Context, in *GetUserAddressRequest, opts ...grpc.CallOption) (*GetUserAddressResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

func (m *defaultUser) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUser) UserCurrInfo(ctx context.Context, in *UserCurrInfoRequest, opts ...grpc.CallOption) (*UserCurrInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserCurrInfo(ctx, in, opts...)
}

func (m *defaultUser) SendSmsCode(ctx context.Context, in *SmsRequest, opts ...grpc.CallOption) (*SmsResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SendSmsCode(ctx, in, opts...)
}

func (m *defaultUser) IsFocueUser(ctx context.Context, in *IsFocusUserRequest, opts ...grpc.CallOption) (*IsFocusUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.IsFocueUser(ctx, in, opts...)
}

func (m *defaultUser) FocueUser(ctx context.Context, in *FocusUserRequest, opts ...grpc.CallOption) (*FocusUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.FocueUser(ctx, in, opts...)
}

func (m *defaultUser) CancelFocueUser(ctx context.Context, in *CancelFocusUserRequest, opts ...grpc.CallOption) (*CancelFocusUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CancelFocueUser(ctx, in, opts...)
}

func (m *defaultUser) UserFocusList(ctx context.Context, in *UserFocusListRequest, opts ...grpc.CallOption) (*UserFocusListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserFocusList(ctx, in, opts...)
}

func (m *defaultUser) UserFansList(ctx context.Context, in *UserFansListRequest, opts ...grpc.CallOption) (*UserFansListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserFansList(ctx, in, opts...)
}

func (m *defaultUser) UserInfoEdit(ctx context.Context, in *UserInfoEditRequest, opts ...grpc.CallOption) (*UserInfoEditResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfoEdit(ctx, in, opts...)
}

func (m *defaultUser) UserAddress(ctx context.Context, in *UserAddressRequest, opts ...grpc.CallOption) (*UserAddressResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserAddress(ctx, in, opts...)
}

func (m *defaultUser) UserUpdateAddress(ctx context.Context, in *UpdateUserAddressRequest, opts ...grpc.CallOption) (*UpdateUserAddressResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserUpdateAddress(ctx, in, opts...)
}

func (m *defaultUser) UserGetAddress(ctx context.Context, in *GetUserAddressRequest, opts ...grpc.CallOption) (*GetUserAddressResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserGetAddress(ctx, in, opts...)
}
