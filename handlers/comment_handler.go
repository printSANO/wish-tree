package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/printSANO/wish-tree/models"
	"github.com/printSANO/wish-tree/services"
)

type CommentHandler struct {
	service services.CommentService
}

func NewCommentHandler(service services.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

// GetCommentsByWishID godoc
// @Summary Get comments by wish ID
// @Description Retrieve comments for a wish by its ID
// @Tags comments
// @Param wish_id path int true "Wish ID"
// @Param limit query int false "Number of items per page"
// @Param page query int false "Page number"
// @Produce json
// @Success 200 {object} models.PaginatedResponse[models.Comment]
// @Failure 404 {object} gin.H
// @Router /comments/{wish_id} [get]
func (h *CommentHandler) GetCommentsByWishID(c *gin.Context) {
	wishID, _ := strconv.Atoi(c.Param("wish_id"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limitAndPage := models.CommentLimit{Limit: limit, Page: page}
	comments, err := h.service.GetCommentsByWishID(uint(wishID), limitAndPage)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comments not found"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// CreateComment godoc
// @Summary Create a comment
// @Description Create a new comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.CreateCommentRequest true "Comment content and wish ID"
// @Success 201 {object} models.Comment
// @Failure 400 {object} gin.H
// @Router /comments [post]
func (h *CommentHandler) CreateComment(c *gin.Context) {
	var req models.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := models.Comment{
		Content: req.Content,
		WishID:  req.WishID,
	}

	if err := h.service.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Delete a comment by its ID
// @Tags comments
// @Param id path int true "Comment ID"
// @Produce json
// @Success 204
// @Failure 404 {object} gin.H
// @Router /comments/{id} [delete]
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment, err := h.service.GetCommentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	if err := h.service.DeleteComment(comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}
	c.Status(http.StatusNoContent)
}
