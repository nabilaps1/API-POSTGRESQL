package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nabilaps1/API-POSTGRESQL/configs"
	"github.com/nabilaps1/API-POSTGRESQL/handlers"
)

func main() { // servidor
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	// declarando as rotas
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
