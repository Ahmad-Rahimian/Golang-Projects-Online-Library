package freebook

import (
	"database/sql"
)

// GetFreeBooks get all free books from database and return list of free books
func GetFreeBooks(db *sql.DB) ([]FreeBook, error) {
	return GetAll(db)
}

// GetFreeBook get free book by id from database and return free book
func GetFreeBook(db *sql.DB, id int) (FreeBook, error) {
	return GetByID(db, id)
}

// CreateFreeBook create new free book in database
func CreateFreeBook(db *sql.DB, book FreeBook) error {
	return Create(db, book)
}

// UpdateFreeBook update free book in database
func UpdateFreeBook(db *sql.DB, book FreeBook) error {
	return Update(db, book)
}

// DeleteFreeBook delete free book from database
func DeleteFreeBook(db *sql.DB, id int) error {
	return Delete(db, id)
}
