package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/scheduler/job/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/scheduler/job/job"

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

func (l *PingLogic) Ping(in *job.Request) (*job.Response, error) {
	// todo: add your logic here and delete this line

	return &job.Response{}, nil
}
