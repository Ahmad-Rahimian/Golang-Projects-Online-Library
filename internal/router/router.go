package router

import (
	"database/sql"

	"online-library/internal/handler"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// swagger

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Online Library API Routes
	r.GET("/books", func(c *gin.Context) { handler.GetBooksHandler(c, db) })
	r.GET("/books/:id", func(c *gin.Context) { handler.GetBookHandler(c, db) })
	r.POST("/books", func(c *gin.Context) { handler.CreateBookHandler(c, db) })
	r.PUT("/books/:id", func(c *gin.Context) { handler.UpdateBookHandler(c, db) })
	r.DELETE("/books/:id", func(c *gin.Context) { handler.DeleteBookHandler(c, db) })

	return r
}
