package error

const (
	SUCCSE = 200
	ERROR  = 500

	ERROR_USERNAME_EXSIT = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOTEXSIT  = 1003

	ERROR_TOKEN_EXSIT    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_DISMATCH = 1007
)

var errorCodeToMsg = map[uint16]string{
	SUCCSE: "OK",
	ERROR:  "FAiL",

	ERROR_USERNAME_EXSIT: "Username already exists!",
	ERROR_PASSWORD_WRONG: "Wrong password!",
	ERROR_USER_NOTEXSIT:  "User not exist!",
}

func ErrMsg(errCode uint16) string {
	return errorCodeToMsg[errCode]
}
