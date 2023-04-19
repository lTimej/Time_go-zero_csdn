package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/city/model"

	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/city/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CityListLogic {
	return &CityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CityListLogic) CityList(req *types.CityRequest) (resp *types.CityResponse, err error) {
	// todo: add your logic here and delete this line
	key := fmt.Sprintf(globalkey.CityByPid, req.Pid)
	var city []*model.City
	ok, _ := l.svcCtx.RedisClient.Exists(key)
	if ok {
		c, _ := l.svcCtx.RedisClient.Get(key)
		json.Unmarshal([]byte(c), &city)
	} else {
		builder := l.svcCtx.CityModel.Builder().Where("pid = ?", req.Pid)
		city, err = l.svcCtx.CityModel.FindAllByPid(l.ctx, builder)
		if err != nil {
			return nil, err
		}
		data, _ := json.Marshal(city)
		err = l.svcCtx.RedisClient.Set(key, string(data))
		if err != nil {
			return nil, err
		}
	}
	resp = new(types.CityResponse)
	for _, c := range city {
		resp.City = append(resp.City, types.CityInfo{
			Id:   c.Id,
			Name: c.Name,
			Pid:  c.Pid,
		})
	}
	return resp, nil
}
