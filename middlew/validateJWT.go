package middlew

import (
	"net/http"

	"github.com/LuisFerTR/twitter_clone/routers"
)

// ValidateJWT authenticate request returned JWT
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Error in token! "+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
