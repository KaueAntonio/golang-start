package main

import (
	"fmt"
	"golang-start/App/App/Config"
	"golang-start/App/App/Handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := Config.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", Handlers.Create)
	r.Put("/{id}", Handlers.Update)
	r.Delete("/{id}", Handlers.Delete)
	r.Get("/", Handlers.List)
	r.Get("/{id}", Handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", Config.GetServerPort()), r)
}
