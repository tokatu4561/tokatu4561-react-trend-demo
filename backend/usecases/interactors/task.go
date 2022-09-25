package interactors

import (
	"myapp/domain/model"
	"myapp/usecases/ports"
)

type TaskUsecase struct {
	Repository ports.TaskRepository
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
