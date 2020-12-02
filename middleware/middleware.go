package middleware

import (
	"log"
	"net/http"
)

// Mid logging requests
func Mid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " request to ", r.URL)
		next.ServeHTTP(w, r)
	})
}
