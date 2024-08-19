package services

import (
	"github.com/printSANO/wish-tree/models"
	"github.com/printSANO/wish-tree/repositories"
)

// CommentService provides methods for managing comments.
type CommentService interface {
	GetCommentByID(id uint) (*models.Comment, error)
	GetCommentsByWishID(wishID uint, limit models.CommentLimit) ([]*models.Comment, error)
	CreateComment(comment *models.Comment) error
	DeleteComment(comment *models.Comment) error
}

type commentService struct {
	repo repositories.CommentRepository
}

// NewCommentService creates a new CommentService.
func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

// GetCommentByID returns a comment by its ID.
func (s *commentService) GetCommentByID(id uint) (*models.Comment, error) {
	return s.repo.FindByID(id)
}

// GetCommentsByWishID returns comments for a wish by its ID.
func (s *commentService) GetCommentsByWishID(wishID uint, limit models.CommentLimit) ([]*models.Comment, error) {
	return s.repo.FindByWishID(wishID, limit)
}

// CreateComment creates a new comment.
func (s *commentService) CreateComment(comment *models.Comment) error {
	return s.repo.Create(comment)
}

// DeleteComment deletes a comment.
func (s *commentService) DeleteComment(comment *models.Comment) error {
	return s.repo.Delete(comment)
}
