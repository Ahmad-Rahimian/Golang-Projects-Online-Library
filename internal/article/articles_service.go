package article

import (
	"database/sql"
)

// GetArticles
func GetArticles(db *sql.DB) ([]Article, error) {
	return GetAll(db)
}

// GetArticle
func GetArticle(db *sql.DB, id int) (Article, error) {
	return GetByID(db, id)
}

// CreateArticle
func CreateArticle(db *sql.DB, Article Article) error {
	return Create(db, Article)
}

// UpdateArticle
func UpdateArticle(db *sql.DB, Article Article) error {
	return Update(db, Article)
}

// DeleteArticle
func DeleteArticle(db *sql.DB, id int) error {
	return Delete(db, id)
}
