package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type DefaultChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDefaultChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DefaultChannelLogic {
	return &DefaultChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DefaultChannelLogic) DefaultChannel(in *channel.DefaultChannelRequest) (*channel.DefaultChannelResponse, error) {
	// todo: add your logic here and delete this line
	build := l.svcCtx.ChannelModel.RowDefaultBuilder().Where("is_default = ?", 1)
	cs, err := l.svcCtx.ChannelModel.FindAllDefaultChannel(l.ctx, build, "sequence")
	if err != nil {
		return nil, err
	}
	resp := new(channel.DefaultChannelResponse)
	ccs := []*channel.ChannelList{}
	for _, c := range cs {
		ccs = append(ccs, &channel.ChannelList{
			Id:          c.ChannelId,
			ChannelName: c.ChannelName,
		})
	}
	resp.Channels = ccs
	return resp, nil
}
