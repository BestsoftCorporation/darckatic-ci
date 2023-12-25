package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func init() {
	var err error
	// Replace the connection string with your PostgreSQL database connection details.
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		println("init db ... [failed]")
	}
	println("init db ... [OK]")
}

type User struct {
	ID       int
	Username string
}
