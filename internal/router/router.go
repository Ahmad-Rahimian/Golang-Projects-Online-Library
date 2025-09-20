package router

import (
	"database/sql"

	"online-library/internal/article"
	"online-library/internal/auth"
	"online-library/internal/freebook"
	"online-library/internal/middleware"
	"online-library/internal/paidbook"
	"online-library/internal/user"
	"online-library/pkg/config"
	"online-library/pkg/redis"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// setup router with all routes and middleware and return gin engine
func SetupRouter(db *sql.DB, cfg *config.Config, jwtSecret string) *gin.Engine {
	r := gin.Default()

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// redis client
	rdb := redis.NewClient(*cfg)

	// user repo + service
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, rdb)

	// --- Auth routes ---
	authHandler := auth.NewHandler(userService, jwtSecret)
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/send-otp", authHandler.SendOTP)
		authRoutes.POST("/verify-otp", authHandler.VerifyOTP)
	}

	// --- FreeBook routes ---
	fbHandler := freebook.NewHandler(db)
	freebookRoutes := r.Group("/freebook")
	{
		freebookRoutes.GET("", fbHandler.GetFreeBooksHandler)
		freebookRoutes.GET("/:id", fbHandler.GetFreeBookHandler)

		freebookRoutes.POST("", middleware.AuthMiddleware(jwtSecret), middleware.AdminOnly(), fbHandler.CreateFreeBookHandler)
		freebookRoutes.PUT("/:id", middleware.AuthMiddleware(jwtSecret), middleware.AdminOnly(), fbHandler.UpdateFreeBookHandler)
		freebookRoutes.DELETE("/:id", middleware.AuthMiddleware(jwtSecret), middleware.AdminOnly(), fbHandler.DeleteFreeBookHandler)
	}

	// --- PaidBook routes ---
	pbHandler := paidbook.NewHandler(db)
	paidbookRoutes := r.Group("/paidbook")
	{
		paidbookRoutes.GET("", pbHandler.GetPaidBooksHandler)
		paidbookRoutes.GET("/:id", pbHandler.GetPaidBookHandler)

		paidbookRoutes.POST("", middleware.AuthMiddleware(jwtSecret), middleware.AdminOnly(), pbHandler.CreatePaidBookHandler)
		paidbookRoutes.PUT("/:id", middleware.AuthMiddleware(jwtSecret), middleware.AdminOnly(), pbHandler.UpdatePaidBookHandler)
		paidbookRoutes.DELETE("/:id", middleware.AuthMiddleware(jwtSecret), middleware.AdminOnly(), pbHandler.DeletePaidBookHandler)
	}

	// --- Article routes ---
	arHandler := article.NewHandler(db)
	articleRoutes := r.Group("/article")
	{
		articleRoutes.GET("", arHandler.GetArticlesHandler)
		articleRoutes.GET("/:id", arHandler.GetArticleHandler)

		articleRoutes.POST("", middleware.AuthMiddleware(jwtSecret), arHandler.CreateArticleHandler)
		articleRoutes.PUT("/:id", middleware.AuthMiddleware(jwtSecret), arHandler.UpdateArticleHandler)
		articleRoutes.DELETE("/:id", middleware.AuthMiddleware(jwtSecret), arHandler.DeleteArticleHandler)
	}

	return r
}
