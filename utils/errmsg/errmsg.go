package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code 1000... 表示用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USE_NOT_EXIST    = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008

	// code 2000... 表示分类模块错误
	ERROR_CATEGORYNAME_USED  = 2001
	ERROR_CATEGORY_NOT_EXIST = 2002

	// code 3000... 表示文章模块错误
	ERROR_ARTICLE_NOT_EXIST = 3001
)

var codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USE_NOT_EXIST:    "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "Token 不存在",
	ERROR_TOKEN_RUNTIME:    "Token 已过期",
	ERROR_TOKEN_WRONG:      "Token 不正确",
	ERROR_TOKEN_TYPE_WRONG: "Token 格式错误",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_CATEGORYNAME_USED: "分类已存在",

	ERROR_ARTICLE_NOT_EXIST: "文章不存在",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
