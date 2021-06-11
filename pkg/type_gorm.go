package pkg

import (
	"time"

	"gorm.io/gorm"
)

// GormModel base model
type GormModel struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"createdAt" sql:"index"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" sql:"index"`
}
