package model

import (
	"encoding/base64"
	"log"

	"github.com/LeoReeYang/GoBlog/utils/errors"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

const KeyLen = 10

// check User wheather exsit
func CheckUserExsit(username string) (errorCode int) {
	var user User

	db.Where("name = ?", username).First(&user)

	if user.Name == "" {
		return errors.ERROR_USER_NOTEXSIT
	}

	return errors.ERROR_USER_EXSIT
}

func AddUser(user *User) int {

	if errcode := CheckUserExsit(user.Name); errcode == errors.ERROR_USER_EXSIT {
		return errcode
	}

	db.Create(user)
	return errors.SUCCESS
}

func EditUser(id int, data *User) int {
	var user User
	maps := make(map[string]interface{})

	maps["name"] = data.Name
	maps["role"] = data.Role

	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

func GetUser(id int) (User, int) {
	var user User

	err := db.Limit(1).Where("id = ?", id).Find(&user).Error

	if err != nil {
		return user, errors.ERROR
	}
	return user, errors.SUCCESS
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
	var offset int

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}

	db.Select("name,password,role,created_at,updated_at").Limit(pageSize).Offset(offset).Find(&users)
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
		return errors.ERROR
	}
	return errors.SUCCESS
}

func (u *User) BeforeSave(db *gorm.DB) (err error) {
	u.Password = ScryptPassword(u.Password)

	return
}

func ScryptPassword(password string) string {
	salt := []byte{12, 22, 57, 23, 15, 64, 2, 9}

	HashPassword, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)

	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(HashPassword)
}
