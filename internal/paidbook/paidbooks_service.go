package paidbook

import (
	"database/sql"
)

func GetBooks(db *sql.DB) ([]PaidBook, error) {
	return GetAll(db)
}

func GetBook(db *sql.DB, id int) (PaidBook, error) {
	return GetByID(db, id)
}

func CreateBook(db *sql.DB, book PaidBook) error {
	return Create(db, book)
}

func UpdateBook(db *sql.DB, book PaidBook) error {
	return Update(db, book)
}

func DeleteBook(db *sql.DB, id int) error {
	return Delete(db, id)
}
