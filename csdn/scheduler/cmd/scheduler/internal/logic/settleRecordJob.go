package logic

import (
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/jobtype"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *Timer) settleRecordScheduler() {

	task := asynq.NewTask(jobtype.ScheduleSaveArticleStatistic, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("*/30 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【settleRecordScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【settleRecordScheduler】 registered an  entry: %q \n", entryID)
}
