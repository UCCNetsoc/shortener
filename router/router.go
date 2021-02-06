package router

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/UCCNetsoc/shortener/middleware"
	"github.com/UCCNetsoc/shortener/views"
)

// Route ...
func Route(r *chi.Mux) {
	r.Get("/{slug}", views.GetURL)
	r.Mount("/api", apiHandler())
}

func apiHandler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Mid)
	r.Post("/", views.PostLink)
	r.Delete("/{slug}", views.DeleteLink)
	return r
}
