package routers

import (
	"net/http"

	"github.com/LuisFerTR/twitter_clone/db"
)

// RemoveTweet removes tweet from database
func RemoveTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "ID parameter must be provided", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)

	if err != nil {
		http.Error(w, "An error has occurred deleting the tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
