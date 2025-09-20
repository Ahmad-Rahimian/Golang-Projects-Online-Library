package db

import (
	"database/sql"
	"fmt"
	"log"

	"online-library/pkg/config"

	_ "github.com/lib/pq"
)

// init db with config and return db and error
func InitDB(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s ",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLmode)

	// open db
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open database : %v ", err)
	}

	// ping db
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping database : %v", err)
	}

	// log success
	log.Println("Database Connected Successfully")
	return db, nil
}
