package db

import (
	"context"
	"time"

	"github.com/LuisFerTR/twitter_clone/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExists given an email address checks if the email address
// already exists in database
func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter_clone")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	return result, err == nil, ID
}