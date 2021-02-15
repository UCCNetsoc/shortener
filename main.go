package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/UCCNetsoc/shortener/config"
	"github.com/UCCNetsoc/shortener/router"
	"github.com/UCCNetsoc/shortener/views"
)

func main() {
	r := chi.NewRouter()
	router.Route(r)
	config.InitConfig()
	views.InitViews()
	http.ListenAndServe(":8080", r)
}
