package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID       int
	Username string
}

// CreateInitialAdminUser generates an initial admin user if it does not exist.
func CreateInitialAdminUser(db *sql.DB) error {
	// Check if the admin user already exists
	var adminUserExists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = 'admin')").Scan(&adminUserExists)
	if err != nil {
		return fmt.Errorf("failed to check if admin user exists: %v", err)
	}

	if adminUserExists {
		fmt.Println("Admin user already exists.")
		return nil
	}

	// Create the initial admin user
	result, err := db.Exec("INSERT INTO users (username) VALUES ('admin')")
	if err != nil {
		return fmt.Errorf("failed to create admin user: %v", err)
	}

	adminUserID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get admin user ID: %v", err)
	}

	fmt.Printf("Admin user created with ID: %d\n", adminUserID)

	return nil
}

func CreateTables(db *sql.DB) error {
	// Create the 'users' table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create 'users' table: %v", err)
	}

	// Create the 'providers' table with 'userId' as a foreign key
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS providers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			provider TEXT NOT NULL,
			token TEXT NOT NULL,
			userId INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users (id)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create 'providers' table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS servers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			host TEXT NOT NULL,
			port INTEGER NOT NULL,
			keyValue TEXT NOT NULL
			FOREIGN KEY (userId) REFERENCES users (id)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create 'providers' table: %v", err)
	}

	fmt.Println("Tables 'users' and 'providers' created successfully.")
	return nil
}

func initDB() error {
	// Open or create an SQLite database file
	db, err := sql.Open("sqlite3", "darkatic-ci.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = CreateTables(db)
	if err != nil {
		return err
	}

	err = CreateInitialAdminUser(db)
	if err != nil {
		return err
	}

	if err != nil {
		log.Fatal(err)
	}

	// Create the initial admin user if it doesn't exist
	err = CreateInitialAdminUser(db)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
