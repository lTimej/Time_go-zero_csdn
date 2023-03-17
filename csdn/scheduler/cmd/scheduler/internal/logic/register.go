package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/scheduler/internal/svc"
)

type Timer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronScheduler(ctx context.Context, svcCtx *svc.ServiceContext) *Timer {
	return &Timer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Timer) Register() {
	// todo: add your logic here and delete this line
	l.settleRecordScheduler()
}
