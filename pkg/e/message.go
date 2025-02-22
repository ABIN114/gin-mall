package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	ErrorExistUser:             "用户已存在！",
	ErrorFailEncryption:        "密码加密失败！",
	ErrorExistUserNotFound:     "用户不存在！",
	ErrorNotCompare:            "密码错误！",
	ErrorAuthToken:             "token 验证失败！",
	ErrorAuthCheckTokenTimeout: "token 过期！",
	ErrorUploadFail:            "图片上传失败！",
	ErrorSendEmail:             "邮件发送失败！",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
