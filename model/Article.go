package model

import (
	errormsg "github.com/LeoReeYang/GoBlog/utils/errormsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string   `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID  uint     `gorm:"type:uint;not null" json:"cid"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category"` //gorm:"foreignKey:CategoryID"
	Description string   `gorm:"type:varchar(200)" json:"description"`
	Content     string   `gorm:"type:longtext" json:"content"`
	Img         string   `gorm:"type:varchar(100)" json:"img"`
}

func AddArticle(article *Article) int {
	if err := db.Create(article).Error; err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

func DeleteArticle(id int) int {
	var article Article
	if err := db.Delete(&article, id).Error; err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

func GetArticle(id int) (Article, int) {
	var article Article
	var errcode = errormsg.SUCCESS

	db.Model(&article).Preload("Category").Find(&article, id)

	if article.CategoryID == 0 {
		errcode = errormsg.ERROR_ARTICLE_NOT_EXIST
	}

	return article, errcode
}

func GetSameCategoryArticleList(CategoryID int, pageSize int, pageNum int) ([]Article, int64, int) {
	var ArticleList []Article
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

	db.Preload("Category").Select("title,category_id,content,created_at,updated_at").Limit(pageSize).Offset(offset).Where("category_id = ?", CategoryID).Find(&ArticleList)

	db.Model(&ArticleList).Where("category_id = ?", CategoryID).Count(&total)

	if err != nil {
		return ArticleList, 0, errormsg.ERROR
	}
	return ArticleList, total, errormsg.SUCCESS
}

func EditArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})

	maps["title"] = data.Title
	maps["category_id"] = data.CategoryID
	maps["description"] = data.Description
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
