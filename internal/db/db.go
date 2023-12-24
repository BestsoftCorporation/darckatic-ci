package db

import (
	"darkatic-ci/internal/repository"
	"darkatic-ci/internal/server"
	"darkatic-ci/internal/source"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

type User struct {
	ID       int
	Username string
}

func InitDB() error {
	var err error
	// Replace the connection string with your PostgreSQL database connection details.
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		return fmt.Errorf("failed to connect database")
	}
	println("init db ... [OK]")
	// Migrate the schema
	DB.AutoMigrate(&server.RemoteServer{})
	DB.AutoMigrate(&source.Source{})
	DB.AutoMigrate(&repository.Repository{})

	return nil
}
