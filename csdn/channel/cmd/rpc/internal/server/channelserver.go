// Code generated by goctl. DO NOT EDIT.
// Source: channel.proto

package server

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"
)

type ChannelServer struct {
	svcCtx *svc.ServiceContext
	channel.UnimplementedChannelServer
}

func NewChannelServer(svcCtx *svc.ServiceContext) *ChannelServer {
	return &ChannelServer{
		svcCtx: svcCtx,
	}
}

func (s *ChannelServer) AllChannel(ctx context.Context, in *channel.ChannelListRequest) (*channel.ChannelListResponse, error) {
	l := logic.NewAllChannelLogic(ctx, s.svcCtx)
	return l.AllChannel(in)
}