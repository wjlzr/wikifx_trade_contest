package errorx

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	// 用户10001-20000
	ERROR_EXIST_TAG                   = 10001
	ERROR_NOT_EXIST_PARAME_X_USER_ID  = 10002
	ERROR_REDIS_CONNECT_EXEPTION      = 10003
	ERROR_NOT_LOGIN                   = 10004
	ERROR_NOT_ACCESS_AUTHORITY        = 10005
	ERROR_USERNAME_PASSWORD_NOT_EMPTY = 10006
	ERROR_FIND_USER_EXEPTION          = 10007
	ERROR_USERNAME_PASSWORD_ERR       = 10008
	ERROR_USER_CREATE_FAIL            = 10009

	// 通用100001-200000
	ERROR_CURRENCY_UNMARSHAL = 100001

	// 标准化http
	StatusUnauthorized = 401
)

var statusText = map[int]string{
	ERROR_EXIST_TAG:                   "已存在该标签名称",
	ERROR_NOT_EXIST_PARAME_X_USER_ID:  "缺少x-user-id参数",
	ERROR_REDIS_CONNECT_EXEPTION:      "用户：%s, 连接redis异常",
	ERROR_NOT_LOGIN:                   "未登录",
	ERROR_NOT_ACCESS_AUTHORITY:        "沒有路径访问权限，请联系管理员",
	ERROR_USERNAME_PASSWORD_NOT_EMPTY: "用户名或密码不能为空",
	ERROR_FIND_USER_EXEPTION:          "查询用户异常",
	ERROR_USERNAME_PASSWORD_ERR:       "账号或密码错误",
	ERROR_USER_CREATE_FAIL:            "添加用户失败",

	ERROR_CURRENCY_UNMARSHAL: "Unmarshal 错误",

	StatusUnauthorized: "Unauthorized",
}

// GetMsg returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func GetMsg(code int) string {
	return statusText[code]
}
