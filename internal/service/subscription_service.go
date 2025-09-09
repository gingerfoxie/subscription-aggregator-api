package service

import (
	model "subscription-service/internal/models"
	"subscription-service/internal/repository"
)

type SubscriptionService interface {
	Create(sub *model.Subscription) error
	GetByID(id uint) (*model.Subscription, error)
	Update(sub *model.Subscription) error
	Delete(id uint) error
	List(filters map[string]interface{}, page, limit int) ([]model.Subscription, int64, error)
	GetTotalCost(userID, serviceName, start, end string) (int, error)
}

type subscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repo: repo}
}

func (s *subscriptionService) Create(sub *model.Subscription) error {
	return s.repo.Create(sub)
}

func (s *subscriptionService) GetByID(id uint) (*model.Subscription, error) {
	return s.repo.GetByID(id)
}

func (s *subscriptionService) Update(sub *model.Subscription) error {
	return s.repo.Update(sub)
}

func (s *subscriptionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *subscriptionService) List(filters map[string]interface{}, page, limit int) ([]model.Subscription, int64, error) {
	offset := (page - 1) * limit
	if page <= 0 {
		offset = 0
	}
	return s.repo.List(filters, limit, offset)
}

func (s *subscriptionService) GetTotalCost(userID, serviceName, start, end string) (int, error) {
	return s.repo.GetTotalCost(userID, serviceName, start, end)
}
