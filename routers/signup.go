package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuisFerTR/twitter_clone/db"
	"github.com/LuisFerTR/twitter_clone/models"
)

// SignUp is a function to create user data in database
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t) // r.Body can only be used once, after that the information is deleted

	if err != nil {
		http.Error(w, "Error in receiving data "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "A user's email address is required to sign up", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters long", http.StatusBadRequest)
		return
	}

	_, found, _ := db.CheckUserExists(t.Email)

	if found {
		http.Error(w, "This email address is already taken", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertSignUpData(t)

	if err != nil {
		http.Error(w, "An error has occurred trying to sign up user", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Sign up user can't be done", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
