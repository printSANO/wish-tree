package handlers

import "github.com/printSANO/wish-tree/services"

// Handler groups all individual handlers.
type Handler struct {
	WishHandler *WishHandler
}

// NewHandler creates a new instance of Handler with all required handlers.
func NewHandler(service *services.Service) *Handler {
	return &Handler{
		WishHandler: NewWishHandler(service.WishService),
	}
}
