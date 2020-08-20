package jwt

import (
	"time"

	"github.com/LuisFerTR/twitter_clone/models"
	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates JWT
func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("(*AprendiendoGoDesdeUdemy15&/")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastName":  t.LastName,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
