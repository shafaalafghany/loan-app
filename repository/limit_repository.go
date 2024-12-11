package repository

import (
	"github.com/shafaalafghany/loan-app/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LimitRepositoryInterface interface {
	AddLimits([]*model.Limit) error
	GetByUserId(string) ([]*model.Limit, error)
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

func (r *LimitRepository) GetByUserId(id string) ([]*model.Limit, error) {
	var limits []*model.Limit
	if err := r.db.Where("user_id = ? AND deleted_at IS NULL", id).Find(&limits).Error; err != nil {
		r.log.Error("failed to get user limit", zap.Error(err))
		return nil, err
	}

	return limits, nil
}
