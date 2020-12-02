package views

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/UCCNetsoc/shortener/database"
	"github.com/UCCNetsoc/shortener/models"
	"github.com/go-chi/chi"
)

var client *database.Client

// InitViews creates an interface to the database client
func InitViews() {
	client = database.InitDatabase()
}

// PostURL create a new url:hash pair
func PostURL(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Slug == "" || req.Domain == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	result := setRedirect(&req)
	if result == 409 {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		log.Println("slug conflict for ", req.Slug)
		return
	}
	if result == 201 {
		http.Error(w, http.StatusText(http.StatusCreated), http.StatusCreated)
		log.Println("created redirect to", req.Target, " on ", req.Domain, req.Slug)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// GetURL resolves and redirects the request
func GetURL(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	domain := r.Host
	redirectURL := getRedirect(domain, slug)
	if redirectURL != "" {
		redirect(w, r, redirectURL)
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
