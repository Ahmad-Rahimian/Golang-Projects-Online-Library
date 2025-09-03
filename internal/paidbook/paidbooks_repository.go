package paidbook

import (
	"database/sql"
	"log"
)

func GetAll(db *sql.DB) ([]PaidBook, error) {
	rows, err := db.Query("SELECT id , title , summary , author, cover_image , pages,price FROM paid_books")
	if err != nil {
		log.Println("Error fetching books :", err)
		return nil, err
	}
	defer rows.Close()

	var books []PaidBook
	for rows.Next() {
		var book PaidBook
		if err := rows.Scan(&book.ID, &book.Title, &book.Summary, &book.Author, &book.Cover_image, &book.Pages, &book.Price); err != nil {
			log.Println("Error Scanning Books : ", err)
		}
		books = append(books, book)
	}
	return books, nil
}

func GetByID(db *sql.DB, id int) (book PaidBook, err error) {
	err = db.QueryRow("SELECT id , title , summary , author, cover_image , pdf_file, pages , price FROM paid_books WHERE id=$1", id).Scan(&book.ID, &book.Title, &book.Summary, &book.Author, &book.Cover_image, &book.Pdf_file, &book.Pages, &book.Price)
	if err != nil {
		log.Println("id Not Found", err)
	}
	return book, nil
}

func Create(db *sql.DB, book PaidBook) (err error) {
	_, err = db.Exec("INSERT INTO paid_books (title , summary , author, cover_image , pdf_file, pages,price) VALUES ($1,$2,$3,$4,$5,$6,$7)", &book.Title, &book.Summary, &book.Author, &book.Cover_image, &book.Pdf_file, &book.Pages, &book.Price)
	if err != nil {
		log.Println("Cannot Create Book")
	}
	return err
}

func Update(db *sql.DB, book PaidBook) (err error) {
	_, err = db.Exec("UPDATE paid_books SET title=$1 ,summary=$2 , author=$3 ,cover_image=$4 , pdf_file=$5 ,pages=$6 ,price=$7 WHERE id=$8 ", &book.Title, &book.Summary, &book.Author, &book.Cover_image, &book.Pdf_file, &book.Pages, &book.Price, &book.ID)
	if err != nil {
		log.Println("Cannot Update Book", err)
	}
	return err
}

func Delete(db *sql.DB, id int) (err error) {
	_, err = db.Exec("DELETE FROM paid_books WHERE id=$1", id)
	if err != nil {
		log.Println("Cannot Delete Book")
	}
	return err
}
