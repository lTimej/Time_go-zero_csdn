package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/xerr"
	"liujun/Time_go-zero_csdn/csdn/user/model"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CancelFocueUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelFocueUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelFocueUserLogic {
	return &CancelFocueUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelFocueUserLogic) CancelFocueUser(in *user.CancelFocusUserRequest) (*user.CancelFocusUserResponse, error) {
	// todo: add your logic here and delete this line
	relation, err := l.svcCtx.UserRelationModel.FindByUserIdTargetUserId(l.ctx, in.UserId, in.TargetId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.ErrDBError, "err:%v", err)
	}
	user_relation := model.UserRelation{
		RelationId:   relation.RelationId,
		UserId:       in.UserId,
		TargetUserId: in.TargetId,
		Relation:     model.RELATION().DELETE,
	}
	err = l.svcCtx.UserRelationModel.Update(l.ctx, &user_relation)
	if err != nil {
		return nil, err
	}
	// relation, err = l.svcCtx.UserRelationModel.FindByUserIdTargetUserId(l.ctx, in.UserId, in.TargetId)
	// if err != nil {
	// 	return nil, err
	// }
	// if relation.Relation == model.RELATION().DELETE {
	// 	userFocuskey := fmt.Sprintf(globalkey.UserFocusByUserId, in.UserId)
	// 	l.svcCtx.RedisClient.Zrem(userFocuskey, utils.Int64ToString(relation.RelationId))
	// }
	return &user.CancelFocusUserResponse{}, nil
}
