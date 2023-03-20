package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"liujun/Time_go-zero_csdn/common/cacheTTL"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
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
	cids, err := l.get_comment_ids(in.Limit, in.Offset, in.ArticleId, 0)
	if err != nil {
		return nil, err
	}
	comment_list := l.get_comment_list(cids)
	total_num := len(comment_list)
	if total_num == 0 {
		return &channel.ArticleCommentListResponse{}, nil
	}
	end_id := utils.TimeToTimeStamp(comment_list[0].CreateTime)
	last_id := utils.TimeToTimeStamp(comment_list[total_num-1].CreateTime)
	comments := make([]*channel.ArticleCommentList, total_num)
	fmt.Printf("len:%d,cap:%d\n", len(comments), cap(comments))
	for index, comment := range comment_list {
		comments[index] = new(channel.ArticleCommentList)
		comments[index].CommentId = comment.CommentId
		comments[index].ParentCommentId = comment.ParentId
		comments[index].Ctime = utils.TimeToString(comment.CreateTime)
		comments[index].AuthorId = comment.UserId
		comments[index].IsTop = comment.IsTop
		comments[index].Content = comment.Content
		comments[index].CommentIsLike = 0
		cCids, err := l.get_comment_ids(in.Limit, in.Offset, 0, comment.CommentId)
		if err != nil {
			return nil, err
		}
		cComment_list := l.get_comment_list(cCids)
		cComments := make([]*channel.CommentList, len(cComment_list))
		fmt.Printf("len:%d,cap:%d\n", len(cComments), cap(cComments))
		comments[index].CComments = cComments
		for indey, cComment := range cComment_list {
			comments[index].CComments[indey] = new(channel.CommentList)
			comments[index].CComments[indey].CommentId = cComment.CommentId
			comments[index].CComments[indey].ParentCommentId = cComment.ParentId
			comments[index].CComments[indey].Ctime = utils.TimeToString(cComment.CreateTime)
			comments[index].CComments[indey].AuthorId = cComment.UserId
			comments[index].CComments[indey].IsTop = cComment.IsTop
			comments[index].CComments[indey].Content = cComment.Content
			comments[index].CComments[indey].CommentIsLike = 0
		}
	}
	return &channel.ArticleCommentListResponse{
		TotalNum: int64(total_num),
		EndId:    end_id,
		LastId:   last_id,
		Comments: comments,
	}, nil
}

func (l *ArticleCommentListLogic) get_comment_ids(limit, offset, aid int64, cid int64) (cids []int64, err error) {
	var collection_key string
	if cid == 0 {
		collection_key = fmt.Sprintf(globalkey.ArticleCommentByAid, aid)
	} else {
		collection_key = fmt.Sprintf(globalkey.ArticleCommentByCid, cid)
	}
	var comments []redis.Pair
	if offset != 0 {
		comments, err = l.svcCtx.RedisClient.ZrangebyscoreWithScoresAndLimit(collection_key, 0, offset, 0, int(limit))
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
	var builder squirrel.SelectBuilder
	if cid == 0 {
		builder = l.svcCtx.ArticleCommentModel.RowBuilder().Where("article_id = ? AND parent_id = ?", aid, cid)
	} else {
		builder = l.svcCtx.ArticleCommentModel.RowBuilder().Where("parent_id = ?", cid)
	}

	comment_list, err := l.svcCtx.ArticleCommentModel.FindAll(l.ctx, builder)
	if err != nil {
		return nil, err
	}
	if comment_list == nil {
		return
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
		if ((offset != 0 && offset < score) || offset == 0) && count < limit {
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
		newsCommentCommentIdKey := fmt.Sprintf(globalkey.ArticleComment, cid)
		data, err := l.svcCtx.RedisClient.Get(newsCommentCommentIdKey)
		if err != nil {
			return
		}
		var comment model.NewsComment
		json.Unmarshal([]byte(data), &comment)
		comments = append(comments, &comment)
	}
	return
}
