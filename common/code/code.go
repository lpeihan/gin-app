package code

type Code = int

const (
	OK           Code = 200
	Unauthorized Code = 401
	CommonError  Code = 0
)

var ErrorMessage = map[Code]string{
	OK:           "成功",
	Unauthorized: "权限不足",
	CommonError:  "参数错误",
}

func GetErrorMessage(code Code) string {
	message, ok := ErrorMessage[code]

	if !ok {
		return "未知错误"
	}

	return message
}
