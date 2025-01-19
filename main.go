package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// User represents a user in our database
type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}

func main() {
	// Database connection parameters
	const (
		dbUser     = "user"
		dbPassword = "password"
		dbHost     = "localhost"
		dbPort     = "3306"
		dbName     = "myapp"
	)

	// Create the connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open database connection using sqlx
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Example: Create users table
	schema := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Example: Insert a user
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	_, err = db.NamedExec(`
		INSERT INTO users (name, email) 
		VALUES (:name, :email)`,
		user)
	if err != nil {
		log.Printf("Failed to insert user: %v", err)
	}

	// Example: Query a single user
	var fetchedUser User
	err = db.Get(&fetchedUser, "SELECT * FROM users WHERE email = ?", "john@example.com")
	if err != nil {
		log.Printf("Failed to fetch user: %v", err)
	} else {
		fmt.Printf("Found user: %+v\n", fetchedUser)
	}

	// Example: Query multiple users
	var users []User
	err = db.Select(&users, "SELECT * FROM users ORDER BY created_at DESC LIMIT 10")
	if err != nil {
		log.Printf("Failed to fetch users: %v", err)
	} else {
		fmt.Printf("Found %d users\n", len(users))
		for _, u := range users {
			fmt.Printf("- %s (%s)\n", u.Name, u.Email)
		}
	}
}
