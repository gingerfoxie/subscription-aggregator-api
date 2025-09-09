package model

import (
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
