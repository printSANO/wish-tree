package server

import (
	"fmt"

	"github.com/printSANO/wish-tree/database"
	"github.com/printSANO/wish-tree/handlers"
	"github.com/printSANO/wish-tree/repositories"
	"github.com/printSANO/wish-tree/services"
)

func Start(port string) {
	db := database.SetupDatabase()
	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	router := setupRouter(handler)
	router.Run(fmt.Sprintf(":%s", port))
}
