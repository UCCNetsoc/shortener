package views

import (
	"crypto/md5"
	"encoding/base32"
	"log"

	"github.com/UCCNetsoc/shortener/models"
)

// returns the md5 hash of the given url
func generateHash(url string) string {
	digest := md5.New()
	digest.Write([]byte(url))
	return base32.StdEncoding.EncodeToString(digest.Sum(nil))
}

// SetRedirect generates a new hashmap entry from unix nano time
func setRedirect(req *models.Request) int {

	// checks for duplicates of the first 5 chars, if it exist, go to 6, if still exists, rehash with time
	if !client.Duplicate(req.Domain, req.Slug) {
		client.CreateNew(req)
		return 201
	}
	return 409
}

// GetRedirect returns the stored url for given slug
func getRedirect(domain, slug string) string {
	target := client.FindRedirect(domain, slug).Target
	if target != "" {
		log.Println("hitting db for ", slug, " redirecting to ", target)
	}
	// return whether target is "" or otherwise
	return target
}
