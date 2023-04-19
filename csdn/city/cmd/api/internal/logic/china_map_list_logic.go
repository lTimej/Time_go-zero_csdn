package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChinaMapListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChinaMapListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChinaMapListLogic {
	return &ChinaMapListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChinaMapListLogic) ChinaMapList(req *types.ChinaMapRequest) (resp *types.ChinaMapResponse, err error) {
	// todo: add your logic here and delete this line
	// {
	// 	"province": [
	// 		{
	// 			"name": "北京市",
	// 			"id": "10000001",
	// 			"pid": "10000000",
	// 		},
	// 		{
	// 			"name": "天津市",
	// 			"id": "10000002",
	// 			"pid": "10000000",
	// 		},
	// 	],
	// 	"city": {
	// 		"10000001": [
	// 			{
	// 				"name": "北京市",
	// 				"id": "10000035",
	// 				"pid": "10000001",
	// 			},
	// 		],
	// 		"10000002": [
	// 			{
	// 				"name": "天津市",
	// 				"id": "10000036",
	// 				"pid": "10000002",
	// 			},
	// 		],
	// 		"10000003": [
	// 			{
	// 				"name": "石家庄市",
	// 				"id": "10000037",
	// 				"pid": "10000003",
	// 			},
	// 			{
	// 				"name": "唐山市",
	// 				"id": "10000038",
	// 				"pid": "10000003",
	// 			},
	// 			{
	// 				"name": "秦皇岛市",
	// 				"id": "10000039",
	// 				"pid": "10000003",
	// 			},
	// 		]
	// 	},
	// 	"district": {
	// 		"10000035": [
	// 			{
	// 				"name": "东城区",
	// 				"id": "10000404",
	// 				"pid": "10000035",
	// 			},
	// 		],
	// 	}
	// }
	return
}
