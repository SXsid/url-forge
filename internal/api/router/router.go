package router

import (
	"net/http"

	"github/SXsid/url-forge/internal/api/handler"
	"github/SXsid/url-forge/ui"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	fileServer := http.FileServerFS(ui.StaticFS)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", handler.HomePageHandler)
	r.Handle("/assests/*", fileServer)
	r.Get("/{code}", handler.Redirect)
	r.Post("/api/create", handler.Register)
	return r
}
