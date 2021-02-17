package mid

import (
	"net/http"

	"github.com/spf13/viper"
)

// Auth checks basicauth
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if ok && user == viper.GetString("shortener.user") && pass == viper.GetString("shortener.password") {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(403)
	})
}
