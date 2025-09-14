package paidbook

import (
	"database/sql"
)

func GetPaidBooks(db *sql.DB) ([]PaidBook, error) {
	return GetAll(db)
}

func GetPaidBook(db *sql.DB, id int) (PaidBook, error) {
	return GetByID(db, id)
}

func CreatePaidBook(db *sql.DB, book PaidBook) error {
	return Create(db, book)
}

func UpdatePaidBook(db *sql.DB, book PaidBook) error {
	return Update(db, book)
}

func DeletePaidBook(db *sql.DB, id int) error {
	return Delete(db, id)
}
