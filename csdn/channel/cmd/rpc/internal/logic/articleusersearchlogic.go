package logic

import (
	"context"
	"github.com/olivere/elastic/v7"
	"reflect"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleUserSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleUserSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleUserSearchLogic {
	return &ArticleUserSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleUserSearchLogic) ArticleUserSearch(in *channel.ArticleUserSearchRequest) (*channel.ArticleUserSearchResponse, error) {
	// todo: add your logic here and delete this line
	aids, err := l.get_article_id("articles", in.Keyword, in.Page, in.PageNum)
	if err != nil {
		return nil, err
	}
	resp := new(channel.ArticleUserSearchResponse)
	resp.TotalNum = 0
	for _, aid := range aids {
		article, err := l.svcCtx.ArticleModel.FindOneByArticleId(l.ctx, aid)
		if err != nil {
			return nil, err
		}
		resp.ArticlesList = append(resp.ArticlesList, &channel.ArticleList{
			ArtId:         article.ArtId,
			UserId:        article.UserId,
			ChannelId:     article.ChannelId,
			Title:         article.Title,
			CreateTime:    article.CreateTime,
			AllowComment:  article.AllowComment,
			Content:       article.Content,
			UserName:      article.UserName,
			HeadPhoto:     article.HeadPhoto,
			Career:        article.Career,
			ReadNum:       0,
			CodeYear:      article.CodeYear,
			CommentNum:    0,
			LikeNum:       0,
			CollectionNum: 0,
		})
		resp.TotalNum++
	}
	resp.Message = "获取成功"
	return resp, nil
}

func (l *ArticleUserSearchLogic) get_article_id(index, keyword string, page, page_num int64) (aids []int64, err error) {
	offset := (page - 1) * page_num
	termQuery := elastic.NewTermQuery("title", keyword)
	searchResult, err := l.svcCtx.EsClient.Search().
		Index(index).
		Query(termQuery).
		Sort("user_id", false).
		From(int(offset)).Size(int(page_num)).
		Pretty(true).
		Do(l.ctx)
	if err != nil {
		return nil, err
	}
	var ttyp channel.ArticleList
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(channel.ArticleList)
		aids = append(aids, t.ArtId)
	}
	return
}
