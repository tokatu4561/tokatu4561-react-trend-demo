package main

import (
	"database/sql"
	"myapp/adapter/controllers"
	"myapp/adapter/gateways"
	"myapp/usecases/interactors"

	"github.com/google/wire"
)

func NewTaskController() *controllers.TaskController {
	wire.Build(controllers.NewTaskController,
		interactors.NewTaskUsecase,
		gateways.NewTaskRepository,
		NewPostgresDatabaseHandler,
	)

	return &controllers.TaskController{}
}

func NewPostgresDatabaseHandler() *sql.DB {
	db, _ := sql.Open("postgres", "user=postgres password=password host=localhost port=5432 dbname=practice sslmode=disable")
	return db
}
