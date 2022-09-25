package main

import (
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

	// mux.Get("/tasks", taskCtl.GetTasks)
	// mux.Post("/tasks", taskCtl.CreateTask)
	// mux.Get("/tasks/${id}", app.GetTask)
	// mux.Put("/tasks/${id}", app.UpdateTask)
	// mux.Delete("/tasks/${id}", app.DeleteTask)

	return mux
}
