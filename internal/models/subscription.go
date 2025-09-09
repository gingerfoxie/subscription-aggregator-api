package models

import (
	"fmt"
	"regexp"
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ServiceName string         `json:"service_name" gorm:"not null;type:varchar(100)"`
	Price       int            `json:"price" gorm:"not null"` // в рублях
	UserID      string         `json:"user_id" gorm:"not null;type:uuid"`
	StartDate   string         `json:"start_date" gorm:"not null;type:varchar(7)"` // формат "MM-YYYY"
	EndDate     *string        `json:"end_date,omitempty" gorm:"type:varchar(7)"`  // опционально
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func (s *Subscription) Validate() error {
	if s.ServiceName == "" {
		return fmt.Errorf("service_name is required")
	}
	if len(s.ServiceName) > 100 {
		return fmt.Errorf("service_name must be <= 100 characters")
	}
	if s.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	if s.UserID == "" {
		return fmt.Errorf("user_id is required")
	}
	if !isValidMonthYear(s.StartDate) {
		return fmt.Errorf("start_date must be in MM-YYYY format, got: %s", s.StartDate)
	}
	if s.EndDate != nil && !isValidMonthYear(*s.EndDate) {
		return fmt.Errorf("end_date must be in MM-YYYY format, got: %s", *s.EndDate)
	}
	return nil
}

func isValidMonthYear(s string) bool {
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])-\d{4}$`)
	return re.MatchString(s)
}
