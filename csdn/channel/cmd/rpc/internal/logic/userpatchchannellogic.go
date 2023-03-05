package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPatchChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPatchChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPatchChannelLogic {
	return &UserPatchChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPatchChannelLogic) UserPatchChannel(in *channel.UserPatchChannelRequest) (*channel.UserPatchChannelResponse, error) {
	// todo: add your logic here and delete this line
	user_channel := model.NewsUserChannel{
		UserId:    in.UserId,
		ChannelId: in.ChannelId,
		IsDeleted: 1,
	}
	err := l.svcCtx.UserChannelModel.Update(l.ctx, &user_channel)
	fmt.Println("和混合双打罚款罚款和", err)
	if err != nil {
		return nil, err
	}
	build := l.svcCtx.UserChannelModel.RowBuilder()
	ucs, err := l.svcCtx.UserChannelModel.FindAllByUserId(l.ctx, build, in.UserId, "")
	if err != nil {
		return nil, err
	}
	resp := new(channel.UserPatchChannelResponse)
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
