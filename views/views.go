package views

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/UCCNetsoc/shortener/database"
)

var client *database.Client

// InitViews creates an interface to the database client
func InitViews() {
	client = database.InitDatabase()
}

// PostURL create a new url:hash pair
func PostURL(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	url := chi.URLParam(r, "*")
	err := setRedirect(url, slug)
	if err == 409 {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		log.Println("slug conflict for ", slug)
		return
	}
	if err == 201 {
		http.Error(w, http.StatusText(http.StatusCreated), http.StatusCreated)
		log.Println("created redirect to", url, " on slug ", slug)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// GetURL resolves and redirects the request
func GetURL(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	redirect(w, r, getRedirect(slug))
}
