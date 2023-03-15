package main

import (
	"database/sql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/godb?parseTime=true")
	if err != nil {
		return
	}
	query := `
		CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
		);
		`
	_, err = db.Exec(query)
	if err != nil {
		return
	}
	username := "april"
	password := "123456"
	datetime := time.Now()
	db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?,?,?)`, username, password, datetime)

}
