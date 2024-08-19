package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	docs "github.com/printSANO/wish-tree/docs"
	"github.com/printSANO/wish-tree/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// setupRouter sets up the routes for the application.
func setupRouter(handler *handlers.Handler) *gin.Engine {
	router := gin.Default()

	// CORS middleware 다 허용
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	docs.SwaggerInfo.Title = "Wish API"
	docs.SwaggerInfo.Description = "This is a sample server for a Wish API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	apiGroup := router.Group("/api/v1")
	{
		// Post routes
		postGroup := apiGroup.Group("/wishes")
		{
			postGroup.GET("/:id", handler.WishHandler.GetWish)
			postGroup.GET("/pending", handler.WishHandler.GetPendingWishes)
			postGroup.GET("/approved", handler.WishHandler.GetApprovedWishes)
			postGroup.GET("/rejected", handler.WishHandler.GetRejectedWishes)
			postGroup.PATCH("/:id", handler.WishHandler.UpdateWish)
			postGroup.DELETE("/:id", handler.WishHandler.DeleteWish)
			postGroup.POST("/", handler.WishHandler.CreateWish)
		}

		// Comment routes
		commentGroup := apiGroup.Group("/comments")
		{
			commentGroup.GET("/:wish_id", handler.CommentHandler.GetCommentsByWishID)
			commentGroup.POST("/", handler.CommentHandler.CreateComment)
			commentGroup.DELETE("/:id", handler.CommentHandler.DeleteComment)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
