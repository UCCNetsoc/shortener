package views

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/UCCNetsoc/shortener/database"
	"github.com/UCCNetsoc/shortener/models"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

var client *database.Client

// InitViews creates an interface to the database client
func InitViews() {
	client = database.InitDatabase()
}

// PostLink creates a new redirect from given slug and url
func PostLink(w http.ResponseWriter, r *http.Request) {
	var req models.Link
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Slug == "" || req.URL == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	result := setRedirect(&req)
	w.WriteHeader(result)
	switch result {
	case 201:
		log.Println("created slug")
	default:
		log.Println("duplicate slug exists for ", req.Slug)
	}
}

// DeleteLink removes a Link
func DeleteLink(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	ok, err := client.DeleteSlug(slug)
	if !ok {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("slug not found")
			w.WriteHeader(404)
			return
		}
		log.Println("couldnt delete slug\n", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(202)
}

// GetURL resolves and redirects the request
func GetURL(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	redirectURL := getRedirect(slug)
	if redirectURL != "" {
		redirect(w, r, redirectURL)
		return
	}
	w.WriteHeader(404)
}
