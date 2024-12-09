package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Limit struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	Tenor     int            `gorm:"not null" json:"tenor"`
	Amount    float64        `gorm:"not null" json:"amount"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
