package ports

import (
	"context"
	"myapp/domain/model"
)

type TaskInputPort interface {
	AddTask(ctx context.Context, task model.Task) error
}

type TaskOutputPort interface {
	OutputUsers([]*model.Task) error
	OutputError(error) error
}

type TaskRepository interface {
	AddTask(ctx context.Context, user *model.Task) ([]*model.Task, error)
	GetTasks(ctx context.Context) ([]*model.Task, error)
}
