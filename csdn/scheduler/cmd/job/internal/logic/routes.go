package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/jobtype"

	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//scheduler job
	mux.Handle(jobtype.ScheduleSaveArticleStatistic, NewSaveArticleStatisticHandler(l.svcCtx))
	//queue job , asynq support queue job
	// wait you fill..

	return mux
}
