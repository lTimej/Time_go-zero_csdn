// Code generated by goctl. DO NOT EDIT.
package types

type ChannelList struct {
	Id          int64  `json:"id"`
	ChannelName string `json:"channel_name"`
}

type ArticleList struct {
	Title         string `json:"title"`
	UserId        string `json:"user_id"`
	CreateTime    string `json:"create_time"`
	ArtId         int64  `json:"art_id"`
	ChannelId     int64  `json:"channel_id"`
	Content       string `json:"content"`
	AllowComment  int32  `json:"allow_comment"`
	UserName      string `json:"user_name"`
	HeadPhoto     string `json:"head_photo"`
	Career        string `json:"career"`
	CodeYear      int32  `json:"code_year"`
	ReadNum       int32  `json:"read_Num"`
	CommentNum    int32  `json:"comment_num"`
	LikeNum       int32  `json:"like_num"`
	CollectionNum int32  `json:"collection_num"`
}

type CommentList struct {
	CommentId       int64  `json:"comment_id"`
	ParentCommentId int64  `json:"parent_comment_id"`
	Ctime           string `json:"ctime"`
	AuthorId        string `json:"author_id"`
	IsTop           int64  `json:"is_top"`
	Content         string `json:"content"`
	CommentIsLike   int64  `json:"comment_is_like"`
}

type ArticleCommentList struct {
	CommentId       int64          `json:"comment_id"`
	ParentCommentId int64          `json:"parent_comment_id"`
	Ctime           string         `json:"ctime"`
	AuthorId        string         `json:"author_id"`
	IsTop           int64          `json:"is_top"`
	Content         string         `json:"content"`
	CommentIsLike   int64          `json:"comment_is_like"`
	CComments       []*CommentList `json:"cComments"`
}

type AllChannelRequest struct {
}

type AllChannelResponse struct {
	Channels []ChannelList `json:"channels"`
}

type DefaultChannelRequest struct {
}

type DefaultChannelResponse struct {
	Channels []ChannelList `json:"default_channel"`
}

type UserChannelRequest struct {
}

type UserChannelResponse struct {
	Channels []ChannelList `json:"channels"`
}

type UserAddChannelRequest struct {
	ChannelId   int64  `json:"channel_id"`
	ChannelName string `json:"channel_name"`
}

type UserAddChannelResponse struct {
	Channels []ChannelList `json:"channels"`
}

type UserPatchChannelRequest struct {
	ChannelId int64 `json:"channel_id"`
}

type UserPatchChannelResponse struct {
	Channels []ChannelList `json:"channels"`
}

type AllArticleRequest struct {
	Page    int32 `json:"page,optional"`
	PageNum int32 `json:"page_num,optional"`
}

type AllArticleResponse struct {
	Articles []ArticleList `json:"articles"`
}

type ArticleStatusRequest struct {
	UserId    string `json:"usid,optional"`
	ArticleId int64  `json:"aid,optional"`
}

type ArticleStatusResponse struct {
	Isfocus       bool  `json:"isfocus"`
	Iscollection  bool  `json:"iscollection"`
	Islike        bool  `json:"islike"`
	CollectionNum int64 `json:"collection_num"`
	LikeNum       int64 `json:"like_num"`
	ReadNum       int64 `json:"read_num"`
	Aid           int64 `json:"aid"`
}

type ArticleReadRequest struct {
	UserId    string `json:"uid,optional"`
	ArticleId int64  `json:"aid,optional"`
}

type ArticleReadResponse struct {
	Message string `json:"message"`
	Aid     int64  `json:"aid"`
}

type ArticleLikeRequest struct {
	ArticleId int64 `json:"aid,optional"`
}

type ArticleLikeList struct {
	HeadPhoto string `json:"head_photo"`
	Aid       int64  `json:"aid"`
}

type ArticleLikeResponse struct {
	UsersInfo []ArticleLikeList `json:"users_info"`
}

type ArticleToLikeRequest struct {
	ArticleId int64 `json:"aid"`
}

type ArticleToLikeResponse struct {
	ArticleId int64 `json:"aid"`
}

type ArticleToDisLikeRequest struct {
	ArticleId int64 `json:"aid"`
}

type ArticleToDisLikeResponse struct {
	Message string `json:"message"`
}

type ArticleToCollectionRequest struct {
	ArticleId int64 `json:"aid"`
}

type ArticleToCollectionResponse struct {
	ArticleId int64 `json:"aid"`
}

type ArticleToDisCollectionRequest struct {
	ArticleId int64 `json:"aid"`
}

type ArticleToDisCollectionResponse struct {
	Message string `json:"message"`
}

type ArticleUserCollectionRequest struct {
	Page    int32 `json:"page,optional"`
	PageNum int32 `json:"page_num,optional"`
}

type ArticleUserCollectionResponse struct {
	Collections []ArticleList `json:"collections"`
	Page        int32         `json:"page"`
	PageNum     int32         `json:"page_num"`
	TotalNum    int32         `json:"total_num"`
}

type ArticleToCommentRequest struct {
	ArticleId int64  `json:"article_id"`
	CommentId int64  `json:"comment_id,optional"`
	Content   string `json:"content"`
}

type ArticleToCommentResponse struct {
	ArticleId       int64 `json:"art_id"`
	CommentParentId int64 `json:"comment_parent_id"`
	CommentId       int64 `json:"comment_id"`
}

type ArticleCommentListRequest struct {
	Ty        string `json:"type,optional"`
	ArticleId int64  `json:"article_id,optional"`
	Offset    int64  `json:"offset,optional"`
	Limit     int64  `json:"limit,optional"`
}

type ArticleCommentListResponse struct {
	Comments []*ArticleCommentList `json:"comments"`
	TotalNum int64                 `json:"total_num"`
	EndId    int64                 `json:"end_id"`
	LastId   int64                 `json:"last_id"`
}

type AllArticleUserRequest struct {
	Page    int32 `json:"page,optional"`
	PageNum int32 `json:"page_num,optional"`
}

type AllArticleUserResponse struct {
	Articles []ArticleList `json:"articles"`
	Page     int32         `json:"page,optional"`
	PageNum  int32         `json:"page_num,optional"`
	TotalNum int64         `json:"total_num,optional"`
}

type ArticleSuggestSearchRequest struct {
	Keyword string `json:"keyword,optional"`
}

type ArticleSuggestSearchResponse struct {
	Searchs []string `json:"searchs"`
}
