package e

const (
	SuccessErrorCode = 0
)

const (
	// 1 开头为系统相关错误码
	FaultErrorCode  = 10000
	ParamsErrorCode = 10001 // 请求参数错误

	NotLoginCode     = 10001 // 用户未登录
	TokenExpiredCode = 10002 // 授权过期
)

const (
	// 2 开头为用户相关错误码
	UserNameOrPasswordNotMatch = 20000
	UsesAbnormal               = 20001
	UserChangePasswordNotMatch       = 20002
)
