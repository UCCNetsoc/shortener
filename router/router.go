package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/UCCNetsoc/shortener/mid"
	"github.com/UCCNetsoc/shortener/views"
)

// Route ...
func Route(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Get("/{slug}", views.GetURL)

	r.Mount("/api", apiRouter())
}

func apiRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(mid.Auth)
	r.Get("/links", views.GetLinks)
	r.Post("/", views.PostLink)
	r.Delete("/{slug}", views.DeleteLink)
	return r
}
