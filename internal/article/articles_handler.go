package article

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

// @Summary      Get all articles
// @Description  Get list of all articles
// @Tags         article
// @Produce      json
// @Success      200  {array}   Article
// @Router       /article [get]
func (h *Handler) GetArticlesHandler(c *gin.Context) {
	articles, err := GetAll(h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, articles)
}

// @Summary      Get article by ID
// @Description  Get a article by its ID
// @Tags         article
// @Produce      json
// @Param        id   path      int  true  "Article ID"
// @Success      200  {object}  Article
// @Router       /article/{id} [get]
func (h *Handler) GetArticleHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := GetByID(h.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

// @Summary      Create new article
// @Description  Add a new article with file upload
// @Tags         article
// @Accept       multipart/form-data
// @Produce      json
// @Param        title       formData  string true  "Article Title"
// @Param        summary     formData  string true  "Article Summary"
// @Param        author      formData  string true  "Article Author"
// @Param        cover_image formData  file   true  "Cover Image"
// @Param        pdf_file    formData  file   true  "PDF File"
// @Success      201  {string}  string "created"
// @Router       /article [post]
func (h *Handler) CreateArticleHandler(c *gin.Context) {
	var article Article

	article.Title = c.PostForm("title")
	summary := c.PostForm("summary")
	if summary != "" {
		article.Summary = summary
	}
	article.Author = c.PostForm("author")

	// Handle cover image upload
	coverFile, err := c.FormFile("cover_image")
	if err == nil {
		coverPath := "uploads/images/" + coverFile.Filename
		if err := c.SaveUploadedFile(coverFile, coverPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save cover image"})
			return
		}
		article.Cover_image = coverPath
	}

	// Handle pdf file upload
	pdfFile, err := c.FormFile("pdf_file")
	if err == nil {
		pdfPath := "uploads/pdfs/" + pdfFile.Filename
		if err := c.SaveUploadedFile(pdfFile, pdfPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save pdf file"})
			return
		}
		article.Pdf_file = &pdfPath
	}

	if err := CreateArticle(h.DB, article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article created", "article": article})
}

// @Summary 	Update article
// @Description Update article details by ID (with optional new files)
// @Tags 		article
// @Accept 		multipart/form-data
// @Produce 	json
// @Param 		id 			path 	 int 	true  "Article ID"
// @Param 		title 		formData string false "Article Title"
// @Param 		summary 	formData string false "Article Summary"
// @Param 		author 		formData string false "Article Author"
// @Param 		cover_image formData file 	false "Cover Image"
// @Param 		pdf_file 	formData file 	false "PDF File"
// @Success 	200 {string}	string 	"updated"
// @Router 		/article/{id} [put]
func (h *Handler) UpdateArticleHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.PostForm("title")
	summary := c.PostForm("summary")
	author := c.PostForm("author")

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

	oldArticle, err := GetByID(h.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	if coverPath == "" {
		coverPath = oldArticle.Cover_image
	}
	if pdfPath == "" {
		pdfPath = *oldArticle.Pdf_file
	}

	var summaryPtr string
	if summary != "" {
		summaryPtr = summary
	} else {
		summaryPtr = oldArticle.Summary
	}

	Article := Article{
		ID:          id,
		Title:       ifEmpty(title, oldArticle.Title),
		Summary:     summaryPtr,
		Author:      ifEmpty(author, oldArticle.Author),
		Cover_image: coverPath,
		Pdf_file:    &pdfPath,
	}

	if err := UpdateArticle(h.DB, Article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article updated", "Article": Article})
}

// @Summary      Delete article
// @Description  Delete a article by ID
// @Tags         article
// @Param        id   path      int  true  "Article ID"
// @Success      200  {string}  string "deleted"
// @Router       /article/{id} [delete]
func (h *Handler) DeleteArticleHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteArticle(h.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
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
