package e

const (
	SuccessErrorCode = 0
)

const (

	// 1 开头为用户相关错误码
	FaultErrorCode  = iota + 10000
	ParamsErrorCode // 参数错误
)

const (
	// 2 开头为用户相关错误码
	UserNameOrPasswordNotMatch = iota + 20000
	UsesAbnormal
)
