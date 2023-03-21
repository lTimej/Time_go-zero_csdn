package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"
	"liujun/Time_go-zero_csdn/csdn/user/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type UserFocusListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFocusListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFocusListLogic {
	return &UserFocusListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFocusListLogic) UserFocusList(in *user.UserFocusListRequest) (*user.UserFocusListResponse, error) {
	// todo: add your logic here and delete this line
	focus_user_ids, err := l.get_focus_user_id(in.UserId)
	fmt.Println(err, "哈哈哈哈哈")
	if err != nil {
		return nil, err
	}
	fans_user_ids, err := l.get_fans_user_id(in.UserId)
	if err != nil {
		return nil, err
	}
	var focus []*user.UserFocus
	for _, target_id := range focus_user_ids {
		user_info, err := l.svcCtx.UserModel.FindOneJoinUserProfileByUserId(l.ctx, target_id.Key)
		if err != nil {
			return nil, err
		}
		var mutual_focus bool
		for _, fans := range fans_user_ids {
			if fans.Key == target_id.Key {
				mutual_focus = true
			}
		}
		focus = append(focus, &user.UserFocus{
			UserId:       user_info.UserId,
			Flag:         "已关注",
			UserName:     user_info.UserName,
			HeadPhoto:    user_info.HeadPhoto,
			Introduction: user_info.Introduce,
			MutualFocus:  mutual_focus,
		})
	}
	// Builder := l.svcCtx.ChannelModel.RowBuilder().Where("is_default = ?", 0)
	return &user.UserFocusListResponse{
		Focus:    focus,
		TotalNum: int64(len(focus)),
		Page:     in.Page,
		PageNum:  in.PageNum,
	}, nil
}

func (l *UserFocusListLogic) get_focus_user_id(user_id string) ([]redis.Pair, error) {
	key := fmt.Sprintf(globalkey.UserFocusByUserId, user_id)
	ok, _ := l.svcCtx.RedisClient.Exists(key)
	if !ok {
		builder := l.svcCtx.UserRelationModel.RowBuilder().Where("relation = ?", model.RELATION().FOLLOW)
		relations, err := l.svcCtx.UserRelationModel.FindFocusByUserId(l.ctx, builder, user_id)
		if err != nil {
			return nil, err
		}
		if relations == nil {
			return nil, nil
		}
		for _, relation := range relations {
			l.svcCtx.RedisClient.Zadd(key, utils.TimeToTimeStamp(relation.CreateTime), relation.TargetUserId)
		}
	}

	focus_user_ids, err := l.svcCtx.RedisClient.ZrangeWithScores(key, 0, -1)
	if err != nil {
		return nil, err
	}
	return focus_user_ids, nil
}

func (l *UserFocusListLogic) get_fans_user_id(user_id string) ([]redis.Pair, error) {
	key := fmt.Sprintf(globalkey.UserFansByUserId, user_id)
	ok, _ := l.svcCtx.RedisClient.Exists(key)
	if !ok {
		builder := l.svcCtx.UserRelationModel.RowBuilder().Where("relation = ?", model.RELATION().FOLLOW)
		relations, err := l.svcCtx.UserRelationModel.FindFansByUserId(l.ctx, builder, user_id)
		if err != nil {
			return nil, err
		}
		if relations == nil {
			return nil, nil
		}
		for _, relation := range relations {
			l.svcCtx.RedisClient.Zadd(key, utils.TimeToTimeStamp(relation.CreateTime), relation.TargetUserId)
		}
	}

	fans_user_ids, err := l.svcCtx.RedisClient.ZrangeWithScores(key, 0, -1)
	if err != nil {
		return nil, err
	}
	return fans_user_ids, nil
}
