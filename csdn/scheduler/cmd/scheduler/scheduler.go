package main

import (
	"context"
	"flag"
	"os"

	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/scheduler/internal/config"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/scheduler/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/scheduler/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/scheduler.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	Timer := logic.NewCronScheduler(ctx, svcCtx)
	Timer.Register()
	if err := svcCtx.Scheduler.Run(); err != nil {
		logx.Errorf("!!!MqueueSchedulerErr!!!  run err:%+v", err)
		os.Exit(1)
	}
}
