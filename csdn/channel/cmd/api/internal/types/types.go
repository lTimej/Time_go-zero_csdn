// Code generated by goctl. DO NOT EDIT.
package types

type AllChannelRequest struct {
}

type ChannelList struct {
	Id          int64  `json:"id"`
	ChannelName string `json:"channel_name"`
}

type AllChannelResponse struct {
	Channels []ChannelList `json:"channels"`
}