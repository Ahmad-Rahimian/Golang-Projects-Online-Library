package freebook

import (
	"database/sql"
)

func GetFreeBooks(db *sql.DB) ([]FreeBook, error) {
	return GetAll(db)
}

func GetFreeBook(db *sql.DB, id int) (FreeBook, error) {
	return GetByID(db, id)
}

func CreateFreeBook(db *sql.DB, book FreeBook) error {
	return Create(db, book)
}

func UpdateFreeBook(db *sql.DB, book FreeBook) error {
	return Update(db, book)
}

func DeleteFreeBook(db *sql.DB, id int) error {
	return Delete(db, id)
}
