package middlew

import (
	"net/http"

	"github.com/LuisFerTR/twitter_clone/db"
)

// CheckDB is the middleware to know database status
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "Lost database connection", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
