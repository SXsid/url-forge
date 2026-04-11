package api

import (
	"net/http"

	"github/SXsid/url-forge/ui"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	fileServer := http.FileServerFS(ui.StaticFS)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", HomePageHandler)
	r.Handle("/assests/*", fileServer)
	r.Get("/{code}", Redirect)
	r.Post("/api/create", Register)
	return r
}
