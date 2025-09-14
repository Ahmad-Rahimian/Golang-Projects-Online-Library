package paidbook

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

// @Summary      Get all paidbooks
// @Description  Get list of all paidbooks
// @Tags         paidbook
// @Produce      json
// @Success      200  {array}   paidbook.PaidBook
// @Router       /paidbook [get]
func (h *Handler) GetPaidBooksHandler(c *gin.Context) {
	books, err := GetAll(h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary      Get paidbook by ID
// @Description  Get a paidbook by its ID
// @Tags         paidbook
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  paidbook.PaidBook
// @Router       /paidbook/{id} [get]
func (h *Handler) GetPaidBookHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := GetByID(h.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary      Create new paidbook
// @Description  Add a new paidbook with file upload
// @Tags         paidbook
// @Accept       multipart/form-data
// @Produce      json
// @Param        title       formData  string true  "Book Title"
// @Param        summary     formData  string false "Book Summary"
// @Param        author      formData  string true  "Book Author"
// @Param        pages       formData  int    true  "Book Pages"
// @Param        price       formData  int	  true  "Book Price"
// @Param        cover_image formData  file   true  "Cover Image"
// @Param        pdf_file    formData  file   true  "PDF File"
// @Success      201  {string}  string "created"
// @Router       /paidbook [post]
func (h *Handler) CreatePaidBookHandler(c *gin.Context) {
	var book PaidBook

	book.Title = c.PostForm("title")
	summary := c.PostForm("summary")
	if summary != "" {
		book.Summary = &summary
	}
	book.Author = c.PostForm("author")
	pages, _ := strconv.Atoi(c.PostForm("pages"))
	book.Pages = pages
	price, _ := strconv.Atoi(c.PostForm("price"))
	book.Price = price

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

	if err := CreatePaidBook(h.DB, book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created", "book": book})
}

// @Summary 	Update paidbook
// @Description Update paidbook details by ID (with optional new files)
// @Tags 		paidbook
// @Accept 		multipart/form-data
// @Produce 	json
// @Param 		id 			path 	 int 	true  "Book ID"
// @Param 		title 		formData string false "Book Title"
// @Param 		summary 	formData string false "Book Summary"
// @Param 		author 		formData string false "Book Author"
// @Param 		pages 		formData int 	false "Book Pages"
// @Param 		price 		formData int	false "Book Price"
// @Param 		cover_image formData file 	false "Cover Image"
// @Param 		pdf_file 	formData file 	false "PDF File"
// @Success 	200 {string}	string 	"updated"
// @Router 		/paidbook/{id} [put]
func (h *Handler) UpdatePaidBookHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.PostForm("title")
	summary := c.PostForm("summary")
	author := c.PostForm("author")
	pages, _ := strconv.Atoi(c.PostForm("pages"))
	price, _ := strconv.Atoi(c.PostForm("price"))

	coverFile, _ := c.FormFile("cover_image")
	pdfFile, _ := c.FormFile("pdf_file")

	var coverPath, pdfPath string

	if coverFile != nil {
		coverPath = "uploads/images/" + coverFile.Filename
		if err := c.SaveUploadedFile(coverFile, coverPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save cover image"})
			return
		}
	}

	if pdfFile != nil {
		pdfPath = "uploads/pdfs/" + pdfFile.Filename
		if err := c.SaveUploadedFile(pdfFile, pdfPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save pdf file"})
			return
		}
	}

	oldBook, err := GetByID(h.DB, id)
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

	var summaryPtr *string
	if summary != "" {
		summaryPtr = &summary
	} else {
		summaryPtr = oldBook.Summary
	}

	book := PaidBook{
		ID:          id,
		Title:       ifEmpty(title, oldBook.Title),
		Summary:     summaryPtr,
		Author:      ifEmpty(author, oldBook.Author),
		Cover_image: coverPath,
		Pdf_file:    pdfPath,
		Pages:       ifZero(pages, oldBook.Pages),
		Price:       ifZero(price, oldBook.Price),
	}

	if err := UpdatePaidBook(h.DB, book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated", "book": book})
}

// @Summary      Delete paidbook
// @Description  Delete a paidbook by ID
// @Tags         paidbook
// @Param        id   path      int  true  "Book ID"
// @Success      200  {string}  string "deleted"
// @Router       /paidbook/{id} [delete]
func (h *Handler) DeletePaidBookHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePaidBook(h.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

// helper funcs
func ifEmpty(val, fallback string) string {
	if val == "" {
		return fallback
	}
	return val
}

func ifZero(val, fallback int) int {
	if val == 0 {
		return fallback
	}
	return val
}
