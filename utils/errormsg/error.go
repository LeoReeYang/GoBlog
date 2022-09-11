package errormsg

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USER_EXIST     = 1001
	ERROR_WRONG_PASSWORD = 1002
	ERROR_USER_NOTEXIST  = 1003
	ERROR_NO_PERMISSION  = 1004

	ERROR_CATEGORY_EXIST     = 2001
	ERROR_CATEGORY_NOT_EXIST = 2002

	ERROR_ARTICLE_NOT_EXIST = 3001

	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_DISMATCH = 1007
)

var errorCodeToMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_USER_EXIST:     "User already exists!",
	ERROR_WRONG_PASSWORD: "Wrong password!",
	ERROR_USER_NOTEXIST:  "User not exist!",
}

func ErrMsg(errCode int) string {
	return errorCodeToMsg[errCode]
}
