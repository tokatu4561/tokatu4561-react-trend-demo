package ports

import (
	"context"
	"myapp/domain/model"
	"myapp/models"
)

type TaskOutputPort interface {
	OutputUsers([]*models.Task) error
	OutputError(error) error
}

type TaskRepository interface {
	AddTask(ctx context.Context, user *model.Task) ([]*models.Task, error)
	GetTasks(ctx context.Context) ([]*models.Task, error)
}
