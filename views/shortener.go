package views

import (
	"crypto/rand"
	"log"
	"net/http"

	"github.com/UCCNetsoc/shortener/models"
)

// SetRedirect generates a new hashmap entry from unix nano time
func setRedirect(req *models.Link) int {

	// if no slug is provided
	if req.Slug == "" {
		byteArr := make ([]byte, 5)
		num, err := rand.Read(byteArr)
		if err != nil && num != 5 {
			log.Println("failed to generate a random string\n", err)
			return http.StatusInternalServerError
		}
	}

	if !client.Duplicate(req.Slug) {
		client.CreateNew(req)
		return http.StatusCreated
	}
	return http.StatusConflict
}

// GetRedirect returns the stored url for given slug
func getRedirect(slug string) string {
	target := client.FindRedirect(slug).URL
	if target != "" {
		log.Println("hitting db for ", slug, " redirecting to ", target)
	}
	// return whether target is "" or otherwise
	return target
}
