package controller

type ResCode int

const (
	SuccessCode ResCode = 10000 + iota
	ErrorCodeInvalidParams
	ErrorCodeUserAlreadyExist
	ErrorCodeUserNotExist
	ErrorCodeInvalidPassword
	ErrorCodeServerBusy
	ErrorCodInvalidAuth
	ErrorCodeNeedLogin
)

var codeMessageMap = map[ResCode]string{
	SuccessCode:               "请求成功！",
	ErrorCodeInvalidParams:    "请求参数非法！",
	ErrorCodeUserAlreadyExist: "用户已存在！",
	ErrorCodeUserNotExist:     "用户不存在！",
	ErrorCodeInvalidPassword:  "用户名或密码不对！",
	ErrorCodeServerBusy:       "服务繁忙！",
	ErrorCodInvalidAuth:       "无效的Token！",
	ErrorCodeNeedLogin:        "请先登录！",
}

func (code ResCode) getMessage() string {
	msg, ok := codeMessageMap[code]
	if !ok {
		msg = codeMessageMap[ErrorCodeServerBusy]
	}
	return msg
}
