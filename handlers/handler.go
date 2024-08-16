package handlers

import "github.com/printSANO/wish-tree/services"

// Handler groups all individual handlers.
type Handler struct {
	UserHandler *UserHandler
	PostHandler *PostHandler
}

// NewHandler creates a new instance of Handler with all required handlers.
func NewHandler(service *services.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(service.UserService),
		PostHandler: NewPostHandler(service.PostService),
	}
}
