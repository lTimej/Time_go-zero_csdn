package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[CODE_EXPIRE_ERROR] = "验证码已过期"
	message[CODE_ERROR] = "验证码错误"
	message[USERNAME_PASSWORD_ERROR] = "用户名或密码错误"
	message[OTHER_ERROR] = "数据库错误"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
