package main

import (
	_ "github.com/LeoReeYang/GoBlog/model"
	_ "github.com/LeoReeYang/GoBlog/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//initialize MySQL
	// model.Initdb()
	//initialize Routers
	// routers.InitRouters()
}
