package jobtype

type DeferSaveArticleStatisticPayload struct {
	ArticleId int64
	UserId    string
	TargetId  string
}