// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: channel.proto

package channel

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

// ChannelClient is the client API for Channel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelClient interface {
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
	ArticleStatusCache(ctx context.Context, in *ArticleStatusCacheRequest, opts ...grpc.CallOption) (*ArticleStatusCacheResponse, error)
	ArticleToCollection(ctx context.Context, in *ArticleToCollectionRequest, opts ...grpc.CallOption) (*ArticleToCollectionResponse, error)
	ArticleToDisCollection(ctx context.Context, in *ArticleToDisCollectionRequest, opts ...grpc.CallOption) (*ArticleToDisCollectionResponse, error)
	ArticleUserCollection(ctx context.Context, in *ArticleUserCollectionRequest, opts ...grpc.CallOption) (*ArticleUserCollectionResponse, error)
	ArticleToComment(ctx context.Context, in *ArticleToCommnetRequest, opts ...grpc.CallOption) (*ArticleToCommentResponse, error)
	ArticleCommentList(ctx context.Context, in *ArticleCommentListRequest, opts ...grpc.CallOption) (*ArticleCommentListResponse, error)
}

type channelClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelClient(cc grpc.ClientConnInterface) ChannelClient {
	return &channelClient{cc}
}

func (c *channelClient) AllChannel(ctx context.Context, in *ChannelListRequest, opts ...grpc.CallOption) (*ChannelListResponse, error) {
	out := new(ChannelListResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/AllChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) DefaultChannel(ctx context.Context, in *DefaultChannelRequest, opts ...grpc.CallOption) (*DefaultChannelResponse, error) {
	out := new(DefaultChannelResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/DefaultChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) UserChannel(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*UserChannelResponse, error) {
	out := new(UserChannelResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/UserChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) UserAddChannel(ctx context.Context, in *UserAddChannelRequest, opts ...grpc.CallOption) (*UserAddChannelResponse, error) {
	out := new(UserAddChannelResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/UserAddChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) UserPatchChannel(ctx context.Context, in *UserPatchChannelRequest, opts ...grpc.CallOption) (*UserPatchChannelResponse, error) {
	out := new(UserPatchChannelResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/UserPatchChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleChannel(ctx context.Context, in *ArticleChannelRequest, opts ...grpc.CallOption) (*ArticleChannelResponse, error) {
	out := new(ArticleChannelResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleStatus(ctx context.Context, in *ArticlestatusRequest, opts ...grpc.CallOption) (*ArticlestatusResponse, error) {
	out := new(ArticlestatusResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleRead(ctx context.Context, in *ArticleReadRequest, opts ...grpc.CallOption) (*ArticleReadResponse, error) {
	out := new(ArticleReadResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleLike(ctx context.Context, in *ArticleLikeRequest, opts ...grpc.CallOption) (*ArticleLikeResponse, error) {
	out := new(ArticleLikeResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleToLike(ctx context.Context, in *ArticleToLikeRequest, opts ...grpc.CallOption) (*ArticleToLikeResponse, error) {
	out := new(ArticleToLikeResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleToLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleToDisLike(ctx context.Context, in *ArticleToDisLikeRequest, opts ...grpc.CallOption) (*ArticleToDisLikeResponse, error) {
	out := new(ArticleToDisLikeResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleToDisLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleStatusCache(ctx context.Context, in *ArticleStatusCacheRequest, opts ...grpc.CallOption) (*ArticleStatusCacheResponse, error) {
	out := new(ArticleStatusCacheResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleStatusCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleToCollection(ctx context.Context, in *ArticleToCollectionRequest, opts ...grpc.CallOption) (*ArticleToCollectionResponse, error) {
	out := new(ArticleToCollectionResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleToCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleToDisCollection(ctx context.Context, in *ArticleToDisCollectionRequest, opts ...grpc.CallOption) (*ArticleToDisCollectionResponse, error) {
	out := new(ArticleToDisCollectionResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleToDisCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleUserCollection(ctx context.Context, in *ArticleUserCollectionRequest, opts ...grpc.CallOption) (*ArticleUserCollectionResponse, error) {
	out := new(ArticleUserCollectionResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleUserCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleToComment(ctx context.Context, in *ArticleToCommnetRequest, opts ...grpc.CallOption) (*ArticleToCommentResponse, error) {
	out := new(ArticleToCommentResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleToComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) ArticleCommentList(ctx context.Context, in *ArticleCommentListRequest, opts ...grpc.CallOption) (*ArticleCommentListResponse, error) {
	out := new(ArticleCommentListResponse)
	err := c.cc.Invoke(ctx, "/channel.Channel/ArticleCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServer is the server API for Channel service.
// All implementations must embed UnimplementedChannelServer
// for forward compatibility
type ChannelServer interface {
	AllChannel(context.Context, *ChannelListRequest) (*ChannelListResponse, error)
	DefaultChannel(context.Context, *DefaultChannelRequest) (*DefaultChannelResponse, error)
	UserChannel(context.Context, *UserChannelRequest) (*UserChannelResponse, error)
	UserAddChannel(context.Context, *UserAddChannelRequest) (*UserAddChannelResponse, error)
	UserPatchChannel(context.Context, *UserPatchChannelRequest) (*UserPatchChannelResponse, error)
	ArticleChannel(context.Context, *ArticleChannelRequest) (*ArticleChannelResponse, error)
	ArticleStatus(context.Context, *ArticlestatusRequest) (*ArticlestatusResponse, error)
	ArticleRead(context.Context, *ArticleReadRequest) (*ArticleReadResponse, error)
	ArticleLike(context.Context, *ArticleLikeRequest) (*ArticleLikeResponse, error)
	ArticleToLike(context.Context, *ArticleToLikeRequest) (*ArticleToLikeResponse, error)
	ArticleToDisLike(context.Context, *ArticleToDisLikeRequest) (*ArticleToDisLikeResponse, error)
	ArticleStatusCache(context.Context, *ArticleStatusCacheRequest) (*ArticleStatusCacheResponse, error)
	ArticleToCollection(context.Context, *ArticleToCollectionRequest) (*ArticleToCollectionResponse, error)
	ArticleToDisCollection(context.Context, *ArticleToDisCollectionRequest) (*ArticleToDisCollectionResponse, error)
	ArticleUserCollection(context.Context, *ArticleUserCollectionRequest) (*ArticleUserCollectionResponse, error)
	ArticleToComment(context.Context, *ArticleToCommnetRequest) (*ArticleToCommentResponse, error)
	ArticleCommentList(context.Context, *ArticleCommentListRequest) (*ArticleCommentListResponse, error)
	mustEmbedUnimplementedChannelServer()
}

// UnimplementedChannelServer must be embedded to have forward compatible implementations.
type UnimplementedChannelServer struct {
}

func (UnimplementedChannelServer) AllChannel(context.Context, *ChannelListRequest) (*ChannelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllChannel not implemented")
}
func (UnimplementedChannelServer) DefaultChannel(context.Context, *DefaultChannelRequest) (*DefaultChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefaultChannel not implemented")
}
func (UnimplementedChannelServer) UserChannel(context.Context, *UserChannelRequest) (*UserChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserChannel not implemented")
}
func (UnimplementedChannelServer) UserAddChannel(context.Context, *UserAddChannelRequest) (*UserAddChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAddChannel not implemented")
}
func (UnimplementedChannelServer) UserPatchChannel(context.Context, *UserPatchChannelRequest) (*UserPatchChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserPatchChannel not implemented")
}
func (UnimplementedChannelServer) ArticleChannel(context.Context, *ArticleChannelRequest) (*ArticleChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleChannel not implemented")
}
func (UnimplementedChannelServer) ArticleStatus(context.Context, *ArticlestatusRequest) (*ArticlestatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleStatus not implemented")
}
func (UnimplementedChannelServer) ArticleRead(context.Context, *ArticleReadRequest) (*ArticleReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleRead not implemented")
}
func (UnimplementedChannelServer) ArticleLike(context.Context, *ArticleLikeRequest) (*ArticleLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleLike not implemented")
}
func (UnimplementedChannelServer) ArticleToLike(context.Context, *ArticleToLikeRequest) (*ArticleToLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleToLike not implemented")
}
func (UnimplementedChannelServer) ArticleToDisLike(context.Context, *ArticleToDisLikeRequest) (*ArticleToDisLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleToDisLike not implemented")
}
func (UnimplementedChannelServer) ArticleStatusCache(context.Context, *ArticleStatusCacheRequest) (*ArticleStatusCacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleStatusCache not implemented")
}
func (UnimplementedChannelServer) ArticleToCollection(context.Context, *ArticleToCollectionRequest) (*ArticleToCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleToCollection not implemented")
}
func (UnimplementedChannelServer) ArticleToDisCollection(context.Context, *ArticleToDisCollectionRequest) (*ArticleToDisCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleToDisCollection not implemented")
}
func (UnimplementedChannelServer) ArticleUserCollection(context.Context, *ArticleUserCollectionRequest) (*ArticleUserCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleUserCollection not implemented")
}
func (UnimplementedChannelServer) ArticleToComment(context.Context, *ArticleToCommnetRequest) (*ArticleToCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleToComment not implemented")
}
func (UnimplementedChannelServer) ArticleCommentList(context.Context, *ArticleCommentListRequest) (*ArticleCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleCommentList not implemented")
}
func (UnimplementedChannelServer) mustEmbedUnimplementedChannelServer() {}

// UnsafeChannelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelServer will
// result in compilation errors.
type UnsafeChannelServer interface {
	mustEmbedUnimplementedChannelServer()
}

func RegisterChannelServer(s grpc.ServiceRegistrar, srv ChannelServer) {
	s.RegisterService(&Channel_ServiceDesc, srv)
}

func _Channel_AllChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).AllChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/AllChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).AllChannel(ctx, req.(*ChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_DefaultChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).DefaultChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/DefaultChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).DefaultChannel(ctx, req.(*DefaultChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_UserChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).UserChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/UserChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).UserChannel(ctx, req.(*UserChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_UserAddChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).UserAddChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/UserAddChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).UserAddChannel(ctx, req.(*UserAddChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_UserPatchChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPatchChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).UserPatchChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/UserPatchChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).UserPatchChannel(ctx, req.(*UserPatchChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleChannel(ctx, req.(*ArticleChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticlestatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleStatus(ctx, req.(*ArticlestatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleRead(ctx, req.(*ArticleReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleLike(ctx, req.(*ArticleLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleToLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleToLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleToLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleToLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleToLike(ctx, req.(*ArticleToLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleToDisLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleToDisLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleToDisLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleToDisLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleToDisLike(ctx, req.(*ArticleToDisLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleStatusCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleStatusCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleStatusCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleStatusCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleStatusCache(ctx, req.(*ArticleStatusCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleToCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleToCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleToCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleToCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleToCollection(ctx, req.(*ArticleToCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleToDisCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleToDisCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleToDisCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleToDisCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleToDisCollection(ctx, req.(*ArticleToDisCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleUserCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleUserCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleUserCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleUserCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleUserCollection(ctx, req.(*ArticleUserCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleToComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleToCommnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleToComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleToComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleToComment(ctx, req.(*ArticleToCommnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_ArticleCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).ArticleCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/ArticleCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).ArticleCommentList(ctx, req.(*ArticleCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Channel_ServiceDesc is the grpc.ServiceDesc for Channel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Channel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.Channel",
	HandlerType: (*ChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllChannel",
			Handler:    _Channel_AllChannel_Handler,
		},
		{
			MethodName: "DefaultChannel",
			Handler:    _Channel_DefaultChannel_Handler,
		},
		{
			MethodName: "UserChannel",
			Handler:    _Channel_UserChannel_Handler,
		},
		{
			MethodName: "UserAddChannel",
			Handler:    _Channel_UserAddChannel_Handler,
		},
		{
			MethodName: "UserPatchChannel",
			Handler:    _Channel_UserPatchChannel_Handler,
		},
		{
			MethodName: "ArticleChannel",
			Handler:    _Channel_ArticleChannel_Handler,
		},
		{
			MethodName: "ArticleStatus",
			Handler:    _Channel_ArticleStatus_Handler,
		},
		{
			MethodName: "ArticleRead",
			Handler:    _Channel_ArticleRead_Handler,
		},
		{
			MethodName: "ArticleLike",
			Handler:    _Channel_ArticleLike_Handler,
		},
		{
			MethodName: "ArticleToLike",
			Handler:    _Channel_ArticleToLike_Handler,
		},
		{
			MethodName: "ArticleToDisLike",
			Handler:    _Channel_ArticleToDisLike_Handler,
		},
		{
			MethodName: "ArticleStatusCache",
			Handler:    _Channel_ArticleStatusCache_Handler,
		},
		{
			MethodName: "ArticleToCollection",
			Handler:    _Channel_ArticleToCollection_Handler,
		},
		{
			MethodName: "ArticleToDisCollection",
			Handler:    _Channel_ArticleToDisCollection_Handler,
		},
		{
			MethodName: "ArticleUserCollection",
			Handler:    _Channel_ArticleUserCollection_Handler,
		},
		{
			MethodName: "ArticleToComment",
			Handler:    _Channel_ArticleToComment_Handler,
		},
		{
			MethodName: "ArticleCommentList",
			Handler:    _Channel_ArticleCommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel.proto",
}
