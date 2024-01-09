package main

import (
	"github.com/MatThHeuss/desafio-metaplane/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/savelists", handlers.SaveListsHandler)
	r.Post("/merge", handlers.MergeHandler)
	r.Get("/health", handlers.HealthCheckHandler)

	http.ListenAndServe(":8080", r)
}
