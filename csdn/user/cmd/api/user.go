package main

import (
	"flag"
	"fmt"

	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/config"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/handler"
	"liujun/Time_go-zero_csdn/csdn/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("http://172.20.16.20:8080"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
