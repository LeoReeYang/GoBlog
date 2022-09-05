package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid         int      `gorm:"type:int;not null" json:"cid"`
	Categroy    Categroy `gorm:"type:varchar(100)" json:"categroy"`
	Description string   `gorm:"type:varchar(200)" json:"description"`
	Content     string   `gorm:"type:longtext" json:"content"`
	Img         string   `gorm:"type:varchar(100)" json:"img"`
}
