package repositories

import (
	"github.com/printSANO/wish-tree/models"
	"gorm.io/gorm"
)

type WishRepository interface {
	FindByID(id uint) (*models.Wish, error)
	GetAll(status models.WishStatus, filter models.Filter) ([]*models.Wish, int64, error)
	UpdateWish(id uint, status models.WishStatus) (*models.Wish, error)
	Create(wish *models.Wish) error
	Delete(wish *models.Wish) error
}

type wishRepository struct {
	db *gorm.DB
}

func NewWishRepository(db *gorm.DB) WishRepository {
	return &wishRepository{db: db}
}

func (r *wishRepository) GetAll(status models.WishStatus, filter models.Filter) ([]*models.Wish, int64, error) {
	var wishes []*models.Wish
	var totalCount int64

	// Base query with status
	query := r.db.Where("is_confirm = ?", status)

	// Apply filters
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if filter.Title != "" {
		query = query.Where("title LIKE ?", "%"+filter.Title+"%")
	}
	if filter.Content != "" {
		query = query.Where("content LIKE ?", "%"+filter.Content+"%")
	}
	query.Model(&models.Wish{}).Count(&totalCount)
	offset := (filter.Page - 1) * filter.Limit
	err := query.Limit(filter.Limit).Offset(offset).Find(&wishes).Error

	return wishes, totalCount, err
}

func (r *wishRepository) FindByID(id uint) (*models.Wish, error) {
	var wish models.Wish
	err := r.db.First(&wish, id).Error
	return &wish, err
}

func (r *wishRepository) UpdateWish(id uint, status models.WishStatus) (*models.Wish, error) {
	var wish models.Wish
	err := r.db.First(&wish, id).Error
	if err != nil {
		return nil, err
	}
	wish.IsConfirm = status
	err = r.db.Save(&wish).Error
	return &wish, err
}

func (r *wishRepository) Create(wish *models.Wish) error {
	return r.db.Create(wish).Error
}

func (r *wishRepository) Delete(wish *models.Wish) error {
	return r.db.Delete(wish).Error
}
