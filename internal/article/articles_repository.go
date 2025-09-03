package article

import (
	"database/sql"
	"log"
)

func GetAll(db *sql.DB) ([]Article, error) {
	rows, err := db.Query("SELECT id , title , summary , author, cover_image FROM articles")
	if err != nil {
		log.Println("Error fetching articles :", err)
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Summary, &article.Author, &article.Cover_image); err != nil {
			log.Println("Error Scanning articles : ", err)
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func GetByID(db *sql.DB, id int) (article Article, err error) {
	err = db.QueryRow("SELECT id , title , summary , author, cover_image , pdf_file FROM articles WHERE id=$1", id).Scan(&article.ID, &article.Title, &article.Summary, &article.Author, &article.Cover_image, &article.Pdf_file)
	if err != nil {
		log.Println("id Not Found", err)
	}
	return article, nil
}

func Create(db *sql.DB, article Article) (err error) {
	_, err = db.Exec("INSERT INTO articles (title , summary , author, cover_image , pdf_file) VALUES ($1,$2,$3,$4,$5)", &article.Title, &article.Summary, &article.Author, &article.Cover_image, &article.Pdf_file)
	if err != nil {
		log.Println("Cannot Create article")
	}
	return err
}

func Update(db *sql.DB, article Article) (err error) {
	_, err = db.Exec("UPDATE articles SET title=$1 ,summary=$2 , author=$3 ,cover_image=$4 , pdf_file=$5  WHERE id=$6 ", &article.Title, &article.Summary, &article.Author, &article.Cover_image, &article.Pdf_file, &article.ID)
	if err != nil {
		log.Println("Cannot Update article", err)
	}
	return err
}

func Delete(db *sql.DB, id int) (err error) {
	_, err = db.Exec("DELETE FROM articles WHERE id=$1", id)
	if err != nil {
		log.Println("Cannot Delete article")
	}
	return err
}
