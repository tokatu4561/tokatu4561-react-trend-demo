package ports

import (
	"context"
	"myapp/domain/model"
)

type TaskInputPort interface {
	AddTask(ctx context.Context, task model.Task) error
	GetTasks(ctx context.Context) error
}

type TaskOutputPort interface {
	OutputTasks([]*model.Task)
	OutputError(error)
}

type TaskRepository interface {
	AddTask(ctx context.Context, task *model.Task) (*model.Task, error)
	GetTasks(ctx context.Context) ([]*model.Task, error)
}
