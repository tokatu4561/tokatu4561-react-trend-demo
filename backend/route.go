package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Post("/payment-intent", app.GetPaymentIntent)
	mux.Post("/auth/login", app.CreateAuthnicateToken)
	mux.Get("/auth/user", app.GetAuthUser)

	mux.Get("/tasks", app.GetTasks)
	mux.Get("/tasks/${id}", app.GetTask)
	mux.Post("/tasks", app.CreateTask)
	mux.Put("/tasks/${id}", app.UpdateTask)
	mux.Delete("/tasks/${id}", app.DeleteTask)

	return mux
}
