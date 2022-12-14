package ports

import (
	"myapp/domain/model"
)

type TaskRepositoryInterface interface {
	AddTask(task *model.Task) (*model.Task, error)
	GetTasks() ([]*model.Task, error)
}

type TaskUseCaseInterface interface {
	AddTask(t *model.Task) (*model.Task, error)
	GetTasks() ([]*model.Task, error)
}
