package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

// create admin user with username, password and role admin and return id, username, password and role
func main() {
	connStr := "host=localhost port=5432 user=admin password=admin-pass dbname=online-library sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	username := "admin"
	password := "admin123"
	role := "admin"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	var id int
	err = db.QueryRow(
		`INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id`,
		username, string(hash), role,
	).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Admin created with ID %d, username: %s, password: %s\n", id, username, password)
}
