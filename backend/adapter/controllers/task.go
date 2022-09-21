package controllers

import (
	"database/sql"
	"myapp/domain/model"
	"myapp/usecases/ports"
	"net/http"
)

type InputFactory func(ports.TaskOutputPort, ports.TaskRepository) ports.TaskInputPort
type RepositoryFacgory func(c *sql.DB) ports.TaskRepository

type TaskController struct {
	outputFactory     func(w http.ResponseWriter) ports.TaskOutputPort
	inputFactory      InputFactory
	repositoryFactory RepositoryFacgory
	clientFactory     *sql.DB
}

// func (t *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) error {

// }

func (t *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	task := model.Task{
		Title: "aaa",
	}

	return func(w http.ResponseWriter) error {
		return t.newInputPort(w).AddTask(ctx, task)
	}
}

func (t *TaskController) newInputPort(w http.ResponseWriter) ports.TaskInputPort {
	outputPort := t.outputFactory(w)
	repository := t.repositoryFactory(t.clientFactory)
	return t.inputFactory(outputPort, repository)
}
