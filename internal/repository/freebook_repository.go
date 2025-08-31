package repository

import (
	"database/sql"
	"log"

	"online-library/internal/domain"
)

func GetAllBooks(db *sql.DB) ([]domain.FreeBook, error) {
	rows, err := db.Query("SELECT id , title , summary , author, cover_image , pages FROM free_books")
	if err != nil {
		log.Println("Error fetching books :", err)
		return nil, err
	}
	var books []domain.FreeBook
	for rows.Next() {
		var book domain.FreeBook
		if err := rows.Scan(&book.ID, &book.Title, &book.Summary, &book.Author, &book.Cover_image, &book.Pages); err != nil {
			log.Println("Error Scanning Books : ", err)
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBookByID(db *sql.DB, id int) (book domain.FreeBook, err error) {
	err = db.QueryRow("SELECT id , title , summary , author, cover_image , pdf_file, pages FROM free_books WHERE id=$1", id).Scan(&book.ID, &book.Title, &book.Summary, &book.Author, &book.Cover_image, &book.Pdf_file, &book.Pages)
	if err != nil {
		log.Println("id Not Found", err)
	}
	return book, nil
}

func CreateBook(db *sql.DB, book domain.FreeBook) (err error) {
	_, err = db.Exec("INSERT INTO free_books (title , summary , author, cover_image , pdf_file, pages) VALUES ($1,$2,$3,$4,$5,$6)", &book.Title, &book.Summary, &book.Author, &book.Cover_image, &book.Pdf_file, &book.Pages)
	if err != nil {
		log.Println("Cannot Create Book")
	}
	return err
}

func UpdateBook(db *sql.DB, book domain.FreeBook) (err error) {
	_, err = db.Exec("UPDATE free_books SET title=$1 ,summary=$2 , author=$3 ,cover_image=$4 , pdf_file=$5 ,pages=$6 WHERE id=$7 ", &book.Title, &book.Summary, &book.Author, &book.Cover_image, &book.Pdf_file, &book.Pages, &book.ID)
	if err != nil {
		log.Println("Cannot Update Book", err)
	}
	return err
}

func DeleteBook(db *sql.DB, id int) (err error) {
	_, err = db.Exec("DELETE FROM free_books WHERE id=$1", id)
	if err != nil {
		log.Println("Cannot Delete Book")
	}
	return err
}
