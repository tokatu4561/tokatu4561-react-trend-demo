package interactors

import (
	"context"
	"myapp/domain/model"
	"myapp/models"
	"myapp/usecases/ports"
)

type TaskCreateInteractor struct {
	OutputPort ports.TaskOutputPort
	Repository ports.TaskRepository
}

type TaskCreateUseCaseInterface interface {
	AddTask(t *models.Task) func(s string) error
}

func (t *TaskCreateInteractor) AddTask(ctx context.Context, task *model.Task) error {
	users, err := t.Repository.AddTask(ctx, task)
	if err != nil {
		return t.OutputPort.OutputError(err)
	}

	return t.OutputPort.OutputUsers(users)
}
