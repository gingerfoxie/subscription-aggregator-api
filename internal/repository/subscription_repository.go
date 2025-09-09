package repository

import (
	model "subscription-service/internal/models"

	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	Create(sub *model.Subscription) error
	GetByID(id uint) (*model.Subscription, error)
	Update(sub *model.Subscription) error
	Delete(id uint) error
	List(filters map[string]interface{}, limit, offset int) ([]model.Subscription, int64, error)
	GetTotalCost(userID, serviceName, start, end string) (int, error)
}

type subscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

func (r *subscriptionRepository) Create(sub *model.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *subscriptionRepository) GetByID(id uint) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.db.First(&sub, id).Error
	return &sub, err
}

func (r *subscriptionRepository) Update(sub *model.Subscription) error {
	return r.db.Save(sub).Error
}

func (r *subscriptionRepository) Delete(id uint) error {
	return r.db.Delete(&model.Subscription{}, id).Error
}

func (r *subscriptionRepository) List(filters map[string]interface{}, limit, offset int) ([]model.Subscription, int64, error) {
	var subscriptions []model.Subscription
	var total int64

	db := r.db.Model(&model.Subscription{}).Where("deleted_at IS NULL")

	if userID, ok := filters["user_id"].(string); ok {
		db = db.Where("user_id = ?", userID)
	}
	if serviceName, ok := filters["service_name"].(string); ok {
		db = db.Where("service_name ILIKE ?", "%"+serviceName+"%")
	}

	db.Count(&total)

	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err := db.Find(&subscriptions).Error
	return subscriptions, total, err
}

// GetTotalCost — суммирует стоимость подписок в указанном диапазоне
func (r *subscriptionRepository) GetTotalCost(userID, serviceName, start, end string) (int, error) {
	var total int
	db := r.db.Model(&model.Subscription{}).Where("deleted_at IS NULL")

	if userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		db = db.Where("service_name ILIKE ?", "%"+serviceName+"%")
	}

	// Фильтрация по дате
	// Подписка активна, если:
	// - start_date <= end AND (end_date >= start OR end_date IS NULL)
	db = db.Where("start_date <= ?", end)
	db = db.Where("end_date IS NULL OR end_date >= ?", start)

	var subscriptions []model.Subscription
	err := db.Find(&subscriptions).Error
	if err != nil {
		return 0, err
	}

	for _, sub := range subscriptions {
		total += sub.Price
	}
	return total, nil
}
