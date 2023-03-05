package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChannelLogic {
	return &UserChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserChannelLogic) UserChannel(req *types.UserChannelRequest) (resp *types.UserChannelResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	fmt.Println(user_id, "==============user_id===============")
	ucs, err := l.svcCtx.ChannelRpc.UserChannel(l.ctx, &channelclient.UserChannelRequest{UserId: user_id})
	fmt.Println(err, "******err*********")
	if err != nil {
		return nil, err
	}
	data := []types.ChannelList{}
	for _, c := range ucs.Channels {
		d := types.ChannelList{
			Id:          c.Id,
			ChannelName: c.ChannelName,
		}
		data = append(data, d)
	}
	return &types.UserChannelResponse{
		Channels: data,
	}, nil
}
