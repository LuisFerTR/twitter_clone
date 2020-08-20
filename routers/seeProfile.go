package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuisFerTR/twitter_clone/db"
)

// SeeProfile extract values from profile
func SeeProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Must send ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)

	if err != nil {
		http.Error(w, "An error has occurred searching the profile "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(profile)
}
