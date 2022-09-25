package interactors

import (
	"myapp/domain/model"
	"myapp/usecases/ports"
)

type TaskUsecase struct {
	Repository ports.TaskRepositoryInterface
}

func NewTaskUsecase(taskRepo ports.TaskRepositoryInterface) ports.TaskUseCaseInterface {
	return &TaskUsecase{
		Repository: taskRepo,
	}
}

func (t *TaskUsecase) GetTasks() ([]*model.Task, error) {
	tasks, err := t.Repository.GetTasks()
	if err != nil {
		return nil, err
	}

	return tasks, err
}

func (t *TaskUsecase) AddTask(task *model.Task) (*model.Task, error) {
	newTask, err := t.Repository.AddTask(task)

	if err != nil {
		return nil, err
	}

	return newTask, nil
}
