package routers

import (
	"encoding/json"
	"github.com/LuisFerTR/twitter_clone/db"
	"github.com/LuisFerTR/twitter_clone/jwt"
	"github.com/LuisFerTR/twitter_clone/models"
	"net/http"
)

// LogIn allows log in the user
func LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Invalid user or password "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Valid email address is required", http.StatusBadRequest)
		return
	}

	document, exists := db.LogInAttempt(t.Email, t.Password)

	if !exists {
		http.Error(w, "Invalid user or password", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Error trying to generate token "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.LogInResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)

	// Cookie example from backend
	//expirationTime := time.Now().Add(24 * time.Hour)
	//http.SetCookie(w, &http.Cookie{
	//	Name: "token",
	//	Value: jwtKey,
	//	Expires: expirationTime,
	//})
}
