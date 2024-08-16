package repositories

import (
	"github.com/printSANO/wish-tree/models"
	"gorm.io/gorm"
)

type WishRepository interface {
	FindByID(id uint) (*models.Wish, error)
	GetAll(status models.WishStatus) ([]*models.Wish, error)
	Create(post *models.Wish) error
	Delete(post *models.Wish) error
}

type wishRepository struct {
	db *gorm.DB
}

func NewWishRepository(db *gorm.DB) WishRepository {
	return &wishRepository{db: db}
}

func (r *wishRepository) GetAll(status models.WishStatus) ([]*models.Wish, error) {
	var posts []*models.Wish
	err := r.db.Where("is_confirm = ?", status).Find(&posts).Error
	return posts, err
}

func (r *wishRepository) FindByID(id uint) (*models.Wish, error) {
	var post models.Wish
	err := r.db.First(&post, id).Error
	return &post, err
}

func (r *wishRepository) Create(post *models.Wish) error {
	return r.db.Create(post).Error
}

func (r *wishRepository) Delete(post *models.Wish) error {
	return r.db.Delete(post).Error
}
