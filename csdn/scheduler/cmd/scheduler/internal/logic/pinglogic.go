package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/scheduler/scheduler/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/scheduler/scheduler/scheduler"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *scheduler.Request) (*scheduler.Response, error) {
	// todo: add your logic here and delete this line

	return &scheduler.Response{}, nil
}
