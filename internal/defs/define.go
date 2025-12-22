package defs

type AppError struct {
	Code    int64
	Message string
}

func NewErrorMsg(code int64, msg string) *AppError {
	return &AppError{code, msg}

}

var (
	Error_ServerError = NewErrorMsg(10000, "Server Error")
	Error_ArgsError   = NewErrorMsg(10001, "参数错误")

	Error_UserNameOrPswError = NewErrorMsg(20002, "用户名或密码错误")
	Error_UserNotFound       = NewErrorMsg(20003, "用户不存在")
	Error_UserExist          = NewErrorMsg(20004, "用户已存在")
)
