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
// @Description  Add a new Book with file upload
// @Tags         books
// @Accept       multipart/form-data
// @Produce      json
// @Param        title       formData  string true  "Book Title"
// @Param        summary     formData  string false "Book Summary"
// @Param        author      formData  string true  "Book Author"
// @Param        pages       formData  int    true  "Book Pages"
// @Param        cover_image formData  file   true  "Cover Image"
// @Param        pdf_file    formData  file   true  "PDF File"
// @Success      201  {string}  string "created"
// @Router       /books [post]
func CreateBookHandler(c *gin.Context, db *sql.DB) {
	var book domain.FreeBook

	book.Title = c.PostForm("title")
	summary := c.PostForm("summary")
	if summary != "" {
		book.Summary = &summary
	}
	book.Author = c.PostForm("author")
	pages, _ := strconv.Atoi(c.PostForm("pages"))
	book.Pages = pages

	// Handle cover image upload
	coverFile, err := c.FormFile("cover_image")
	if err == nil {
		coverPath := "uploads/images/" + coverFile.Filename
		if err := c.SaveUploadedFile(coverFile, coverPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save cover image"})
			return
		}
		book.Cover_image = coverPath
	}

	// Handle pdf file upload
	pdfFile, err := c.FormFile("pdf_file")
	if err == nil {
		pdfPath := "uploads/pdfs/" + pdfFile.Filename
		if err := c.SaveUploadedFile(pdfFile, pdfPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save pdf file"})
			return
		}
		book.Pdf_file = pdfPath
	}

	if err := service.CreateBook(db, book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created", "book": book})
}

// @Summary 	Update Book
// @Description Update Book details by ID (with optional new files)
// @Tags 		books
// @Accept 		multipart/form-data
// @Produce 	json
// @Param 		id 			path 	 int 	true  "Book ID"
// @Param 		title 		formData string false "Book Title"
// @Param 		summary 	formData string false "Book Summary"
// @Param 		author 		formData string false "Book Author"
// @Param 		pages 		formData int 	false "Book Pages"
// @Param 		cover_image formData file 	false "Cover Image"
// @Param 		pdf_file 	formData file 	false "PDF File"
// @Success 	200 {string}	string 	"updated"
// @Router 		/books/{id} [put]
func UpdateBookHandler(c *gin.Context, db *sql.DB) {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.PostForm("title")
	summary := c.PostForm("summary")
	author := c.PostForm("author")
	pages, _ := strconv.Atoi(c.PostForm("pages"))

	coverFile, err1 := c.FormFile("cover_image")
	pdfFile, err2 := c.FormFile("pdf_file")

	var coverPath, pdfPath string

	if err1 == nil {
		coverPath = "/uploads/images" + coverFile.Filename
		c.SaveUploadedFile(coverFile, coverPath)
	}

	if err2 == nil {
		pdfPath = "/uploads/pdfs" + pdfFile.Filename
		c.SaveUploadedFile(pdfFile, pdfPath)
	}

	oldBook, err := service.GetBookByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	if coverPath == "" {
		coverPath = oldBook.Cover_image
	}
	if pdfPath == "" {
		pdfPath = oldBook.Pdf_file
	}

	book := domain.FreeBook{
		ID:          id,
		Title:       title,
		Summary:     &summary,
		Author:      author,
		Cover_image: coverPath,
		Pdf_file:    pdfPath,
		Pages:       pages,
	}

	if err := service.UpdateBook(db, book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated", "book": book})
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
