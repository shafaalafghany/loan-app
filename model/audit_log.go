package model

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;"`
	Action    string    `gorm:"type:varchar(50);not null;"`
	Resource  string    `gorm:"type:varchar(100);not null;"`
	Timestamp time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP;"`
}
