package model

import (
	"github.com/LeoReeYang/GoBlog/utils/error"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int;default:2" json:"role"`
}

// check User wheather exsit
func CheckUserExsit(username string) (errorCode int) {
	var user User

	db.Where("name = ?", username).First(&user)

	if user.Name == "" {
		return error.ERROR_USER_NOTEXSIT
	}

	return error.ERROR_USER_EXSIT
}

func AddUser(user *User) (errorCode int) {

	if errcode := CheckUserExsit(user.Name); errcode == error.ERROR_USER_EXSIT {
		return errcode
	}

	db.Create(user)
	return error.SUCCESS
}

func EditUser(id int, data *User) int {
	var user User
	maps := make(map[string]interface{})

	maps["username"] = data.Name
	maps["role"] = data.Role

	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		return error.ERROR
	}
	return error.SUCCESS
}

func GetUser(id int) (User, int) {
	var user User

	err := db.Limit(1).Where("id = ?", id).Find(&user).Error

	if err != nil {
		return user, error.ERROR
	}
	return user, error.SUCCESS
}

func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	// if userName != "" {
	// 	db.Select("id,username,role,created_at").Where("username LIKE ?", userName+"%").
	// 		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)

	// 	db.Model(&users).Where(
	// 		"username LIKE ?", userName+"%",
	// 	).Count(&total)
	// 	return users, total
	// }

	db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total
}

func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return error.ERROR
	}
	return error.SUCCESS
}
