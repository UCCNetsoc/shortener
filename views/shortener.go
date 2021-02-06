package views

import (
	"log"

	"github.com/UCCNetsoc/shortener/models"
)

// SetRedirect generates a new hashmap entry from unix nano time
func setRedirect(req *models.Link) int {
	if !client.Duplicate(req.Slug) {
		client.CreateNew(req)
		return 201
	}
	return 409
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
