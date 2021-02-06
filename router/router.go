package router

import (
	"github.com/go-chi/chi"

	"github.com/UCCNetsoc/shortener/middleware"
	"github.com/UCCNetsoc/shortener/views"
)

// Route ...
func Route(r *chi.Mux) {
	r.Get("/{slug}", views.GetURL)
	r.With(middleware.Mid).Post("/", views.PostLink)
	r.With(middleware.Mid).Delete("/{slug}", views.DeleteLink)
}
