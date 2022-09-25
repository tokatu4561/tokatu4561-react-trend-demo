package main

import (
	"database/sql"
	"myapp/adapter/controllers"
	"myapp/adapter/gateways"
	"myapp/usecases/interactors"
	"myapp/usecases/ports"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func routes(mux *chi.Mux) http.Handler {
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	taskCtl := controllers.NewTaskController(InjectTodoUsecase())

	mux.Get("/tasks", taskCtl.GetTasks)
	mux.Post("/tasks", taskCtl.CreateTask)
	// mux.Get("/tasks/${id}", app.GetTask)
	// mux.Put("/tasks/${id}", app.UpdateTask)
	// mux.Delete("/tasks/${id}", app.DeleteTask)

	return mux
}

func InjectDB() gateways.DatabaseHandler {
	db, _ := sql.Open("postgres", "user=postgres password=password host=localhost port=5432 dbname=practice sslmode=disable")
	return gateways.DatabaseHandler{
		Conn: db,
	}
}

func InjectTaskRepository() ports.TaskRepository {
	sqlHandler := InjectDB()
	return gateways.NewTaskRepository(sqlHandler)
}

func InjectTodoUsecase() ports.TaskUseCaseInterface {
	taskRepo := InjectTaskRepository()
	return &interactors.TaskUsecase{
		Repository: taskRepo,
	}
}
