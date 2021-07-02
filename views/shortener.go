package views

import (
	"log"
	"math/rand"

	"github.com/UCCNetsoc/shortener/models"
)

// SetRedirect generates a new hashmap entry from unix nano time
func setRedirect(req *models.Link) *models.Link {

	// if no slug is provided
	if req.Slug == "" {
		runes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i:=0;i<6;i++ {
			req.Slug += string(runes[rand.Intn(61)])
		}
	}

	if !client.Duplicate(req.Slug) {
		resp, err := client.CreateNew(req)
		if err != nil || resp == nil {
			log.Println("failed to create shortened link")
		}
		return resp
	}
	return nil
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
