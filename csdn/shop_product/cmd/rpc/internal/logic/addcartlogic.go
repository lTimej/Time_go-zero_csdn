package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCartLogic) AddCart(in *product.AddCartRequest) (*product.AddCartResponse, error) {
	// todo: add your logic here and delete this line
	key := ""
	if in.UserId == "0" {
		key = globalkey.AnonymityUserCartList
	} else {
		key = fmt.Sprintf(globalkey.UserCartList, in.UserId)
	}

	l.svcCtx.RedisClient.Hincrby(key, utils.Int64ToString(in.SkuId), 1)
	return &product.AddCartResponse{}, nil
}