package globalkey

const (
	ArticleStatus        = "article:status:%s"        //文章状态
	ArticleIds           = "cache:article:ids"        //文章id
	ArticleLikeNum       = "article:like:count"       //文章点赞数
	ArticleReadNum       = "article:read:count"       //文章阅读量
	ArticleCollectionNum = "article:collection:count" //文章收藏量
)

const (
	ArticleCommentByAid = "article:comment:article:%d"
	ArticleCommentByCid = "article:comment:comment:%d"
	ArticleComment      = "article:comment:%d"
)

const (
	UserFocusByUserId = "user:focus:%s"
	UserFansByUserId  = "user:fans:%s"
	UserArticleSearch = "user:article:search:%s"
)

const (
	UserContactByUserId = "user:contact:%s"
	UserChatCount       = "user:chat:count:%s"
)

const (
	ProductCategory = "product:category"
)
