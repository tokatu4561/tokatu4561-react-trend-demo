package controllers

import (
	"database/sql"
	"myapp/domain/model"
	"myapp/usecases/ports"
	"net/http"
	"strconv"
)

type InputFactory func(ports.TaskOutputPort, ports.TaskRepository) ports.TaskInputPort
type RepositoryFacgory func(c *sql.DB) ports.TaskRepository
type OutputFactory func(w http.ResponseWriter) ports.TaskOutputPort

type TaskController struct {
	outputFactory     OutputFactory
	inputFactory      InputFactory
	repositoryFactory RepositoryFacgory
}

func NewTaskController(outputFactory OutputFactory, inputFactory InputFactory, repositoryFactory RepositoryFactory) User {
	return &TaskController{
		outputFactory:     outputFactory,
		inputFactory:      inputFactory,
		repositoryFactory: repositoryFactory,
	}
}

func (t *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	t.newInputPort(w).GetTasks()
}

func (t *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, _ := strconv.Atoi(r.Form.Get("user_id"))
	task := model.Task{
		UserID: userID,
		Title:  r.Form.Get("title"),
	}

	t.newInputPort(w).AddTask(ctx, task)
}

func (t *TaskController) newInputPort(w http.ResponseWriter) ports.TaskInputPort {
	outputPort := t.outputFactory(w)
	repository := t.repositoryFactory(t.clientFactory)
	return t.inputFactory(outputPort, repository)
}
