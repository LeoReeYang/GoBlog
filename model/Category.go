package model

import "gorm.io/gorm"

type Categroy struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
