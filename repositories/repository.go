package repositories

import "gorm.io/gorm"

type Repository struct {
	WishRepository WishRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		WishRepository: NewWishRepository(db),
	}
}
