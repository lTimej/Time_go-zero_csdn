package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/user/model"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FocueUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFocueUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FocueUserLogic {
	return &FocueUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FocueUserLogic) FocueUser(in *user.FocusUserRequest) (*user.FocusUserResponse, error) {
	// todo: add your logic here and delete this line
	user_relation := model.UserRelation{
		UserId:       in.UserId,
		TargetUserId: in.TargetId,
		Relation:     model.RELATION().FOLLOW,
	}
	ur, err := l.svcCtx.UserRelationModel.FindByUserIdTargetUserId(l.ctx, in.UserId, in.TargetId)
	if err != nil {
		return nil, err
	}
	if ur == nil {
		_, err = l.svcCtx.UserRelationModel.Insert(l.ctx, &user_relation)
		if err != nil {
			return nil, err
		}
	} else {
		user_relation.RelationId = ur.RelationId
		err = l.svcCtx.UserRelationModel.Update(l.ctx, &user_relation)
		if err != nil {
			return nil, err
		}
	}
	return &user.FocusUserResponse{
		TargetId: in.TargetId,
	}, nil
}
