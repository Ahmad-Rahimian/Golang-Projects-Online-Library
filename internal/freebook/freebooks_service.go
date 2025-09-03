package freebook

import (
	"database/sql"
)

func GetBooks(db *sql.DB) ([]FreeBook, error) {
	return GetAll(db)
}

func GetBook(db *sql.DB, id int) (FreeBook, error) {
	return GetByID(db, id)
}

func CreateBook(db *sql.DB, book FreeBook) error {
	return Create(db, book)
}

func UpdateBook(db *sql.DB, book FreeBook) error {
	return Update(db, book)
}

func DeleteBook(db *sql.DB, id int) error {
	return Delete(db, id)
}
