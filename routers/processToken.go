package routers

import (
	"errors"
	"strings"

	"github.com/LuisFerTR/twitter_clone/db"
	"github.com/LuisFerTR/twitter_clone/models"
	"github.com/dgrijalva/jwt-go"
)

// Email value used in all endpoints
var Email string

// IDUser value used in all endpoints
var IDUser string

// ProcessToken process the given token to extract its values later
func ProcessToken(tokenString string) (*models.Claim, bool, string, error) {
	myKey := []byte("(*AprendiendoGoDesdeUdemy15&/")
	claims := &models.Claim{}
	splitToken := strings.Split(tokenString, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("invalid token format")
	}

	tokenString = strings.TrimSpace(splitToken[1])

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := db.CheckUserExists(claims.Email)

		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}

		return claims, found, IDUser, nil
	}

	if token != nil && !token.Valid {
		return nil, false, "", errors.New("invalid token")
	}

	return nil, false, "", err
}
