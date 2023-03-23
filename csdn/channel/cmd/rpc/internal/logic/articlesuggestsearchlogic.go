package logic

import (
	"context"

	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/internal/svc"
	"liujun/Time_go-zero_csdn/csdn/channel/cmd/rpc/types/channel"

	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleSuggestSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleSuggestSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleSuggestSearchLogic {
	return &ArticleSuggestSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleSuggestSearchLogic) ArticleSuggestSearch(in *channel.ArticleSuggestSearchRequest) (*channel.ArticleSuggestSearchResponse, error) {
	// todo: add your logic here and delete this line
	searchs, err := l.GetSuggest("completions", in.Keyword)
	if err != nil {
		return nil, err
	}
	return &channel.ArticleSuggestSearchResponse{Searchs: searchs}, nil
}

func (l *ArticleSuggestSearchLogic) GetSuggest(index, keyword string) ([]string, error) {
	ctx := context.Background()
	my_suggest := "my-suggest"
	suggester := elastic.NewCompletionSuggester(my_suggest).Fuzziness(0).
		Text(keyword).Field("title").SkipDuplicates(true)
	searchSource := elastic.NewSearchSource().
		Suggester(suggester).
		FetchSource(false).
		TrackScores(true)
	searchResult, err := l.svcCtx.EsClient.Search().
		Index(index).
		SearchSource(searchSource).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	// fmt.Println(searchResult.Suggest[my_suggest][0].Text)
	// fmt.Println(searchResult.Suggest[my_suggest][0].Offset)
	// fmt.Println(searchResult.Suggest[my_suggest][0].Length)
	// fmt.Println(searchResult.Suggest[my_suggest][0].Options)
	options := searchResult.Suggest[my_suggest][0].Options
	searchs := []string{}
	for _, option := range options {
		searchs = append(searchs, option.Text)
		//fmt.Println(option.Source)
	}
	return searchs, nil
}
