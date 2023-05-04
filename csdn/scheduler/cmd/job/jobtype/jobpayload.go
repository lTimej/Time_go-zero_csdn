package jobtype

type DeferSaveArticleStatisticPayload struct {
	ArticleId int64
	UserId    string
	TargetId  string
}

type DeferCloseProductOrderPayload struct {
	Sn string
}
