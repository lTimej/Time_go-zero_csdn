package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/api/internal/types"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPatchChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPatchChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPatchChannelLogic {
	return &UserPatchChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPatchChannelLogic) UserPatchChannel(req *types.UserPatchChannelRequest) (resp *types.UserPatchChannelResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	ucs, err := l.svcCtx.ChannelRpc.UserPatchChannel(l.ctx, &channelclient.UserPatchChannelRequest{
		ChannelId: req.ChannelId,
		UserId:    user_id,
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
	return &types.UserPatchChannelResponse{
		Channels: data,
	}, nil
}
