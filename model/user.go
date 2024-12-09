package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	NIK          string         `gorm:"type:varchar(16);unique;not null" json:"nik"`
	FullName     string         `gorm:"type:varchar(100);not null" json:"full_name"`
	LegalName    string         `gorm:"type:varchar(100);not null" json:"legal_name"`
	TempatLahir  string         `gorm:"type:varchar(50);not null" json:"tempat_lahir"`
	TanggalLahir time.Time      `gorm:"not null" json:"tanggal_lahir"`
	Gaji         float64        `gorm:"not null" json:"gaji"`
	FotoKTP      string         `gorm:"type:varchar(255)" json:"foto_ktp"`
	FotoSelfie   string         `gorm:"type:varchar(255)" json:"foto_selfie"`
	Limits       []Limit        `json:"limits"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
