package domains

import (
	"github.com/dipeshdulal/clean-gin/models"
	"gorm.io/gorm"
)

type TaskService interface {
	WithTrx(trxHandle *gorm.DB) TaskService
	GetOneTask(id uint) (models.Task, error)
	GetAllTask() ([]models.Task, error)
	CreateTask(models.Task) error
	DeleteTask(id uint) error
	UpdateTask(id uint, task models.Task) error
}
