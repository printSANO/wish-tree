package server

import (
	"github.com/gin-gonic/gin"
	"github.com/printSANO/wish-tree/handlers"
)

// 라우터 설정
func setupRouter(handler *handlers.Handler) *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	{
		// Post routes
		postGroup := apiGroup.Group("/wishes")
		{
			postGroup.GET("/:id", handler.WishHandler.GetWish)
			postGroup.POST("/", handler.WishHandler.CreateWish)
		}
	}
	return router
}
