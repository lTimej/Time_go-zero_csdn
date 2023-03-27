package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/rpc/types/im"

	redisclient "github.com/go-redis/redis/v8"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMessageListLogic {
	return &UserMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserMessageListLogic) UserMessageList(in *im.UserMessageListRequest) (*im.UserMessageListResponse, error) {
	// todo: add your logic here and delete this line
	//从缓存获取
	key := fmt.Sprintf(globalkey.UserContactByUserId, in.UserId)
	target_ids, err := l.svcCtx.RedisIm.ZRevRange(l.ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	//缓存没有从数据库取
	if len(target_ids) == 0 {
		builder := l.svcCtx.UserContact.RowDefaultBuilder().Where("owner_id = ?", in.UserId)
		contacts, err := l.svcCtx.UserContact.FindAllByUserId(l.ctx, "", builder)
		if err != nil {
			return nil, err
		}
		for _, contact := range contacts {
			z := &redisclient.Z{
				Score:  float64(utils.TimeToTimeStamp(contact.CreatedAt)),
				Member: contact.TargetId,
			}
			l.svcCtx.RedisIm.ZAdd(l.ctx, key, z)
		}
	}
	infos := []*im.UserInfo{}
	for _, target_id := range target_ids {
		user, _ := l.svcCtx.UserBasic.FindOne(l.ctx, target_id)
		infos = append(infos, &im.UserInfo{
			UserName:  user.UserName,
			HeadPhoto: user.ProfilePhoto,
			Introduce: user.Introduction,
		})
	}

	// l.svcCtx.UserContact.FindOne()
	return &im.UserMessageListResponse{Userinfos: infos}, nil
}
