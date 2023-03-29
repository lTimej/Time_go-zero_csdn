package logic

import (
	"context"
	"encoding/json"
	"errors"
	"liujun/Time_go-zero_csdn/common/ctxdata"

	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChatRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatRecordLogic {
	return &UserChatRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserChatRecordLogic) UserChatRecord(req *types.UserChatRecordRequest) (resp *types.UserChatRecordResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	key := "msg_" + req.TargetUserId + "_" + user_id
	data, err := l.svcCtx.RedisIm.ZRange(l.ctx, key, req.Page, req.PageNum).Result()
	if err != nil {
		return nil, errors.New("获取消息失败")
	}
	var records types.UserChatRecords
	resp = new(types.UserChatRecordResponse)
	for _, record := range data {
		json.Unmarshal([]byte(record), &records)
		resp.ChatRecords = append(resp.ChatRecords, records)
	}
	return resp, nil
}