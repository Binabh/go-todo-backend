package services

import (
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

// UserService service layer
type TaskService struct {
	logger     lib.Logger
	repository repository.TaskRepository
}

// CreateTask implements domains.TaskService.
func (t TaskService) CreateTask(task models.Task) error {
	return t.repository.Create(&task).Error
}

// DeleteTask implements domains.TaskService.
func (t TaskService) DeleteTask(id uint) error {
	return t.repository.Delete(&models.Task{}, id).Error
}

// GetAllTask implements domains.TaskService.
func (t TaskService) GetAllTask() (tasks []models.Task, err error) {
	return tasks, t.repository.Find(&tasks).Error
}

// GetOneTask implements domains.TaskService.
func (t TaskService) GetOneTask(id uint) (task models.Task, err error) {
	return task, t.repository.Find(&task, id).Error
}

// UpdateTask implements domains.TaskService.
func (t TaskService) UpdateTask(id uint, task models.Task) error {
	task.ID = id
	return t.repository.Updates(&task).Error
}

// NewUserService creates a new userservice
func NewTaskService(logger lib.Logger, repository repository.TaskRepository) domains.TaskService {
	return TaskService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s TaskService) WithTrx(trxHandle *gorm.DB) domains.TaskService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}
