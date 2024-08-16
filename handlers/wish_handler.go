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
	post, err := h.service.GetWishByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *WishHandler) CreateWish(c *gin.Context) {
	var postInput struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.BindJSON(&postInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPost := &models.Wish{
		Title:   postInput.Title,
		Content: postInput.Content,
	}

	if err := h.service.CreateWish(newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, newPost)
}
