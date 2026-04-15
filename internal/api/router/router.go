package router

import (
	"net/http"

	"github/SXsid/url-forge/internal"
	"github/SXsid/url-forge/ui"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(app *internal.Applicaton) http.Handler {
	r := chi.NewRouter()
	fileServer := http.FileServerFS(ui.StaticFS)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", app.UrlHandler.HomePageHandler)
	r.Handle("/assests/*", fileServer)
	r.Get("/{code}", app.UrlHandler.Redirect)
	r.Post("/api/create", app.UrlHandler.Register)
	// sereve reenere html page
	r.Get("/status/{code}", app.UrlHandler.GetAnylitics)
	return r
}
