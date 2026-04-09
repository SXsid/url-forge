package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello world</h1>"))
	})
	r.Get("/{code}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.PathValue("code"))
	})

	return r
}
