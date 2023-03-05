package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddChannelLogic {
	return &UserAddChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddChannelLogic) UserAddChannel(req *types.UserAddChannelRequest) (resp *types.UserAddChannelResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	ucs, err := l.svcCtx.ChannelRpc.UserAddChannel(l.ctx, &channelclient.UserAddChannelRequest{
		UserId:      user_id,
		ChannelId:   req.ChannelId,
		ChannelName: req.ChannelName,
	})
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
	return &types.UserAddChannelResponse{
		Channels: data,
	}, nil
}
