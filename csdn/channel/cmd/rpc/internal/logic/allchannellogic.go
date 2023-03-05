package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"
)

type AllChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllChannelLogic {
	return &AllChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllChannelLogic) AllChannel(in *channel.ChannelListRequest) (*channel.ChannelListResponse, error) {
	// todo: add your logic here and delete this line
	Builder := l.svcCtx.ChannelModel.RowBuilder()
	cs, err := l.svcCtx.ChannelModel.FindAll(l.ctx, Builder, "")
	fmt.Println(cs, "+++++++++++")
	if err != nil {
		return nil, err
	}
	resp := new(channel.ChannelListResponse)
	ccs := []*channel.ChannelList{}
	for _, c := range cs {
		ccs = append(ccs, &channel.ChannelList{
			Id:          c.ChannelId,
			ChannelName: c.ChannelName,
		})
	}
	resp.Channels = ccs
	fmt.Println(resp.Channels, "--------------")
	return resp, nil
}
