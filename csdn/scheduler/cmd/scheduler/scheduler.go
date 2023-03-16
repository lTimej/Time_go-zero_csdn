package main

import (
	"flag"
	"fmt"

	"liujun/Time_go-zero_csdn/csdn/scheduler/scheduler/internal/config"
	"liujun/Time_go-zero_csdn/csdn/scheduler/scheduler/internal/server"
	"liujun/Time_go-zero_csdn/csdn/scheduler/scheduler/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/scheduler/scheduler/scheduler"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/scheduler.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		scheduler.RegisterSchedulerServer(grpcServer, server.NewSchedulerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
