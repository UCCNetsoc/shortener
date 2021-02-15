package router

import (
	"net/http"

	"github.com/go-chi/chi"
	mid "github.com/go-chi/chi/middleware"

	"github.com/UCCNetsoc/shortener/middleware"
	"github.com/UCCNetsoc/shortener/views"
)

// Route ...
func Route(r *chi.Mux) {
	r.Use(mid.Logger)
	r.Get("/{slug}", views.GetURL)
	r.Post("/api", views.PostLink)
	r.Delete("/api/{slug}", views.DeleteLink)
}