package router

import (
	"net/http"
	"time"

	"github/SXsid/url-forge/internal"
	"github/SXsid/url-forge/ui"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(app *internal.Applicaton) http.Handler {
	r := chi.NewRouter()
	fileServer := http.FileServerFS(ui.StaticFS)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST"},
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type"},
		// for dev in 0
		MaxAge: int((time.Hour * 0).Seconds()),
	}))

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
