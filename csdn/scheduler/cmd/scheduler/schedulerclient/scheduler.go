// Code generated by goctl. DO NOT EDIT.
// Source: scheduler.proto

package schedulerclient

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/scheduler/scheduler/scheduler"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = scheduler.Request
	Response = scheduler.Response

	Scheduler interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultScheduler struct {
		cli zrpc.Client
	}
)

func NewScheduler(cli zrpc.Client) Scheduler {
	return &defaultScheduler{
		cli: cli,
	}
}

func (m *defaultScheduler) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := scheduler.NewSchedulerClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
