package services

import (
	"github.com/printSANO/wish-tree/models"
	"github.com/printSANO/wish-tree/repositories"
)

type WishService interface {
	GetWishByID(id uint) (*models.Wish, error)
	GetAllApprovedWishes() ([]*models.Wish, error)
	GetAllPendingWishes() ([]*models.Wish, error)
	CreateWish(post *models.Wish) error
	DeleteWish(post *models.Wish) error
}

type wishService struct {
	repo repositories.WishRepository
}

func NewWishService(repo repositories.WishRepository) WishService {
	return &wishService{repo: repo}
}

func (s *wishService) GetWishByID(id uint) (*models.Wish, error) {
	return s.repo.FindByID(id)
}

func (s *wishService) GetAllApprovedWishes() ([]*models.Wish, error) {
	return s.repo.GetAll(models.Approved)
}

func (s *wishService) GetAllPendingWishes() ([]*models.Wish, error) {
	return s.repo.GetAll(models.Pending)
}

func (s *wishService) CreateWish(post *models.Wish) error {
	return s.repo.Create(post)
}

func (s *wishService) DeleteWish(post *models.Wish) error {
	return s.repo.Delete(post)
}
