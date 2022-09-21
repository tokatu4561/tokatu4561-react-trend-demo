package interactors

import (
	"context"
	"myapp/domain/model"
	"myapp/models"
	"myapp/usecases/ports"
)

type TaskInteractor struct {
	OutputPort ports.TaskOutputPort
	Repository ports.TaskRepository
}

type TaskUseCaseInterface interface {
	AddTask(t *models.Task) func(s string) error
}

func (t *TaskInteractor) AddTask(ctx context.Context, task *model.Task) error {
	users, err := t.Repository.AddTask(ctx, task)
	if err != nil {
		return t.OutputPort.OutputError(err)
	}

	return t.OutputPort.OutputUsers(users)
}
