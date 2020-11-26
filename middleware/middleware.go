package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// Auth checks provided api token
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " request to ", r.URL)
		authHeader := r.Header.Get("Authorization")
		splitAuth := strings.Split(authHeader, "Bearer ")
		if len(splitAuth) != 2 || splitAuth[1] != viper.GetString("api.token") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			log.Println("Unable to authenticate")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Mid logging
func Mid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, " request to ", r.URL)
		next.ServeHTTP(w, r)
	})
}
