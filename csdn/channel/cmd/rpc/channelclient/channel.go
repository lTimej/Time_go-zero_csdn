// Code generated by goctl. DO NOT EDIT.
// Source: channel.proto

package channelclient

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ArticleChannelRequest        = channel.ArticleChannelRequest
	ArticleChannelResponse       = channel.ArticleChannelResponse
	ArticleLikeRequest           = channel.ArticleLikeRequest
	ArticleLikeResponse          = channel.ArticleLikeResponse
	ArticleLikeResponse_UserInfo = channel.ArticleLikeResponse_UserInfo
	ArticleList                  = channel.ArticleList
	ArticleReadRequest           = channel.ArticleReadRequest
	ArticleReadResponse          = channel.ArticleReadResponse
	ArticleToDisLikeRequest      = channel.ArticleToDisLikeRequest
	ArticleToDisLikeResponse     = channel.ArticleToDisLikeResponse
	ArticleToLikeRequest         = channel.ArticleToLikeRequest
	ArticleToLikeResponse        = channel.ArticleToLikeResponse
	ArticlestatusRequest         = channel.ArticlestatusRequest
	ArticlestatusResponse        = channel.ArticlestatusResponse
	ChannelList                  = channel.ChannelList
	ChannelListRequest           = channel.ChannelListRequest
	ChannelListResponse          = channel.ChannelListResponse
	DefaultChannelRequest        = channel.DefaultChannelRequest
	DefaultChannelResponse       = channel.DefaultChannelResponse
	UserAddChannelRequest        = channel.UserAddChannelRequest
	UserAddChannelResponse       = channel.UserAddChannelResponse
	UserChannelRequest           = channel.UserChannelRequest
	UserChannelResponse          = channel.UserChannelResponse
	UserPatchChannelRequest      = channel.UserPatchChannelRequest
	UserPatchChannelResponse     = channel.UserPatchChannelResponse

	Channel interface {
		AllChannel(ctx context.Context, in *ChannelListRequest, opts ...grpc.CallOption) (*ChannelListResponse, error)
		DefaultChannel(ctx context.Context, in *DefaultChannelRequest, opts ...grpc.CallOption) (*DefaultChannelResponse, error)
		UserChannel(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*UserChannelResponse, error)
		UserAddChannel(ctx context.Context, in *UserAddChannelRequest, opts ...grpc.CallOption) (*UserAddChannelResponse, error)
		UserPatchChannel(ctx context.Context, in *UserPatchChannelRequest, opts ...grpc.CallOption) (*UserPatchChannelResponse, error)
		ArticleChannel(ctx context.Context, in *ArticleChannelRequest, opts ...grpc.CallOption) (*ArticleChannelResponse, error)
		ArticleStatus(ctx context.Context, in *ArticlestatusRequest, opts ...grpc.CallOption) (*ArticlestatusResponse, error)
		ArticleRead(ctx context.Context, in *ArticleReadRequest, opts ...grpc.CallOption) (*ArticleReadResponse, error)
		ArticleLike(ctx context.Context, in *ArticleLikeRequest, opts ...grpc.CallOption) (*ArticleLikeResponse, error)
		ArticleToLike(ctx context.Context, in *ArticleToLikeRequest, opts ...grpc.CallOption) (*ArticleToLikeResponse, error)
		ArticleToDisLike(ctx context.Context, in *ArticleToDisLikeRequest, opts ...grpc.CallOption) (*ArticleToDisLikeResponse, error)
	}

	defaultChannel struct {
		cli zrpc.Client
	}
)

func NewChannel(cli zrpc.Client) Channel {
	return &defaultChannel{
		cli: cli,
	}
}

func (m *defaultChannel) AllChannel(ctx context.Context, in *ChannelListRequest, opts ...grpc.CallOption) (*ChannelListResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.AllChannel(ctx, in, opts...)
}

func (m *defaultChannel) DefaultChannel(ctx context.Context, in *DefaultChannelRequest, opts ...grpc.CallOption) (*DefaultChannelResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.DefaultChannel(ctx, in, opts...)
}

func (m *defaultChannel) UserChannel(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*UserChannelResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.UserChannel(ctx, in, opts...)
}

func (m *defaultChannel) UserAddChannel(ctx context.Context, in *UserAddChannelRequest, opts ...grpc.CallOption) (*UserAddChannelResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.UserAddChannel(ctx, in, opts...)
}

func (m *defaultChannel) UserPatchChannel(ctx context.Context, in *UserPatchChannelRequest, opts ...grpc.CallOption) (*UserPatchChannelResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.UserPatchChannel(ctx, in, opts...)
}

func (m *defaultChannel) ArticleChannel(ctx context.Context, in *ArticleChannelRequest, opts ...grpc.CallOption) (*ArticleChannelResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.ArticleChannel(ctx, in, opts...)
}

func (m *defaultChannel) ArticleStatus(ctx context.Context, in *ArticlestatusRequest, opts ...grpc.CallOption) (*ArticlestatusResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.ArticleStatus(ctx, in, opts...)
}

func (m *defaultChannel) ArticleRead(ctx context.Context, in *ArticleReadRequest, opts ...grpc.CallOption) (*ArticleReadResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.ArticleRead(ctx, in, opts...)
}

func (m *defaultChannel) ArticleLike(ctx context.Context, in *ArticleLikeRequest, opts ...grpc.CallOption) (*ArticleLikeResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.ArticleLike(ctx, in, opts...)
}

func (m *defaultChannel) ArticleToLike(ctx context.Context, in *ArticleToLikeRequest, opts ...grpc.CallOption) (*ArticleToLikeResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.ArticleToLike(ctx, in, opts...)
}

func (m *defaultChannel) ArticleToDisLike(ctx context.Context, in *ArticleToDisLikeRequest, opts ...grpc.CallOption) (*ArticleToDisLikeResponse, error) {
	client := channel.NewChannelClient(m.cli.Conn())
	return client.ArticleToDisLike(ctx, in, opts...)
}
