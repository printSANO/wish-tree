package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/printSANO/wish-tree/models"
	"github.com/printSANO/wish-tree/services"
)

type WishHandler struct {
	service services.WishService
}

func NewWishHandler(service services.WishService) *WishHandler {
	return &WishHandler{service: service}
}

func (h *WishHandler) GetWish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	wish, err := h.service.GetWishByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wish not found"})
		return
	}
	c.JSON(http.StatusOK, wish)
}

func (h *WishHandler) CreateWish(c *gin.Context) {
	var wish models.Wish
	if err := c.ShouldBindJSON(&wish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateWish(&wish); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create wish"})
		return
	}
	c.JSON(http.StatusCreated, wish)
}
