package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "root:123@(127.0.0.1:3306)/test?parseTime=true")

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

	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	_, err = db.Exec(query)

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)

	userID, err := result.LastInsertId()

	fmt.Println(userID)

	fmt.Println("Connect successfully")
}
