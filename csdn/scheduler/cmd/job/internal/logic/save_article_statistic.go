package logic

import (
	"context"
	"fmt"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/channelclient"
	"liujun/Time_go-zero_csdn/csdn/scheduler/cmd/job/internal/svc"

	"github.com/hibiken/asynq"
)

type SaveArticleStatisticHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSaveArticleStatisticHandler(svcCtx *svc.ServiceContext) *SaveArticleStatisticHandler {
	return &SaveArticleStatisticHandler{
		svcCtx: svcCtx,
	}
}

// every one minute exec : if return err != nil , asynq will retry
func (l *SaveArticleStatisticHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	_, err := l.svcCtx.ChannelRpc.ArticleStatusCache(ctx, &channelclient.ArticleStatusCacheRequest{})
	if err != nil {
		fmt.Println("缓存失败，err:", err)
		return nil
	}
	fmt.Println("缓存成功！")
	// fmt.Printf("shcedule job demo -----> every one minute exec \n")
	return nil
}
