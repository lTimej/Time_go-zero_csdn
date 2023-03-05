package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/sliceSet"
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
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	ret, err := l.svcCtx.ChannelRpc.AllChannel(l.ctx, &channelclient.ChannelListRequest{})
	if err != nil {
		return nil, err
	}
	ucs, err := l.svcCtx.ChannelRpc.UserChannel(l.ctx, &channelclient.UserChannelRequest{UserId: user_id})
	data1 := []types.ChannelList{}
	for _, c := range ret.Channels {
		d := types.ChannelList{
			Id:          c.Id,
			ChannelName: c.ChannelName,
		}
		data1 = append(data1, d)
	}
	data2 := []types.ChannelList{}
	for _, c := range ucs.Channels {
		d := types.ChannelList{
			Id:          c.Id,
			ChannelName: c.ChannelName,
		}
		data2 = append(data2, d)
	}
	type d struct {
		Channels interface{} `json:"channels"`
	}
	res := []types.ChannelList{}
	datas := sliceSet.Mines(data1, data2)
	for _, data := range datas {
		res = append(res, types.ChannelList{Id: data["Id"].(int64), ChannelName: data["ChannelName"].(string)})
	}
	return &types.AllChannelResponse{
		Channels: res,
	}, nil
}
