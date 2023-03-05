// Code generated by goctl. DO NOT EDIT.
package types

type ChannelList struct {
	Id          int64  `json:"id"`
	ChannelName string `json:"channel_name"`
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
