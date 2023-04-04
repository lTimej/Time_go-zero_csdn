package logic

import (
	"context"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/rpc/productclient"

	"github.com/jinzhu/copier"

	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/shop_product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProducSputListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProducSputListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProducSputListLogic {
	return &ProducSputListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProducSputListLogic) ProducSputList(req *types.ProductSpuListRequest) (resp *types.ProductSpuListResponse, err error) {
	// todo: add your logic here and delete this line
	spus, err := l.svcCtx.ProductRpc.ProductSpuList(l.ctx, &productclient.ProductSpuListRequest{CategoryId: req.CategoryId})
	if err != nil {
		return nil, err
	}
	resp = new(types.ProductSpuListResponse)
	copier.Copy(resp, spus)
	return
}
