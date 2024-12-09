package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	ContractNo  string         `gorm:"type:varchar(50);not null" json:"contract_no"`
	OTR         float64        `gorm:"not null" json:"otr"`
	AdminFee    float64        `gorm:"not null" json:"admin_fee"`
	Installment float64        `gorm:"not null" json:"installment"`
	Interest    float64        `gorm:"not null" json:"interest"`
	AssetName   string         `gorm:"type:varchar(100);not null" json:"asset_name"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
