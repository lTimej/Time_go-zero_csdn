package resp

type ResponseSuccess struct {
	Data interface{} `json:"data"`
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
