package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/LeoReeYang/GoBlog/utils"
	"github.com/LeoReeYang/GoBlog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 定义错误
var (
	errTokenHasExpired  = errors.New("Token Expired")
	errTokenNotValidYet = errors.New("Token Invalid Yet")
	errTokenMalformed   = errors.New("Token Malformed")
	errTokenInvalid     = errors.New("Token Invalid")
)

func SetToken(username string) (string, int) {
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	tokenstring, err := CreateToken(*claims)
	if err != nil {
		return "", errormsg.ERROR
	}
	return tokenstring, errormsg.SUCCESS
}

func CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// SignedString need a []byte ,not a string
	return token.SignedString([]byte(utils.JWTKey))
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// ParseWithClaims need a []byte ,not a string
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(utils.JWTKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errTokenHasExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errTokenNotValidYet
			} else {
				return nil, errTokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errTokenInvalid
	}

	return nil, errTokenInvalid
}

func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		tokenHeader := ctx.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			code = errormsg.ERROR_TOKEN_EXIST
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errormsg.ErrMsg(code),
			})
			ctx.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errormsg.ErrMsg(code),
			})
			ctx.Abort()
			return
		}

		claim, err := ParseToken(checkToken[1])
		if err != nil {
			if err == errTokenHasExpired {
				ctx.JSON(http.StatusOK, gin.H{
					"status":  errormsg.ERROR,
					"message": errTokenHasExpired,
					"data":    nil,
				})
				ctx.Abort()
				return
			}
			// 其他错误
			ctx.JSON(http.StatusOK, gin.H{
				"status":  errormsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			ctx.Abort()
			return
		}

		ctx.Set("username", claim)
		ctx.Next()
	}
}
