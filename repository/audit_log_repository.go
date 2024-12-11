package repository

import (
	"github.com/shafaalafghany/loan-app/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuditLogRepositoryInterface interface {
	Add(*model.AuditLog) error
}

type AuditLogRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewAuditLogRepository(db *gorm.DB, log *zap.Logger) AuditLogRepositoryInterface {
	return &AuditLogRepository{
		db:  db,
		log: log,
	}
}

func (r *AuditLogRepository) Add(data *model.AuditLog) error {
	if err := r.db.Create(&data).Error; err != nil {
		r.log.Info("failed to create new audit log", zap.Any("error", err))
		return err
	}

	return nil
}
