package err

// 错误码的对应中文解释

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	ErrNotExistRoom : "不存在该房间",
	ErrAuthCheckTokenFail : "Token鉴权失败",
	ErrAuthCheckTokenTimeout : "Token已超时",
	ErrAuthToken : "Token生成失败",
	ErrAuth : "Token错误",
}

func GetErrMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}