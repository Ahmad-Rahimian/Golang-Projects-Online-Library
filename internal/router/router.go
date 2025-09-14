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
		freebookRoutes.GET("", fbHandler.GetFreeBooksHandler)          // GET /freebook
		freebookRoutes.GET("/:id", fbHandler.GetFreeBookHandler)       // GET /freebook/:id
		freebookRoutes.POST("", fbHandler.CreateFreeBookHandler)       // POST /freebook
		freebookRoutes.PUT("/:id", fbHandler.UpdateFreeBookHandler)    // PUT /freebook/:id
		freebookRoutes.DELETE("/:id", fbHandler.DeleteFreeBookHandler) // DELETE /freebook/:id
	}

	pbHandler := paidbook.NewHandler(db)

	paidbookRoutes := r.Group("/paidbook")
	{
		paidbookRoutes.GET("", pbHandler.GetPaidBooksHandler)          // GET /paidbook
		paidbookRoutes.GET("/:id", pbHandler.GetPaidBookHandler)       // GET /paidbook/:id
		paidbookRoutes.POST("", pbHandler.CreatePaidBookHandler)       // POST /paidbook
		paidbookRoutes.PUT("/:id", pbHandler.UpdatePaidBookHandler)    // PUT /paidbook/:id
		paidbookRoutes.DELETE("/:id", pbHandler.DeletePaidBookHandler) // DELETE /paidbook/:id
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
