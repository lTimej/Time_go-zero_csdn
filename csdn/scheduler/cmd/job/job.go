package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/internal/config"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/internal/logic"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/job.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	conf.MustLoad(*configFile, &c, conf.UseEnv())

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	// conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, svcCtx)
	mux := cronJob.Register()
	if err := svcCtx.AsynqServer.Run(mux); err != nil {
		fmt.Println(err)
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
