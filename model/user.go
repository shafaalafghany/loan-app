package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	Email        string         `gorm:"type.varchar(100); unique;not null" json:"email"`
	Password     string         `gorm:"type.varchar(255);not null" json:"-"`
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

type UserRequest struct {
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	NIK          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Gaji         float64   `json:"gaji"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	User    User   `json:"data"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
