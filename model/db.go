package model

import (
	"fmt"
	"log"
	"time"

	"github.com/LeoReeYang/GoBlog/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	utils.DBUser, utils.DBPassWords, utils.DBHost, utils.DBPort, utils.DBName)

func InitDB() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Categroy{}, &Article{})

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
}
