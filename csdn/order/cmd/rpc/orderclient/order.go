// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package orderclient

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/order/cmd/rpc/types/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	OrderCreateRequest  = order.OrderCreateRequest
	OrderCreateResponse = order.OrderCreateResponse
	OrderDescInfo       = order.OrderDescInfo
	OrderDescRequest    = order.OrderDescRequest
	OrderDescResponse   = order.OrderDescResponse
	OrderGetRequest     = order.OrderGetRequest
	OrderGetResponse    = order.OrderGetResponse
	OrderInfo           = order.OrderInfo
	OrderSpec           = order.OrderSpec
	OrderUpdateRequest  = order.OrderUpdateRequest
	OrderUpdateResponse = order.OrderUpdateResponse
	Sku                 = order.Sku

	Order interface {
		OrderCreate(ctx context.Context, in *OrderCreateRequest, opts ...grpc.CallOption) (*OrderCreateResponse, error)
		OrderGet(ctx context.Context, in *OrderGetRequest, opts ...grpc.CallOption) (*OrderGetResponse, error)
		OrderDesc(ctx context.Context, in *OrderDescRequest, opts ...grpc.CallOption) (*OrderDescResponse, error)
		OrderUpdate(ctx context.Context, in *OrderUpdateRequest, opts ...grpc.CallOption) (*OrderUpdateResponse, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) OrderCreate(ctx context.Context, in *OrderCreateRequest, opts ...grpc.CallOption) (*OrderCreateResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.OrderCreate(ctx, in, opts...)
}

func (m *defaultOrder) OrderGet(ctx context.Context, in *OrderGetRequest, opts ...grpc.CallOption) (*OrderGetResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.OrderGet(ctx, in, opts...)
}

func (m *defaultOrder) OrderDesc(ctx context.Context, in *OrderDescRequest, opts ...grpc.CallOption) (*OrderDescResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.OrderDesc(ctx, in, opts...)
}

func (m *defaultOrder) OrderUpdate(ctx context.Context, in *OrderUpdateRequest, opts ...grpc.CallOption) (*OrderUpdateResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.OrderUpdate(ctx, in, opts...)
}