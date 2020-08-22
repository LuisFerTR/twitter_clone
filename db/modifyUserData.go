package db

import (
	"context"
	"time"

	"github.com/LuisFerTR/twitter_clone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyUserData allows modify user profile data
func ModifyUserData(user models.User, ID string)  (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter_clone")
	col := db.Collection("users")

	data := make(map[string]interface{})

	if len(user.Name) > 0 {
		data["name"] = user.Name
	}

	if len(user.LastName) > 0 {
		data["lastName"] = user.LastName
	}

	data["birthDate"] = user.BirthDate

	if len(user.LastName) > 0 {
		data["email"] = user.Email
	}

	if len(user.Name) > 0 {
		data["password"] = user.Password
	}

	if len(user.LastName) > 0 {
		data["avatar"] = user.Avatar
	}

	if len(user.LastName) > 0 {
		data["banner"] = user.Banner
	}

	if len(user.Name) > 0 {
		data["biography" ] = user.Biography
	}

	if len(user.LastName) > 0 {
		data["location"] = user.Location
	}

	if len(user.LastName) > 0 {
		data["website"] = user.Website
	}

	updateString := bson.M{
		"$set": data,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq":objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
