// Code generated by goctl. DO NOT EDIT.
// Source: job.proto

package jobclient

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/scheduler/job/job"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = job.Request
	Response = job.Response

	Job interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultJob struct {
		cli zrpc.Client
	}
)

func NewJob(cli zrpc.Client) Job {
	return &defaultJob{
		cli: cli,
	}
}

func (m *defaultJob) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}