package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddChannelLogic {
	return &UserAddChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddChannelLogic) UserAddChannel(in *channel.UserAddChannelRequest) (*channel.UserAddChannelResponse, error) {
	// todo: add your logic here and delete this line
	user_channel := model.NewsUserChannel{
		UserId:    in.UserId,
		ChannelId: in.ChannelId,
	}
	_, err := l.svcCtx.UserChannelModel.Insert(l.ctx, &user_channel)
	if err != nil {
		return nil, err
	}
	build := l.svcCtx.UserChannelModel.RowBuilder()
	ucs, err := l.svcCtx.UserChannelModel.FindAllByUserId(l.ctx, build, in.UserId, "")
	if err != nil {
		return nil, err
	}
	resp := new(channel.UserAddChannelResponse)
	ccs := []*channel.ChannelList{}
	for _, c := range ucs {
		ccs = append(ccs, &channel.ChannelList{
			Id:          c.ChannelId,
			ChannelName: c.ChannelName,
		})
	}
	resp.Channels = ccs
	return resp, nil
}
