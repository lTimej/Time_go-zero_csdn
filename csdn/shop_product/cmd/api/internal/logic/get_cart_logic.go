package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/common/ctxdata"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/productclient"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartLogic) GetCart(req *types.GetCartRequest) (resp *types.GetCartResponse, err error) {
	// todo: add your logic here and delete this line
	user_id := ctxdata.GetUidFromCtx(l.ctx)
	cart_info, err := l.svcCtx.ProductRpc.GetCart(l.ctx, &productclient.GetCartRequest{UserId: user_id})
	if err != nil {
		return nil, err
	}
	resp = new(types.GetCartResponse)
	copier.Copy(resp, cart_info)
	return resp, nil
}
