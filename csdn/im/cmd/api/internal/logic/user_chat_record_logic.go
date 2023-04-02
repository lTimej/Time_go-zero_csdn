package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/globalkey"

	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"

	"github.com/go-redis/redis/v8"
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
	fmt.Println(user_id, "===========")
	var key string
	if req.TargetUserId > user_id {
		key = "msg_" + user_id + "_" + req.TargetUserId
	} else {
		key = "msg_" + req.TargetUserId + "_" + user_id
	}
	fmt.Println(key, "---------------")
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
	user_chat_count_key := fmt.Sprintf(globalkey.UserChatCount, user_id)
	l.svcCtx.RedisIm.ZAdd(l.ctx, user_chat_count_key, &redis.Z{float64(0), req.TargetUserId})
	l.svcCtx.RedisIm.Zrem(l.ctx, user_chat_count_key, req.TargetUserId)
	return resp, nil
}
