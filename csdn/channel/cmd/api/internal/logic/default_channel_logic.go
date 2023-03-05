package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DefaultChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDefaultChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DefaultChannelLogic {
	return &DefaultChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DefaultChannelLogic) DefaultChannel(req *types.DefaultChannelRequest) (resp *types.DefaultChannelResponse, err error) {
	// todo: add your logic here and delete this line
	ret, err := l.svcCtx.ChannelRpc.DefaultChannel(l.ctx, &channelclient.DefaultChannelRequest{})
	if err != nil {
		return nil, err
	}
	data := []types.ChannelList{}
	for _, c := range ret.Channels {
		d := types.ChannelList{
			Id:          c.Id,
			ChannelName: c.ChannelName,
		}
		data = append(data, d)
	}
	return &types.DefaultChannelResponse{
		Channels: data,
	}, nil
}
