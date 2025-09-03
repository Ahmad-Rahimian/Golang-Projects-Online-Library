package article

import (
	"database/sql"
)

func GetArticles(db *sql.DB) ([]Article, error) {
	return GetAll(db)
}

func GetArticle(db *sql.DB, id int) (Article, error) {
	return GetByID(db, id)
}

func CreateArticle(db *sql.DB, Article Article) error {
	return Create(db, Article)
}

func UpdateArticle(db *sql.DB, Article Article) error {
	return Update(db, Article)
}

func DeleteArticle(db *sql.DB, id int) error {
	return Delete(db, id)
}
