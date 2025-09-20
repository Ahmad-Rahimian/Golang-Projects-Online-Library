package paidbook

import (
	"database/sql"
)

// GetPaidBooks get all paid books from database and return list of paid books
func GetPaidBooks(db *sql.DB) ([]PaidBook, error) {
	return GetAll(db)
}

// GetPaidBook get paid book by id from database and return paid book
func GetPaidBook(db *sql.DB, id int) (PaidBook, error) {
	return GetByID(db, id)
}

// CreatePaidBook create new paid book in database
func CreatePaidBook(db *sql.DB, book PaidBook) error {
	return Create(db, book)
}

// UpdatePaidBook update paid book in database
func UpdatePaidBook(db *sql.DB, book PaidBook) error {
	return Update(db, book)
}

// DeletePaidBook delete paid book from database
func DeletePaidBook(db *sql.DB, id int) error {
	return Delete(db, id)
}
