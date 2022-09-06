package error

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USER_EXSIT     = 1001
	ERROR_WRONG_PASSWORD = 1002
	ERROR_USER_NOTEXSIT  = 1003

	ERROR_TOKEN_EXSIT    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_DISMATCH = 1007
)

var errorCodeToMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_USER_EXSIT:     "User already exists!",
	ERROR_WRONG_PASSWORD: "Wrong password!",
	ERROR_USER_NOTEXSIT:  "User not exist!",
}

func ErrMsg(errCode int) string {
	return errorCodeToMsg[errCode]
}
