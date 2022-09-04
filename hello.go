package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "test_mysql:123@(host.docker.internal)/test?parseTime=true")

	if err != nil {
		fmt.Printf("%s when open DB", err)
	}

	defer db.Close()
	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connect successfully")
}
