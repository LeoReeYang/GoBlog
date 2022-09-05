package model

import (
	"github.com/LeoReeYang/GoBlog/utils/error"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type varchar(20);not null" json:"password"`
	Role     int    `gorm:"type :int" json:"role"`
}

// check User wheather exsit
func CheckUserExsit(username string) (errorCode int) {
	var user User

	db.Where("name = ?", username).First(&user)

	if user.Username == "" {
		return error.ERROR_USER_NOTEXSIT
	}

	return error.ERROR_USERNAME_EXSIT
}

func AddUser(user *User) (errorCode int) {

	if errcode := CheckUserExsit(user.Username); errcode == error.ERROR_USERNAME_EXSIT {
		return errcode
	}

	db.Create(user)
	return error.SUCCSE
}
