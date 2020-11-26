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
func setRedirect(url, slug string) int {

	// checks for duplicates of the first 5 chars, if it exist, go to 6, if still exists, rehash with time
	if !client.Duplicate(slug) {
		redirectURL := models.RedirectURL{Slug: slug, URL: url}
		client.CreateNew(&redirectURL)
		return 201
	}
	return 409
}

// GetRedirect returns the stored url for given hash
func getRedirect(slug string) string {
	// if url not in cache, look for it in database
	value := models.RedirectURL{URL: getFromDB(slug), Slug: slug}
	log.Println("hitting db for ", slug, " redirecting to ", value.URL)

	return value.URL
}

// getFromDB if not in cache
func getFromDB(hash string) string {
	return client.FindByHash(hash).URL
}
