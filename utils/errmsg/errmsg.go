package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... 用户模块错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003

	// 用户携带token不存在
	ERROR_TOKEN_EXIST = 1004
	// 用户token过期
	ERROR_TOKEN_RUNTIME = 1005
	// 用户token不对应
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	// code = 2000... 文章模块错误

	// code = 3000...  分类模块错误
)

var CodeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误！",
	ERROR_USER_NOT_EXIST:   "用户不存在！",
	ERROR_TOKEN_EXIST:      "TOKEN 不存在！",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期！",
	ERROR_TOKEN_WRONG:      "TOKEN不正确！",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN 格式错误！",
}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}