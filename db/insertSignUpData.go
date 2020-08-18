package db

import (
	"context"
	"time"

	"github.com/LuisFerTR/twitter_clone/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertSingUpData is the last stop to insert
// user data in database
func InsertSignUpData(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter_clone")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
