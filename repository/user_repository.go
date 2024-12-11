package repository

import (
	"github.com/shafaalafghany/loan-app/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(*model.User) error
	GetByEmail(string) (*model.User, error)
	GetById(string) (*model.User, error)
}

type UserRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepositoryInterface {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

func (r *UserRepository) Create(data *model.User) error {
	if err := r.db.Create(&data).Error; err != nil {
		r.log.Error("failed to create new user", zap.Any("error", err))
		return err
	}

	return nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	r.log.Info("getting user data by email", zap.Any("data", email))
	var user *model.User
	if err := r.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetById(id string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
