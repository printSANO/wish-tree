package repositories

import (
	"github.com/printSANO/wish-tree/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	FindByID(id uint) (*models.Comment, error)
	FindByWishID(wishID uint, limit models.CommentLimit) ([]*models.Comment, error)
	Create(comment *models.Comment) error
	Delete(comment *models.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) FindByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.First(&comment, id).Error
	return &comment, err
}

func (r *commentRepository) FindByWishID(wishID uint, limit models.CommentLimit) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := r.db.Where("wish_id = ?", wishID).Limit(limit.Limit).Offset((limit.Page - 1) * limit.Limit).Find(&comments).Error
	return comments, err
}

func (r *commentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) Delete(comment *models.Comment) error {
	return r.db.Delete(comment).Error
}
