package svc

import (
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/scheduler/internal/config"

	"github.com/hibiken/asynq"
)

type ServiceContext struct {
	Config    config.Config
	Scheduler *asynq.Scheduler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Scheduler: newScheduler(c),
	}
}
