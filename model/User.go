package model

import (
	"errors"
	"log"

	errormsg "github.com/LeoReeYang/GoBlog/utils/errormsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

const KeyLen = 10

func CheckUserExist(name string) (errorCode int) {
	var user User

	db.Where("name = ?", name).First(&user)

	if user.ID == 0 {
		return errormsg.ERROR_USER_NOTEXIST
	}

	return errormsg.ERROR_USER_EXIST
}

func AddUser(user *User) int {

	if errcode := CheckUserExist(user.Name); errcode == errormsg.ERROR_USER_EXIST {
		return errcode
	}

	db.Create(user)
	return errormsg.SUCCESS
}

func EditUser(id int, data *User) int {
	var user User
	infos := make(map[string]interface{})

	infos["name"] = data.Name
	infos["role"] = data.Role

	err := db.Model(&user).Where("id = ?", id).Updates(infos).Error

	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

func GetUser(id int) (User, int) {
	var user User
	var errcode = errormsg.SUCCESS

	db.Model(&User{}).First(&user, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		errcode = errormsg.ERROR_USER_NOTEXIST
	}
	return user, errcode
}

func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	var offset int

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}

	db.Select("name,role,created_at,updated_at").Limit(pageSize).Offset(offset).Find(&users)
	db.Model(&users).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total
}

func DeleteUser(id int) int {
	var user User
	if err := db.Delete(&user, id).Error; err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

func EditPassword(id int, user *User) int {
	err = db.Select("password").Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

func (u *User) BeforeSave(db *gorm.DB) (err error) {
	u.Password = ScryptPassword(u.Password)
	u.Role = 1
	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) (err error) {
	u.Password = ScryptPassword(u.Password)
	return nil
}

func ScryptPassword(password string) string {
	// salt := []byte{12, 22, 57, 23, 15, 64, 2, 9}

	// hashPassword, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}
	// return base64.StdEncoding.EncodeToString(hashPassword)
	return string(hashPassword)
}

func UserLogin(name string, password string) int {
	var user User

	db.First(&user, "name = ?", name)

	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return errormsg.ERROR_USER_NOTEXIST
	} else if passwordErr != nil {
		// user.Password = user.Password + " " + password
		return errormsg.ERROR_WRONG_PASSWORD
	} else if user.Role != 1 {
		return errormsg.ERROR_NO_PERMISSION
	}

	return errormsg.SUCCESS
}

func AdminLogin(name string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.First(&user, "name = ?", name)
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errormsg.ERROR_USER_NOTEXIST
	} else if PasswordErr != nil {
		return user, errormsg.ERROR_WRONG_PASSWORD
	}

	return user, errormsg.SUCCESS
}
