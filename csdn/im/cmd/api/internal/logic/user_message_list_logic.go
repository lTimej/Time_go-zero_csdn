package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/im/cmd/api/internal/types"

	redisclient "github.com/go-redis/redis/v8"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMessageListLogic {
	return &UserMessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMessageListLogic) UserMessageList(req *types.UserMessageListRequest) (resp *types.UserMessageListResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	key := fmt.Sprintf(globalkey.UserContactByUserId, user_id)
	var target_ids []string
	target_ids, err = l.svcCtx.RedisIm.ZRevRange(l.ctx, key, 0, 1).Result()
	if err != nil {
		fmt.Println(err, "111111111111")
		return nil, err
	}
	fmt.Println(999999, "===", len(target_ids))
	//缓存没有从数据库取
	if len(target_ids) == 0 {
		builder := l.svcCtx.UserContact.RowDefaultBuilder().Where("owner_id = ?", user_id)
		contacts, err := l.svcCtx.UserContact.FindAllByUserId(l.ctx, "", builder)
		if err != nil {
			fmt.Println(err, "啥？？？？", user_id)
			return nil, err
		}
		for _, contact := range contacts {
			target_ids = append(target_ids, contact.TargetId)
			z := &redisclient.Z{
				Score:  float64(utils.TimeToTimeStamp(contact.CreatedAt)),
				Member: contact.TargetId,
			}
			l.svcCtx.RedisIm.ZAdd(l.ctx, key, z)
		}
	}
	infos := []types.UserInfo{}
	for _, target_id := range target_ids {
		user, err := l.svcCtx.UserBasic.FindOne(l.ctx, target_id)
		if err != nil {
			fmt.Println(err, "2222222222222222")
		}
		infos = append(infos, types.UserInfo{
			UserName:  user.UserName,
			HeadPhoto: user.ProfilePhoto,
			Introduce: user.Introduction,
			UserId:    user.UserId,
		})
	}
	return &types.UserMessageListResponse{UserInfos: infos}, nil
}
