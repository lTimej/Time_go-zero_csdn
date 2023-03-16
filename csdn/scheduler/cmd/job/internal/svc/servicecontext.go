package svc

import "liujun/Time_go-zero_csdn/csdn/scheduler/job/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
