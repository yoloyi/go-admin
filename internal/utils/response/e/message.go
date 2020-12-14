package e

var codeMap = map[int]string{
	SuccessErrorCode:           "成功",
	FaultErrorCode:             "系统出现异常，请稍后再试",
	UserNameOrPasswordNotMatch: "账号或密码不正确",
	UsesAbnormal:               "账户异常，请联系管理员",
}

func CodeToMessage(code int) string {
	if message, ok := codeMap[code]; ok {
		return message
	}

	return CodeToMessage(FaultErrorCode)
}
