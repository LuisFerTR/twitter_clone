package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuisFerTR/twitter_clone/db"
	"github.com/LuisFerTR/twitter_clone/models"
)

// ModifyProfile modifies user profile data
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Wrong data "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err := db.ModifyUserData(t, IDUser)

	if err != nil {
		http.Error(w, "An error has occurred trying to modify user profile data. Try again "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "User profile data couldn't be modified", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}