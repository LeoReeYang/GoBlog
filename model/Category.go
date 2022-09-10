package model

import (
	errormsg "github.com/LeoReeYang/GoBlog/utils/errormsg"
)

// gorm:"primarykey"
type Category struct {
	ID   uint   `gorm:"primaryKey;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// check User wheather exsit
func CheckCategoryExist(name string) (errorCode int) {
	var category Category

	db.Where("name = ?", name).First(&category)

	if category.Name == "" {
		return errormsg.ERROR_CATEGORY_NOT_EXIST
	}

	return errormsg.ERROR_CATEGORY_EXIST
}

func AddCategory(category *Category) int {

	if errcode := CheckCategoryExist(category.Name); errcode == errormsg.ERROR_CATEGORY_EXIST {
		return errcode
	}

	db.Create(category)
	return errormsg.SUCCESS
}

func DeleteCategory(id int) int {
	var category Category
	if err := db.Delete(&category, id).Error; err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

func GetCategory(id int) (Category, int) {
	var category Category
	var errcode = errormsg.SUCCESS

	db.Model(&Category{}).First(&category, id)

	return category, errcode
}

func EditCategory(category *Category) int {

	err := db.Model(&category).Update("name", category.Name).Error

	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

func GetCategoryList(pageSize int, pageNum int) ([]Category, int64) {
	var categoryList []Category
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

	db.Select("id,name,created_at,updated_at").Limit(pageSize).Offset(offset).Find(&categoryList)
	db.Model(&categoryList).Count(&total)

	if err != nil {
		return categoryList, 0
	}
	return categoryList, total
}
