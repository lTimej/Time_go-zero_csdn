package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChannelLogic {
	return &UserChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserChannelLogic) UserChannel(in *channel.UserChannelRequest) (*channel.UserChannelResponse, error) {
	// todo: add your logic here and delete this line
	build := l.svcCtx.UserChannelModel.RowBuilder()
	res, err := l.svcCtx.UserChannelModel.FindAllByUserId(l.ctx, build, in.UserId, "sequence")
	if err != nil {
		return nil, err
	}
	resp := new(channel.UserChannelResponse)
	ccs := []*channel.ChannelList{}
	for _, c := range res {
		ccs = append(ccs, &channel.ChannelList{
			Id:          c.ChannelId,
			ChannelName: c.ChannelName,
		})
	}
	resp.Channels = ccs
	return resp, nil
}
