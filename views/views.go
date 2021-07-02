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

// GetLinks fetches all links from db
func GetLinks(w http.ResponseWriter, r *http.Request) {
	links := client.FetchLinks()
	encoded, err := json.Marshal(links)
	if err != nil {
		log.Println("couldn't fetch links")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
}

// PostLink creates a new redirect from given slug and url
func PostLink(w http.ResponseWriter, r *http.Request) {
	var req models.Link
	log.Println("creating slug")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("couldn't read body or url was not present")
		return
	}
	result := setRedirect(&req)
	encoded, err := json.Marshal(&result)
	if err != nil {
		log.Println("couldn't create link")
		w.WriteHeader(http.StatusConflict)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(encoded)
	log.Println("Encoded json: ", encoded)
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
