package main

import (
	"github.com/LeoReeYang/GoBlog/model"
	"github.com/LeoReeYang/GoBlog/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// initialize MySQL
	model.InitDB()
	// initialize Routers
	routers.InitRouter()
}
