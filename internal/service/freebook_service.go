package service

import (
	"database/sql"

	"online-library/internal/domain"
	"online-library/internal/repository"
)

func GetAllBooks(db *sql.DB) ([]domain.FreeBook, error) {
	return repository.GetAllBooks(db)
}

func GetBookByID(db *sql.DB, id int) (domain.FreeBook, error) {
	return repository.GetBookByID(db, id)
}

func CreateBook(db *sql.DB, c domain.FreeBook) error {
	return repository.CreateBook(db, c)
}

func UpdateBook(db *sql.DB, c domain.FreeBook) error {
	return repository.UpdateBook(db, c)
}

func DeleteBook(db *sql.DB, id int) error {
	return repository.DeleteBook(db, id)
}
