package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"liujun/Time_go-zero_csdn/common/cacheTTL"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleCommentListLogic {
	return &ArticleCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleCommentListLogic) ArticleCommentList(in *channel.ArticleCommentListRequest) (*channel.ArticleCommentListResponse, error) {
	// todo: add your logic here and delete this line
	cids, err := l.get_comment_ids(in)
	if err != nil {
		return nil, err
	}
	l.get_comment_list(cids)
	return &channel.ArticleCommentListResponse{}, nil
}

func (l *ArticleCommentListLogic) get_comment_ids(in *channel.ArticleCommentListRequest) (cids []int64, err error) {
	collection_key := fmt.Sprintf(globalkey.ArticleCommentByAid, in.ArticleId)
	comments := []redis.Pair{}
	if in.Offset != 0 {
		comments, err = l.svcCtx.RedisClient.ZrangebyscoreWithScoresAndLimit(collection_key, 0, in.Offset, 0, int(in.Limit))
	} else {
		comments, err = l.svcCtx.RedisClient.ZrangeWithScores(collection_key, 0, -1)
	}
	if err != nil {
		return nil, err
	}
	total_num := len(comments)
	if total_num > 0 { //有缓存
		for _, comment := range comments {
			cids = append(cids, utils.StringToInt64(comment.Key))
		}
		return
	}
	//没有缓存从数据库取
	builder := l.svcCtx.ArticleCommentModel.RowBuilder()
	comment_list, err := l.svcCtx.ArticleCommentModel.FindAll(l.ctx, builder)
	if err != nil {
		return nil, err
	}
	comment_caches := redis.Pair{}
	var count int64
	for _, c := range comment_list {
		if c.IsTop != 0 {
			c.CreateTime = c.CreateTime.Add(cacheTTL.MAXTTL)
		}
		score := utils.TimeToTimeStamp(c.CreateTime)
		comment_caches.Key = utils.Int64ToString(c.CommentId)
		comment_caches.Score = score
		if ((in.Offset != 0 && in.Offset < score) || in.Offset == 0) && count < in.Limit {
			cids = append(cids, c.CommentId)
			count++
		}
	}
	l.svcCtx.RedisClient.Zadds(collection_key, comment_caches)
	l.svcCtx.RedisClient.Expire(collection_key, cacheTTL.ArticleCommentByAid)
	return
}

func (l *ArticleCommentListLogic) get_comment_list(cids []int64) (comments []*model.NewsComment) {
	for _, cid := range cids {
		newsCommentCommentIdKey := fmt.Sprintf(globalkey.ArticleCommentByCid, cid)
		data, err := l.svcCtx.RedisClient.Get(newsCommentCommentIdKey)
		if err != nil {
			return
		}
		var comment model.NewsComment
		json.Unmarshal([]byte(data), &comment)
		comments = append(comments, &comment)
	}
	return comments
}
