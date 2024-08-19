package server

import (
	"github.com/gin-gonic/gin"
	docs "github.com/printSANO/wish-tree/docs"
	"github.com/printSANO/wish-tree/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// setupRouter sets up the routes for the application.
func setupRouter(handler *handlers.Handler) *gin.Engine {
	router := gin.Default()

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
			postGroup.GET("/rejected", handler.WishHandler.GetPendingWishes)
			postGroup.PATCH("/:id", handler.WishHandler.UpdateWish)
			postGroup.DELETE("/:id", handler.WishHandler.DeleteWish)
			postGroup.POST("/", handler.WishHandler.CreateWish)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
