package router

import (
	"database/sql"

	"online-library/internal/article"
	"online-library/internal/freebook"
	"online-library/internal/paidbook"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// swagger

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fbHandler := freebook.NewHandler(db)

	freebookRoutes := r.Group("/freebook")
	{
		freebookRoutes.GET("", fbHandler.GetBooksHandler)          // GET /freebook
		freebookRoutes.GET("/:id", fbHandler.GetBookHandler)       // GET /freebook/:id
		freebookRoutes.POST("", fbHandler.CreateBookHandler)       // POST /freebook
		freebookRoutes.PUT("/:id", fbHandler.UpdateBookHandler)    // PUT /freebook/:id
		freebookRoutes.DELETE("/:id", fbHandler.DeleteBookHandler) // DELETE /freebook/:id
	}

	pbHandler := paidbook.NewHandler(db)

	paidbookRoutes := r.Group("/paidbook")
	{
		paidbookRoutes.GET("", pbHandler.GetBooksHandler)          // GET /paidbook
		paidbookRoutes.GET("/:id", pbHandler.GetBookHandler)       // GET /paidbook/:id
		paidbookRoutes.POST("", pbHandler.CreateBookHandler)       // POST /paidbook
		paidbookRoutes.PUT("/:id", pbHandler.UpdateBookHandler)    // PUT /paidbook/:id
		paidbookRoutes.DELETE("/:id", pbHandler.DeleteBookHandler) // DELETE /paidbook/:id
	}

	arHandler := article.NewHandler(db)

	articleRoutes := r.Group("/article")
	{
		articleRoutes.GET("", arHandler.GetArticlesHandler)          // GET /article
		articleRoutes.GET("/:id", arHandler.GetArticleHandler)       // GET /article/:id
		articleRoutes.POST("", arHandler.CreateArticleHandler)       // POST /article
		articleRoutes.PUT("/:id", arHandler.UpdateArticleHandler)    // PUT /article/:id
		articleRoutes.DELETE("/:id", arHandler.DeleteArticleHandler) // DELETE /article/:id
	}

	return r
}
