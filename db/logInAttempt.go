package db

import (
	"github.com/LuisFerTR/twitter_clone/models"
	"golang.org/x/crypto/bcrypt"
)

// LogInAttempt checks credentials in database
func LogInAttempt(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserExists(email)

	if !found {
		return user, false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return models.User{}, false
	}

	return user, true
}
