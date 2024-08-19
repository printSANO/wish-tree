package services

import (
	"github.com/printSANO/wish-tree/models"
	"github.com/printSANO/wish-tree/repositories"
)

type WishService interface {
	GetWishByID(id uint) (*models.Wish, error)
	GetAllApprovedWishes(filter models.Filter) ([]*models.Wish, int64, error)
	GetAllPendingWishes(filter models.Filter) ([]*models.Wish, int64, error)
	GetRejectedWishes(filter models.Filter) ([]*models.Wish, int64, error)
	UpdateWishToApproved(id uint) (*models.Wish, error)
	UpdateWishToRejected(id uint) (*models.Wish, error)
	CreateWish(wish *models.Wish) error
	DeleteWish(wish *models.Wish) error
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

func (s *wishService) GetAllApprovedWishes(filter models.Filter) ([]*models.Wish, int64, error) {
	return s.repo.GetAll(models.Approved, filter)
}

func (s *wishService) GetAllPendingWishes(filter models.Filter) ([]*models.Wish, int64, error) {
	return s.repo.GetAll(models.Pending, filter)
}

func (s *wishService) GetRejectedWishes(filter models.Filter) ([]*models.Wish, int64, error) {
	return s.repo.GetAll(models.Rejected, filter)
}

func (s *wishService) UpdateWishToApproved(id uint) (*models.Wish, error) {
	return s.repo.UpdateWish(id, models.Approved)
}

func (s *wishService) UpdateWishToRejected(id uint) (*models.Wish, error) {
	return s.repo.UpdateWish(id, models.Rejected)
}

func (s *wishService) CreateWish(wish *models.Wish) error {
	return s.repo.Create(wish)
}

func (s *wishService) DeleteWish(wish *models.Wish) error {
	return s.repo.Delete(wish)
}
