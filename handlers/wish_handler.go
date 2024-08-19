package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/printSANO/wish-tree/models"
	"github.com/printSANO/wish-tree/services"
)

// WishHandler handles requests related to wishes.
type WishHandler struct {
	service services.WishService
}

func NewWishHandler(service services.WishService) *WishHandler {
	return &WishHandler{service: service}
}

// GetWish godoc
// @Summary Get a wish by ID
// @Description Retrieve a wish by its ID
// @Tags wishes
// @Param id path int true "Wish ID"
// @Produce json
// @Success 200 {object} models.Wish
// @Failure 404 {object} gin.H
// @Router /wishes/{id} [get]
func (h *WishHandler) GetWish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	wish, err := h.service.GetWishByID(uint(id))
	if wish.IsConfirm != models.Approved {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wish not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wish not found"})
		return
	}
	c.JSON(http.StatusOK, wish)
}

// GetApprovedWishes godoc
// @Summary Get all approved wishes
// @Description Retrieve all wishes with the status 'approved'
// @Tags wishes
// @Param category query string false "Category"
// @Param title query string false "Title"
// @Param content query string false "Content"
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Produce json
// @Success 200 {object} models.PaginatedResponse[models.Wish]
// @Failure 400 {object} gin.H
// @Router /wishes/approved [get]
func (h *WishHandler) GetApprovedWishes(c *gin.Context) {
	var filter models.Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	wishes, totalCount, err := h.service.GetAllApprovedWishes(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch wishes"})
		return
	}
	response := models.PaginatedResponse[*models.Wish]{
		Data: wishes,
		Pagination: models.Pagination{
			Total: totalCount,
			Page:  filter.Page,
			Limit: filter.Limit,
		},
	}
	c.JSON(http.StatusOK, response)
}

// GetPendingWishes godoc
// @Summary Get all pending wishes
// @Description Retrieve all wishes with the status 'pending'
// @Tags wishes
// @Param category query string false "Category"
// @Param title query string false "Title"
// @Param content query string false "Content"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Produce json
// @Success 200 {object} models.PaginatedResponse[models.Wish]
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /wishes/pending [get]
func (h *WishHandler) GetPendingWishes(c *gin.Context) {
	var filter models.Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	wishes, totalCount, err := h.service.GetAllPendingWishes(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch wishes"})
		return
	}
	response := models.PaginatedResponse[*models.Wish]{
		Data: wishes,
		Pagination: models.Pagination{
			Total: totalCount,
			Page:  filter.Page,
			Limit: filter.Limit,
		},
	}
	c.JSON(http.StatusOK, response)
}

// GetRejectedWishes godoc
// @Summary Get all rejected wishes
// @Description Retrieve all wishes with the status 'rejected'
// @Tags wishes
// @Param category query string false "Category"
// @Param title query string false "Title"
// @Param content query string false "Content"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Produce json
// @Success 200 {object} models.PaginatedResponse[models.Wish]
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /wishes/rejected [get]
func (h *WishHandler) GetRejectedWishes(c *gin.Context) {
	var filter models.Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	wishes, totalCount, err := h.service.GetRejectedWishes(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch wishes"})
		return
	}
	response := models.PaginatedResponse[*models.Wish]{
		Data: wishes,
		Pagination: models.Pagination{
			Total: totalCount,
			Page:  filter.Page,
			Limit: filter.Limit,
		},
	}
	c.JSON(http.StatusOK, response)
}

// CreateWish godoc
// @Summary Create a new wish
// @Description Create a new wish with the provided category, content, and title
// @Tags wishes
// @Accept json
// @Produce json
// @Param wish body models.CreateWishRequest true "Wish data"
// @Success 201 {object} models.Wish
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /wishes [post]
func (h *WishHandler) CreateWish(c *gin.Context) {
	// Create a variable of the input struct to validate the request body
	var req models.CreateWishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wish := models.Wish{
		Category: req.Category,
		Content:  req.Content,
		Title:    req.Title,
	}

	if err := h.service.CreateWish(&wish); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create wish"})
		return
	}

	c.JSON(http.StatusCreated, wish)
}

// UpdateWish godoc
// @Summary Update the status of a wish
// @Description Update the status of a wish by its ID
// @Tags wishes
// @Param id path int true "Wish ID"
// @Param status query string true "New status of the wish (approve/reject)"
// @Produce json
// @Success 200 {object} models.Wish
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /wishes/{id} [patch]
func (h *WishHandler) UpdateWish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status := c.Query("status")
	wish, err := h.service.GetWishByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wish not found"})
		return
	}
	if wish.IsConfirm != models.Pending {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wish is not pending"})
		return
	}
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}
	switch status {
	case "approve":
		wish, err := h.service.UpdateWishToApproved(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update wish"})
			return
		}
		c.JSON(http.StatusOK, wish)
	case "reject":
		wish, err := h.service.UpdateWishToRejected(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update wish"})
			return
		}
		c.JSON(http.StatusOK, wish)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
	}
}

// DeleteWish godoc
// @Summary Delete a wish
// @Description Delete a wish by its ID
// @Tags wishes
// @Param id path int true "Wish ID"
// @Produce json
// @Success 200 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /wishes/{id} [delete]
func (h *WishHandler) DeleteWish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	wish, err := h.service.GetWishByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wish not found"})
		return
	}
	if err := h.service.DeleteWish(wish); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete wish"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Wish deleted successfully"})
}
