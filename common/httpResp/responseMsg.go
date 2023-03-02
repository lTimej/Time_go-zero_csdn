package httpResp

import (
	"encoding/json"
)

type ResponseSuccess struct {
	Data interface{} `json:"data"`
}
type ResponseLoginSuccess struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginSuccess(data interface{}) (resp *ResponseLoginSuccess) {
	res, _ := json.Marshal(data)
	_ = json.Unmarshal(res, &resp)
	return
}
func Success(data interface{}) *ResponseSuccess {
	return &ResponseSuccess{
		Data: data,
	}
}

type ResponseError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Error(errCode uint32, errMsg string) *ResponseError {
	return &ResponseError{errCode, errMsg}
}
