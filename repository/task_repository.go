package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

type TaskRepository struct {
	lib.Database
	logger lib.Logger
}

// NewUserRepository creates a new user repository
func NewTaskRepository(db lib.Database, logger lib.Logger) TaskRepository {
	return TaskRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r TaskRepository) WithTrx(trxHandle *gorm.DB) TaskRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
