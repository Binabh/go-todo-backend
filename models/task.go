package models

import (
	"time"
)

// Task model
type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName gives table name of model
func (t Task) TableName() string {
	return "tasks"
}
