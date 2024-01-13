package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description,omitempty"`
	Deadline    time.Time  `json:"deadline,omitempty"`
	Status      TaskStatus `json:"status" validate:"required,taskstatus"`
}

type TaskStatus string

const (
	TaskStatusTodo    TaskStatus = "TODO"
	TaskStatusDoing   TaskStatus = "DOING"
	TaskStatusDone    TaskStatus = "DONE"
	TaskStatusDeleted TaskStatus = "DELETED"
)
