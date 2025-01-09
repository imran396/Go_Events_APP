package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:i@/goDev")
	if err != nil {
		panic(fmt.Sprintf("Database connection failed: %s", err.Error()))
	}
	return db
}


func CreatTable(db *sql.DB  ){
	tableQuery :=`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY, -- Unique identifier
		username VARCHAR(50) NOT NULL,     -- Username, maximum 50 characters
		email VARCHAR(100) UNIQUE NOT NULL, -- Email, must be unique
		password_hash VARCHAR(255) NOT NULL, -- Password hash for authentication
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation time
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Record update time
	);`
	
	_, err := db.Exec(tableQuery)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
