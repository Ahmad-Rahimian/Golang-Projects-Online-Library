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
// @Security 	 BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        title       		formData  string true  "Article Title"
// @Param        short_summary      formData  string true  "Article Short Summary"
// @Param        full_text      	formData  string true  "Article Full Text"
// @Param        author      		formData  string true  "Article Author"
// @Param        cover_image 		formData  file   true  "Cover Image"
// @Param        reading_time      	formData  int	 false  "Article Reading Time"
// @Success      201  {string}  string "created"
// @Router       /article [post]
func (h *Handler) CreateArticleHandler(c *gin.Context) {
	var article Article

	article.Title = c.PostForm("title")
	short_summary := c.PostForm("short_summary")
	if short_summary != "" {
		article.Short_summary = short_summary
	}
	full_text := c.PostForm("full_text")
	if full_text != "" {
		article.Full_text = full_text
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
	reading_time, _ := strconv.Atoi(c.PostForm("reading_time"))
	article.Reading_time = &reading_time

	if err := CreateArticle(h.DB, article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article created", "article": article})
}

// @Summary 	Update article
// @Description Update article details by ID (with optional new files)
// @Tags 		article
// @Security 	BearerAuth
// @Accept 		multipart/form-data
// @Produce 	json
// @Param 		id 				path 	 int 	true  "Article ID"
// @Param 		title 			formData string false "Article Title"
// @Param       short_summary   formData string false  "Article Short Summary"
// @Param       full_text      	formData string false  "Article Full Text"
// @Param 		author 			formData string false "Article Author"
// @Param 		cover_image 	formData file 	false "Cover Image"
// @Param       reading_time    formData int	false  "Article Reading Time"
// @Success 	200 {string}	string 	"updated"
// @Router 		/article/{id} [put]
func (h *Handler) UpdateArticleHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.PostForm("title")
	short_summary := c.PostForm("short_summary")
	full_text := c.PostForm("full_text")
	author := c.PostForm("author")
	coverFile, _ := c.FormFile("cover_image")
	reading_time, _ := strconv.Atoi(c.PostForm("reading_time"))

	var coverPath string
	if coverFile != nil {
		coverPath = "uploads/images/" + coverFile.Filename
		if err := c.SaveUploadedFile(coverFile, coverPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save cover image"})
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

	var short_summaryPtr string
	if short_summary != "" {
		short_summaryPtr = short_summary
	} else {
		short_summaryPtr = oldArticle.Short_summary
	}

	Article := Article{
		ID:            id,
		Title:         ifEmpty(title, oldArticle.Title),
		Short_summary: short_summaryPtr,
		Full_text:     ifEmpty(full_text, oldArticle.Full_text),
		Author:        ifEmpty(author, oldArticle.Author),
		Cover_image:   coverPath,
		Reading_time:  &reading_time,
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
// @Security 	 BearerAuth
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
