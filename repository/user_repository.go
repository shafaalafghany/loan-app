package repository

import (
	"fmt"

	"github.com/shafaalafghany/loan-app/model"
	"github.com/shafaalafghany/loan-app/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(*model.User) error
	GetByEmail(string) (*model.User, error)
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
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r.log.Info("creating data user", zap.Any("data", data))
		if errTx := tx.Create(&data).Error; errTx != nil {
			r.log.Error("failed to create new user", zap.Any("error", errTx))
			return errTx
		}

		audit := util.GetDefaultModelAuditLog("CREATE", "user", *data)
		r.log.Info("creating data audit", zap.Any("data", audit))
		if errTx := tx.Create(&audit).Error; errTx != nil {
			r.log.Error("failed to create new audit", zap.Any("error", errTx))
			return errTx
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	r.log.Info("getting user data by email", zap.Any("data", email))
	var user *model.User
	if err := r.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		fmt.Println("ERROR", err)
		return nil, err
	}

	return user, nil
}
