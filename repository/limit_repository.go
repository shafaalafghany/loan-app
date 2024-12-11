package repository

import (
	"github.com/shafaalafghany/loan-app/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LimitRepositoryInterface interface {
	AddLimits([]*model.Limit) error
}

type LimitRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewLimitRepository(db *gorm.DB, log *zap.Logger) LimitRepositoryInterface {
	return &LimitRepository{
		db:  db,
		log: log,
	}
}

func (r *LimitRepository) AddLimits(data []*model.Limit) error {
	if err := r.db.Create(data).Error; err != nil {
		r.log.Error("error create user limits", zap.Error(err))
		return err
	}

	return nil
}
