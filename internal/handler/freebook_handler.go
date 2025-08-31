package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"online-library/internal/domain"
	"online-library/internal/service"

	"github.com/gin-gonic/gin"
)

// @Summary      Get all books
// @Description  Get list of all books
// @Tags         books
// @Produce      json
// @Success      200  {array}   domain.FreeBook
// @Router       /books [get]
func GetBooksHandler(c *gin.Context, db *sql.DB) {
	books, err := service.GetAllBooks(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary      Get Book by ID
// @Description  Get a Book by its ID
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  domain.FreeBook
// @Router       /books/{id} [get]
func GetBookHandler(c *gin.Context, db *sql.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	books, err := service.GetBookByID(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary      Create new Book
// @Description  Add a new Book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        Book  body      domain.FreeBook  true  "Book data"
// @Success      201  {string}  string "created"
// @Router       /books [post]
func CreateBookHandler(c *gin.Context, db *sql.DB) {
	var Book domain.FreeBook
	if err := c.ShouldBindJSON(&Book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateBook(db, Book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Book created"})
}

// @Summary      Update Book
// @Description  Update Book details by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Book ID"
// @Param        Book  body      domain.FreeBook  true  "Updated Book"
// @Success      200  {string}  string "updated"
// @Router       /books/{id} [put]
func UpdateBookHandler(c *gin.Context, db *sql.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Book domain.FreeBook
	if err := c.ShouldBindJSON(&Book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Book.ID = id
	if err := service.UpdateBook(db, Book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book updated"})
}

// @Summary      Delete Book
// @Description  Delete a Book by ID
// @Tags         books
// @Param        id   path      int  true  "Book ID"
// @Success      200  {string}  string "deleted"
// @Router       /books/{id} [delete]
func DeleteBookHandler(c *gin.Context, db *sql.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteBook(db, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
