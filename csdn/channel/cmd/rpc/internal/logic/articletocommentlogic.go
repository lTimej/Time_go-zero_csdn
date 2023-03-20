package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"liujun/Time_go-zero_csdn/common/cacheTTL"
	"liujun/Time_go-zero_csdn/common/globalkey"
	"liujun/Time_go-zero_csdn/common/utils"
	"liujun/Time_go-zero_csdn/csdn/channel/model"

	"github.com/pkg/errors"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleToCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleToCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleToCommentLogic {
	return &ArticleToCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleToCommentLogic) ArticleToComment(in *channel.ArticleToCommnetRequest) (*channel.ArticleToCommentResponse, error) {
	// todo: add your logic here and delete this line
	comment_id := in.CommentId
	Aid := in.ArticleId
	content := in.Content
	UserId := in.UserId
	c, err := l.svcCtx.ArticleCommentModel.FindOneByArticleIdUserIdParentId(l.ctx, Aid, UserId, comment_id)
	if err != nil {
		return nil, err
	}
	if c != nil {
		return nil, errors.New("用户已评论")
	}
	comment := model.NewsComment{
		UserId:     UserId,
		ArticleId:  Aid,
		Content:    content,
		ParentId:   comment_id,
		LikeCount:  0,
		ReplyCount: 0,
		IsTop:      0,
		Status:     1,
	}
	l.svcCtx.ArticleCommentModel.Insert(l.ctx, &comment)
	comment_obj, err := l.svcCtx.ArticleCommentModel.FindOneByArticleIdUserIdParentId(l.ctx, Aid, UserId, comment_id)
	if err != nil {
		return nil, err
	}
	comment_cid_key := fmt.Sprintf(globalkey.ArticleComment, comment_obj.CommentId)
	data, _ := json.Marshal(comment_obj)
	l.svcCtx.RedisClient.Setex(comment_cid_key, string(data), cacheTTL.ArticleCommentByCid)
	var comment_aid_key string
	if comment_id == 0 {
		comment_aid_key = fmt.Sprintf(globalkey.ArticleCommentByAid, Aid)
		_, err = l.svcCtx.RedisClient.Zadd(comment_aid_key, utils.TimeToTimeStamp(comment_obj.CreateTime), utils.Int64ToString(comment_obj.CommentId))
	} else {
		comment_aid_key = fmt.Sprintf(globalkey.ArticleCommentByCid, comment_id) //comment_id是父id  comment_obj.CommentId是子id
		_, err = l.svcCtx.RedisClient.Zadd(comment_aid_key, utils.TimeToTimeStamp(comment_obj.CreateTime), utils.Int64ToString(comment_obj.CommentId))
	}
	l.svcCtx.RedisClient.Expire(comment_aid_key, cacheTTL.ArticleCommentByAid)
	if err != nil {
		fmt.Println("评论缓存失败 err:", err)
	}
	resp := new(channel.ArticleToCommentResponse)
	resp.ArticleId = Aid
	resp.CommentId = comment_obj.CommentId
	resp.CommentParentId = comment_id
	return resp, nil
}
