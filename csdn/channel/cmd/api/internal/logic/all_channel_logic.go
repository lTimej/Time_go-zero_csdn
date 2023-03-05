package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllChannelLogic {
	return &AllChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllChannelLogic) AllChannel(req *types.AllChannelRequest) (resp *types.AllChannelResponse, err error) {
	// todo: add your logic here and delete this line
	ret, err := l.svcCtx.ChannelRpc.AllChannel(l.ctx, &channelclient.ChannelListRequest{})
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
	return &types.AllChannelResponse{
		Channels: data,
	}, nil
}
