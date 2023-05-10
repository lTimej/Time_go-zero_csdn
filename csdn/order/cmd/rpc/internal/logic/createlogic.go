package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分布式事务
func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// todo: add your logic here and delete this line

	return &order.CreateResponse{}, nil
}
