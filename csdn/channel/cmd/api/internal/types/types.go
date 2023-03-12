// Code generated by goctl. DO NOT EDIT.
package types

type ChannelList struct {
	Id          int64  `json:"id"`
	ChannelName string `json:"channel_name"`
}

type ArticleList struct {
	Title         string `json:"title"`
	UserId        int64  `json:"user_id"`
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
	ArticleId string `json:"aid,optional"`
}

type ArticleStatusResponse struct {
	Isfocus       bool   `json:"isfocus"`
	Iscollection  bool   `json:"iscollection"`
	Islike        bool   `json:"islike"`
	CollectionNum int64  `json:"collection_num"`
	LikeNum       int64  `json:"like_num"`
	ReadNum       int64  `json:"read_num"`
	Aid           string `json:"aid"`
}

type ArticleReadRequest struct {
	UserId    string `json:"uid,optional"`
	ArticleId int64  `json:"aid,optional"`
}

type ArticleReadResponse struct {
	Message string `json:"message"`
	Aid     int64  `json:"aid"`
}
